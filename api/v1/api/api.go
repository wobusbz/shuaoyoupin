package v1Api

import "shuaoyoupin/internal/server"

type V1Api struct {
	srvRelation server.Relationer
	srvUser     server.ApiUserer
}
