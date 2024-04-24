// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/examservice/answers/submit": {
            "post": {
                "description": "Creates or updates an answer based on the provided request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Answers"
                ],
                "summary": "Create or update answer",
                "parameters": [
                    {
                        "description": "Answer request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AnswerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/dto.AnswerResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            }
        },
        "/examservice/exams": {
            "post": {
                "description": "Create a new exam or update an existing one based on the provided request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exams"
                ],
                "summary": "Create or update an exam",
                "parameters": [
                    {
                        "description": "Exam request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ExamRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/dto.ExamResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            }
        },
        "/examservice/exams/": {
            "get": {
                "description": "Retrieve a list of all exams filtered by topic and subTopic.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exams"
                ],
                "summary": "Get all exams",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by topic",
                        "name": "topic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by subTopic",
                        "name": "subTopic",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.ListExamsResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            }
        },
        "/examservice/exams/{id}": {
            "get": {
                "description": "Retrieve an exam based on the provided exam ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exams"
                ],
                "summary": "Get an exam by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exam ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/dto.Exam"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "404": {
                        "description": "Exam not found",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an exam by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exams"
                ],
                "summary": "Delete an exam by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exam ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/dto.ExamResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "404": {
                        "description": "Exam not found",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            }
        },
        "/examservice/ping/": {
            "get": {
                "description": "Pings the server and returns \"Okay\" if successful.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ping"
                ],
                "summary": "Pings the server.",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Flag indicating whether to ping the database",
                        "name": "db",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.PingResponse"
                        }
                    }
                }
            }
        },
        "/examservice/questions": {
            "post": {
                "description": "Create or update questions based on the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questions"
                ],
                "summary": "Create or update questions",
                "parameters": [
                    {
                        "description": "Question request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.QuestionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created or updated questions",
                        "schema": {
                            "$ref": "#/definitions/dto.QuestionResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            }
        },
        "/examservice/questions/": {
            "get": {
                "description": "Retrieve a list of all questions.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questions"
                ],
                "summary": "Get all questions",
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.ListQuestionResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            }
        },
        "/examservice/questions/{id}": {
            "get": {
                "description": "Retrieves a question based on the provided ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questions"
                ],
                "summary": "Retrieve a question by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/dto.QuestionByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "404": {
                        "description": "Question not found",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a question by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questions"
                ],
                "summary": "Delete a question by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/dto.QuestionResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "404": {
                        "description": "Question not found",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AnswerRequest": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.QuestionAnswer"
                    }
                },
                "examId": {
                    "type": "string"
                },
                "result": {
                    "$ref": "#/definitions/dto.Result"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "dto.AnswerResponse": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.Choice": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.Exam": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "difficulty_level": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "integer"
                },
                "exam_fee": {
                    "type": "number"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "start_time": {
                    "type": "integer"
                },
                "sub_topic": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                }
            }
        },
        "dto.ExamRequest": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "difficulty_level": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "integer"
                },
                "exam_fee": {
                    "type": "number"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "start_time": {
                    "type": "integer"
                },
                "sub_topic": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                }
            }
        },
        "dto.ExamResponse": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.ListExamsResponse": {
            "type": "object",
            "properties": {
                "exam": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Exam"
                    }
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.ListQuestionResponse": {
            "type": "object",
            "properties": {
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Question"
                    }
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.PingResponse": {
            "type": "object",
            "properties": {
                "message": {},
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.Question": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "choices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Choice"
                    }
                },
                "correct": {
                    "type": "string"
                },
                "created_at": {
                    "type": "integer"
                },
                "explanation": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "dto.QuestionAnswer": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string"
                },
                "correctAnswer": {
                    "type": "string"
                },
                "questionId": {
                    "type": "string"
                }
            }
        },
        "dto.QuestionByIdResponse": {
            "type": "object",
            "properties": {
                "question": {
                    "$ref": "#/definitions/dto.Question"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.QuestionRequest": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "choices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Choice"
                    }
                },
                "correct": {
                    "type": "string"
                },
                "explanation": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "dto.QuestionResponse": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.Result": {
            "type": "object",
            "properties": {
                "attempted": {
                    "type": "integer"
                },
                "correct": {
                    "type": "integer"
                }
            }
        },
        "utils.CustomError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
