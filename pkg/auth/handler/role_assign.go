package handler

import (
	"encoding/json"
	"net/http"

	"github.com/skygeario/skygear-server/pkg/auth"
	"github.com/skygeario/skygear-server/pkg/core/auth/authinfo"
	"github.com/skygeario/skygear-server/pkg/core/auth/authz"
	"github.com/skygeario/skygear-server/pkg/core/auth/authz/policy"
	"github.com/skygeario/skygear-server/pkg/core/db"
	"github.com/skygeario/skygear-server/pkg/core/handler"
	"github.com/skygeario/skygear-server/pkg/core/inject"
	"github.com/skygeario/skygear-server/pkg/core/server"
	"github.com/skygeario/skygear-server/pkg/core/skyerr"
)

func AttachRoleAssignHandler(
	server *server.Server,
	authDependency auth.DependencyMap,
) *server.Server {
	server.Handle("/role/assign", &RoleAssignHandlerFactory{
		authDependency,
	}).Methods("OPTIONS", "POST")
	return server
}

type RoleAssignHandlerFactory struct {
	Dependency auth.DependencyMap
}

func (f RoleAssignHandlerFactory) NewHandler(request *http.Request) http.Handler {
	h := &RoleAssignHandler{}
	inject.DefaultRequestInject(h, f.Dependency, request)
	return handler.APIHandlerToHandler(h, h.TxContext)
}

func (f RoleAssignHandlerFactory) ProvideAuthzPolicy() authz.Policy {
	return policy.AllOf(
		authz.PolicyFunc(policy.DenyNoAccessKey),
		authz.PolicyFunc(policy.RequireAuthenticated),
		policy.AnyOf(
			authz.PolicyFunc(policy.RequireAdminRole),
			authz.PolicyFunc(policy.RequireMasterKey),
		),
		authz.PolicyFunc(policy.DenyDisabledUser),
	)
}

type RoleAssignRequestPayload struct {
	Roles   []string `json:"roles"`
	UserIDs []string `json:"users"`
}

func (p RoleAssignRequestPayload) Validate() error {
	if p.Roles == nil || len(p.Roles) == 0 {
		return skyerr.NewInvalidArgument("unspecified roles in request", []string{"roles"})
	}
	if p.UserIDs == nil || len(p.UserIDs) == 0 {
		return skyerr.NewInvalidArgument("unspecified users in request", []string{"users"})
	}

	return nil
}

// RoleAssignHandler allow system administrator to batch assign roles to
// users
//
// RoleAssignHandler required user with admin role.
// All specified users will assign to all roles specified. Roles not already
// existed in DB will be created. Users not already existed will be ignored.
//
// curl -X POST -H "Content-Type: application/json" \
//   -d @- http://localhost:3000/role/assign <<EOF
// {
//     "roles": [
//        "writer",
//        "user"
//     ],
//     "users": [
//        "95db1e34-0cc0-47b0-8a97-3948633ce09f",
//        "3df4b52b-bd58-4fa2-8aee-3d44fd7f974d"
//     ]
// }
// EOF
//
// {
//     "result": "OK"
// }
type RoleAssignHandler struct {
	AuthInfoStore authinfo.Store `dependency:"AuthInfoStore"`
	TxContext     db.TxContext   `dependency:"TxContext"`
}

func (h RoleAssignHandler) WithTx() bool {
	return true
}

func (h RoleAssignHandler) DecodeRequest(request *http.Request) (handler.RequestPayload, error) {
	payload := RoleAssignRequestPayload{}
	err := json.NewDecoder(request.Body).Decode(&payload)
	return payload, err
}

func (h RoleAssignHandler) Handle(req interface{}) (resp interface{}, err error) {
	payload := req.(RoleAssignRequestPayload)
	if err = h.AuthInfoStore.AssignRoles(payload.UserIDs, payload.Roles); err != nil {
		err = skyerr.MakeError(err)
		return
	}

	resp = "OK"
	return
}