package errors

import (
	"reflect"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/internal"
	"sm-box/pkg/errors/internal/grpc"
	"sm-box/pkg/errors/internal/rest_api"
	"sm-box/pkg/errors/internal/ws"
	"sm-box/pkg/errors/types"
	"testing"
)

func TestToError_ForError(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr Error
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000002",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error with details. "),
					Details: new(details.Details).
						Set("key", "value"),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000003",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).Text("Example error with details and fields. "),
					Details: new(details.Details).
						Set("key", "value").
						SetFields(types.DetailsField{
							Key:     new(details.FieldKey).Add("test"),
							Message: new(messages.TextMessage).Text("123"),
						}),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToError(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToError() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToError_ForRestAPI(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr Error
	}

	tests := []testCase[RestAPI]{
		{
			name: "Case 1",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000002",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error with details. "),
					Details: new(details.Details).
						Set("key", "value"),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000003",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).Text("Example error with details and fields. "),
					Details: new(details.Details).
						Set("key", "value").
						SetFields(types.DetailsField{
							Key:     new(details.FieldKey).Add("test"),
							Message: new(messages.TextMessage).Text("123"),
						}),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToError(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToError() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToError_ForWebSocket(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr Error
	}

	tests := []testCase[WebSocket]{
		{
			name: "Case 1",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000002",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error with details. "),
					Details: new(details.Details).
						Set("key", "value"),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000003",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).Text("Example error with details and fields. "),
					Details: new(details.Details).
						Set("key", "value").
						SetFields(types.DetailsField{
							Key:     new(details.FieldKey).Add("test"),
							Message: new(messages.TextMessage).Text("123"),
						}),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToError(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToError() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToError_ForGrpc(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr Error
	}

	tests := []testCase[Grpc]{
		{
			name: "Case 1",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000002",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error with details. "),
					Details: new(details.Details).
						Set("key", "value"),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Error(&internal.Internal{
				Store: &internal.Store{
					ID:     "T-000003",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).Text("Example error with details and fields. "),
					Details: new(details.Details).
						Set("key", "value").
						SetFields(types.DetailsField{
							Key:     new(details.FieldKey).Add("test"),
							Message: new(messages.TextMessage).Text("123"),
						}),

					Others: &internal.StoreOthers{
						RestAPI: &internal.RestAPIStore{
							StatusCode: 500,
						},
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToError(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToError() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToRestAPI_ForError(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr RestAPI
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToRestAPI(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToRestAPI() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToRestAPI_ForRestAPI(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr RestAPI
	}

	tests := []testCase[RestAPI]{
		{
			name: "Case 1",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToRestAPI(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToRestAPI() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToRestAPI_ForWebSocket(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr RestAPI
	}

	tests := []testCase[WebSocket]{
		{
			name: "Case 1",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToRestAPI(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToRestAPI() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToRestAPI_ForGrpc(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr RestAPI
	}

	tests := []testCase[Grpc]{
		{
			name: "Case 1",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: RestAPI(&rest_api.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToRestAPI(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToRestAPI() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToWebSocket_ForError(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr WebSocket
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToWebSocket(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToWebSocket() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToWebSocket_ForRestAPI(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr WebSocket
	}

	tests := []testCase[RestAPI]{
		{
			name: "Case 1",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToWebSocket(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToWebSocket() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToWebSocket_ForWebSocket(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr WebSocket
	}

	tests := []testCase[WebSocket]{
		{
			name: "Case 1",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToWebSocket(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToWebSocket() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToWebSocket_ForGrpc(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr WebSocket
	}

	tests := []testCase[Grpc]{
		{
			name: "Case 1",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: WebSocket(&ws.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToWebSocket(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToWebSocket() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToGrpc_ForError(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr Grpc
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[Error]{
				err: Error(&internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToGrpc(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToGrpc() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToGrpc_ForRestAPI(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr Grpc
	}

	tests := []testCase[RestAPI]{
		{
			name: "Case 1",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[RestAPI]{
				err: RestAPI(&rest_api.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToGrpc(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToGrpc() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToGrpc_ForWebSocket(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr Grpc
	}

	tests := []testCase[WebSocket]{
		{
			name: "Case 1",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[WebSocket]{
				err: WebSocket(&ws.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToGrpc(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToGrpc() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}

func TestToGrpc_ForGrpc(t *testing.T) {
	type args[T Error] struct {
		err T
	}

	type testCase[T Error] struct {
		name       string
		args       args[T]
		wantNewErr Grpc
	}

	tests := []testCase[Grpc]{
		{
			name: "Case 1",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000001",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error. "),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 2",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000002",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).
								Text("Example error with details. "),
							Details: new(details.Details).
								Set("key", "value"),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
		{
			name: "Case 3",
			args: args[Grpc]{
				err: Grpc(&grpc.Internal{
					Internal: &internal.Internal{
						Store: &internal.Store{
							ID:     "T-000003",
							Type:   types.TypeSystem,
							Status: types.StatusFatal,

							Message: new(messages.TextMessage).Text("Example error with details and fields. "),
							Details: new(details.Details).
								Set("key", "value").
								SetFields(types.DetailsField{
									Key:     new(details.FieldKey).Add("test"),
									Message: new(messages.TextMessage).Text("123"),
								}),

							Others: &internal.StoreOthers{
								RestAPI: &internal.RestAPIStore{
									StatusCode: 500,
								},
								WebSocket: &internal.WebSocketStore{
									StatusCode: 500,
								},
							},
						},
					},
				}),
			},
			wantNewErr: Grpc(&grpc.Internal{
				Internal: &internal.Internal{
					Store: &internal.Store{
						ID:     "T-000003",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).Text("Example error with details and fields. "),
						Details: new(details.Details).
							Set("key", "value").
							SetFields(types.DetailsField{
								Key:     new(details.FieldKey).Add("test"),
								Message: new(messages.TextMessage).Text("123"),
							}),

						Others: &internal.StoreOthers{
							RestAPI: &internal.RestAPIStore{
								StatusCode: 500,
							},
							WebSocket: &internal.WebSocketStore{
								StatusCode: 500,
							},
						},
					},
				},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewErr := ToGrpc(tt.args.err); !reflect.DeepEqual(gotNewErr, tt.wantNewErr) {
				t.Errorf("ToGrpc() = %v, want %v", gotNewErr, tt.wantNewErr)
			}
		})
	}
}
