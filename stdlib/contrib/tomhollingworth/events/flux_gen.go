// DO NOT EDIT: This file is autogenerated via the builtin command.

package events

import (
	ast "github.com/influxdata/flux/ast"
	runtime "github.com/influxdata/flux/runtime"
)

func init() {
	runtime.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 17,
					Line:   13,
				},
				File:   "duration.flux",
				Source: "package events\n\n// duration will calculate the duration between records\n// for each record. The duration calculated is between\n// the current record and the next. The last record will\n// compare against either the stopColum (default: _stop)\n// or a stop timestamp value.\n//\n// `timeColumn` - Optional string. Default '_time'. The value used to calculate duration\n// `columnName` - Optional string. Default 'duration'. The name of the result column\n// `stopColumn` - Optional string. Default '_stop'. The name of the column to compare the last record on\n// `stop` - Optional Time. Use a fixed time to compare the last record against instead of stop column.\nbuiltin duration",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 17,
						Line:   13,
					},
					File:   "duration.flux",
					Source: "builtin duration",
					Start: ast.Position{
						Column: 1,
						Line:   13,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 17,
							Line:   13,
						},
						File:   "duration.flux",
						Source: "duration",
						Start: ast.Position{
							Column: 9,
							Line:   13,
						},
					},
				},
				Name: "duration",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 162,
							Line:   13,
						},
						File:   "duration.flux",
						Source: "(<-tables: [A], ?unit: duration, ?timeColumn: string, ?columnName: string, ?stopColumn: string, ?stop: time) => [B] where A: Record, B: Record",
						Start: ast.Position{
							Column: 20,
							Line:   13,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{&ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 151,
								Line:   13,
							},
							File:   "duration.flux",
							Source: "A: Record",
							Start: ast.Position{
								Column: 142,
								Line:   13,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 151,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 145,
									Line:   13,
								},
							},
						},
						Name: "Record",
					}},
					Tvar: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 143,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "A",
								Start: ast.Position{
									Column: 142,
									Line:   13,
								},
							},
						},
						Name: "A",
					},
				}, &ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 162,
								Line:   13,
							},
							File:   "duration.flux",
							Source: "B: Record",
							Start: ast.Position{
								Column: 153,
								Line:   13,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 162,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 156,
									Line:   13,
								},
							},
						},
						Name: "Record",
					}},
					Tvar: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 154,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "B",
								Start: ast.Position{
									Column: 153,
									Line:   13,
								},
							},
						},
						Name: "B",
					},
				}},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 135,
								Line:   13,
							},
							File:   "duration.flux",
							Source: "(<-tables: [A], ?unit: duration, ?timeColumn: string, ?columnName: string, ?stopColumn: string, ?stop: time) => [B]",
							Start: ast.Position{
								Column: 20,
								Line:   13,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 34,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "<-tables: [A]",
								Start: ast.Position{
									Column: 21,
									Line:   13,
								},
							},
						},
						Kind: "Pipe",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 29,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "tables",
									Start: ast.Position{
										Column: 23,
										Line:   13,
									},
								},
							},
							Name: "tables",
						},
						Ty: &ast.ArrayType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 34,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "[A]",
									Start: ast.Position{
										Column: 31,
										Line:   13,
									},
								},
							},
							ElementType: &ast.TvarType{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 33,
											Line:   13,
										},
										File:   "duration.flux",
										Source: "A",
										Start: ast.Position{
											Column: 32,
											Line:   13,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 33,
												Line:   13,
											},
											File:   "duration.flux",
											Source: "A",
											Start: ast.Position{
												Column: 32,
												Line:   13,
											},
										},
									},
									Name: "A",
								},
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 51,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "?unit: duration",
								Start: ast.Position{
									Column: 36,
									Line:   13,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 41,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "unit",
									Start: ast.Position{
										Column: 37,
										Line:   13,
									},
								},
							},
							Name: "unit",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 51,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "duration",
									Start: ast.Position{
										Column: 43,
										Line:   13,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 51,
											Line:   13,
										},
										File:   "duration.flux",
										Source: "duration",
										Start: ast.Position{
											Column: 43,
											Line:   13,
										},
									},
								},
								Name: "duration",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 72,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "?timeColumn: string",
								Start: ast.Position{
									Column: 53,
									Line:   13,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 64,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "timeColumn",
									Start: ast.Position{
										Column: 54,
										Line:   13,
									},
								},
							},
							Name: "timeColumn",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 72,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "string",
									Start: ast.Position{
										Column: 66,
										Line:   13,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 72,
											Line:   13,
										},
										File:   "duration.flux",
										Source: "string",
										Start: ast.Position{
											Column: 66,
											Line:   13,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 93,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "?columnName: string",
								Start: ast.Position{
									Column: 74,
									Line:   13,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 85,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "columnName",
									Start: ast.Position{
										Column: 75,
										Line:   13,
									},
								},
							},
							Name: "columnName",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 93,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "string",
									Start: ast.Position{
										Column: 87,
										Line:   13,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 93,
											Line:   13,
										},
										File:   "duration.flux",
										Source: "string",
										Start: ast.Position{
											Column: 87,
											Line:   13,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 114,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "?stopColumn: string",
								Start: ast.Position{
									Column: 95,
									Line:   13,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 106,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "stopColumn",
									Start: ast.Position{
										Column: 96,
										Line:   13,
									},
								},
							},
							Name: "stopColumn",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 114,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "string",
									Start: ast.Position{
										Column: 108,
										Line:   13,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 114,
											Line:   13,
										},
										File:   "duration.flux",
										Source: "string",
										Start: ast.Position{
											Column: 108,
											Line:   13,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 127,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "?stop: time",
								Start: ast.Position{
									Column: 116,
									Line:   13,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 121,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "stop",
									Start: ast.Position{
										Column: 117,
										Line:   13,
									},
								},
							},
							Name: "stop",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 127,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "time",
									Start: ast.Position{
										Column: 123,
										Line:   13,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 127,
											Line:   13,
										},
										File:   "duration.flux",
										Source: "time",
										Start: ast.Position{
											Column: 123,
											Line:   13,
										},
									},
								},
								Name: "time",
							},
						},
					}},
					Return: &ast.ArrayType{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 135,
									Line:   13,
								},
								File:   "duration.flux",
								Source: "[B]",
								Start: ast.Position{
									Column: 132,
									Line:   13,
								},
							},
						},
						ElementType: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 134,
										Line:   13,
									},
									File:   "duration.flux",
									Source: "B",
									Start: ast.Position{
										Column: 133,
										Line:   13,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 134,
											Line:   13,
										},
										File:   "duration.flux",
										Source: "B",
										Start: ast.Position{
											Column: 133,
											Line:   13,
										},
									},
								},
								Name: "B",
							},
						},
					},
				},
			},
		}},
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "duration.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 15,
						Line:   1,
					},
					File:   "duration.flux",
					Source: "package events",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 15,
							Line:   1,
						},
						File:   "duration.flux",
						Source: "events",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "events",
			},
		},
	}},
	Package: "events",
	Path:    "contrib/tomhollingworth/events",
}