// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-09-14 12:59:02.300437264 -0400 EDT m=+0.030366227

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "PID.",
        "contact": {
            "name": "Nicholas Spangler",
            "email": "nicholas.spangler@charter.com"
        },
        "license": {},
        "version": "2.0"
    },
    "host": "localhost:9005",
    "basePath": "/peek",
    "paths": {
        "/v2/login/{login}": {
            "get": {
                "description": "PID.",
                "consumes": [
                    "application/json"
                ],
                "summary": "Finds a PID and returns the full account detail object",
                "parameters": [
                    {
                        "type": "string",
                        "description": "PID",
                        "name": "login",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Agent"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Agent": {
            "type": "object",
            "properties": {
                "Area": {
                    "type": "string"
                },
                "Avaya": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Avaya"
                    }
                },
                "Company": {
                    "type": "string"
                },
                "Division": {
                    "type": "string"
                },
                "EID": {
                    "type": "string"
                },
                "FirstName": {
                    "type": "string"
                },
                "LastName": {
                    "type": "string"
                },
                "Location": {
                    "type": "string"
                },
                "PID": {
                    "type": "string"
                },
                "Region": {
                    "type": "string"
                }
            }
        },
        "models.Avaya": {
            "type": "object",
            "properties": {
                "LoginId": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Skills": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Skill"
                    }
                },
                "Switch": {
                    "type": "string"
                },
                "Updated": {
                    "type": "string"
                }
            }
        },
        "models.Skill": {
            "type": "object",
            "properties": {
                "ACD": {
                    "type": "integer"
                },
                "BLGVENDOR": {
                    "type": "string"
                },
                "CHANGE_ID": {
                    "type": "string"
                },
                "CHANGE_NOTES": {
                    "type": "string"
                },
                "COMMENTS": {
                    "type": "string"
                },
                "COMPANY": {
                    "type": "string"
                },
                "CUSTCALLCONTACT": {
                    "type": "string"
                },
                "CUSTTYPE": {
                    "type": "string"
                },
                "DISPLAYNM": {
                    "type": "string"
                },
                "END_DATE": {
                    "type": "string"
                },
                "FACENM": {
                    "type": "string"
                },
                "FCREXCLUSION": {
                    "type": "string"
                },
                "FCRGROUPING": {
                    "type": "string"
                },
                "FCROVERRIDE": {
                    "type": "integer"
                },
                "FCSTGRP": {
                    "type": "string"
                },
                "HIGHLVLDESC": {
                    "type": "string"
                },
                "INSRT_TS": {
                    "type": "string"
                },
                "LANG": {
                    "type": "string"
                },
                "LOB_TYPE": {
                    "type": "string"
                },
                "Level": {
                    "type": "string"
                },
                "MARKET": {
                    "type": "string"
                },
                "MSO": {
                    "type": "string"
                },
                "NRTSUMRPTID": {
                    "type": "integer"
                },
                "OWNER_DESC": {
                    "type": "string"
                },
                "Ordinal": {
                    "type": "integer"
                },
                "QUEUEID": {
                    "type": "integer"
                },
                "QUEUETYPE": {
                    "type": "string"
                },
                "SECONDARYSWITCHNM": {
                    "type": "string"
                },
                "STAFFHRSEXCLUDE": {
                    "type": "string"
                },
                "START_DATE": {
                    "type": "string"
                },
                "SUBLOB": {
                    "type": "string"
                },
                "SUBPRODUCTFIELD": {
                    "type": "string"
                },
                "SVCLVLSEC": {
                    "type": "integer"
                },
                "SWITCHNM": {
                    "type": "string"
                },
                "SYSTEMSOURCECODE": {
                    "type": "string"
                },
                "SYS_AUTH_ID": {
                    "type": "string"
                },
                "SYS_CREATED_BY": {
                    "type": "string"
                },
                "SYS_CREATION_DATE": {
                    "type": "string"
                },
                "SYS_ENT_STATE": {
                    "type": "string"
                },
                "SYS_ERR_CODE": {
                    "type": "string"
                },
                "SYS_ERR_SVRTY": {
                    "type": "string"
                },
                "SYS_LAST_MODIFIED_BY": {
                    "type": "string"
                },
                "SYS_LAST_MODIFIED_DATE": {
                    "type": "string"
                },
                "SYS_NC_TYPE": {
                    "type": "string"
                },
                "SYS_SOURCE": {
                    "type": "string"
                },
                "SYS_TARGET_ID": {
                    "type": "integer"
                },
                "SkillId": {
                    "type": "string"
                },
                "acd": {
                    "type": "integer"
                },
                "hostname": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}