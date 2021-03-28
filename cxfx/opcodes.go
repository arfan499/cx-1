// +build cxfx

// 24:58
// 16/02 : 20:20->23:50
package cxfx

import (
	. "github.com/skycoin/cx/cx"
	"github.com/skycoin/cx/cx/ast"
	"github.com/skycoin/cx/cx/constants"
	"github.com/skycoin/cx/cxos"
)

const (
	// gogl
	OP_GL_INIT = iota + cxos.END_OF_OS_OPS
	OP_GL_DESTROY
	OP_GL_STRS
	OP_GL_FREE
	OP_GL_NEW_TEXTURE
	OP_GL_NEW_TEXTURE_CUBE
	OP_GL_NEW_GIF
	OP_GL_FREE_GIF
	OP_GL_GIF_FRAME_TO_TEXTURE
	OP_GL_UPLOAD_IMAGE_TO_TEXTURE
	OP_CX_RELEASE_TEXTURE
	OP_CX_TEXTURE_GET_PIXEL
	OP_GL_APPEND_F32
	OP_GL_APPEND_UI32
	OP_GL_APPEND_UI16

	// gl_0_0
	OP_GL_MATRIX_MODE
	OP_GL_ROTATEF
	OP_GL_TRANSLATEF
	OP_GL_LOAD_IDENTITY
	OP_GL_PUSH_MATRIX
	OP_GL_POP_MATRIX
	OP_GL_ENABLE_CLIENT_STATE
	OP_GL_COLOR3F
	OP_GL_COLOR4F
	OP_GL_BEGIN
	OP_GL_END
	OP_GL_NORMAL3F
	OP_GL_VERTEX_2F
	OP_GL_VERTEX_3F
	OP_GL_LIGHTFV
	OP_GL_FRUSTUM
	OP_GL_TEX_ENVI
	OP_GL_ORTHO
	OP_GL_SCALEF
	OP_GL_TEX_COORD_2D
	OP_GL_TEX_COORD_2F

	// gl_1_0
	OP_GL_CULL_FACE
	OP_GL_FRONT_FACE
	OP_GL_HINT
	OP_GL_SCISSOR
	OP_GL_TEX_PARAMETERI
	OP_GL_TEX_IMAGE_2D
	OP_GL_CLEAR
	OP_GL_CLEAR_COLOR
	OP_GL_CLEAR_STENCIL
	OP_GL_CLEAR_DEPTH
	OP_GL_STENCIL_MASK
	OP_GL_COLOR_MASK
	OP_GL_DEPTH_MASK
	OP_GL_DISABLE
	OP_GL_ENABLE
	OP_GL_BLEND_FUNC
	OP_GL_STENCIL_FUNC
	OP_GL_STENCIL_OP
	OP_GL_DEPTH_FUNC
	OP_GL_GET_ERROR
	OP_GL_GET_TEX_LEVEL_PARAMETERIV
	OP_GL_DEPTH_RANGE
	OP_GL_VIEWPORT

	// gl_1_1
	OP_GL_DRAW_ARRAYS
	OP_GL_DRAW_ELEMENTS
	OP_GL_BIND_TEXTURE
	OP_GL_DELETE_TEXTURES
	OP_GL_GEN_TEXTURES

	// gl_1_3
	OP_GL_ACTIVE_TEXTURE

	// gl_1_4
	OP_GL_BLEND_FUNC_SEPARATE

	// gl_1_5
	OP_GL_BIND_BUFFER
	OP_GL_DELETE_BUFFERS
	OP_GL_GEN_BUFFERS
	OP_GL_BUFFER_DATA
	OP_GL_BUFFER_SUB_DATA

	// gl_2_0
	OP_GL_DRAW_BUFFERS
	OP_GL_STENCIL_OP_SEPARATE
	OP_GL_STENCIL_FUNC_SEPARATE
	OP_GL_STENCIL_MASK_SEPARATE
	OP_GL_ATTACH_SHADER
	OP_GL_BIND_ATTRIB_LOCATION
	OP_GL_COMPILE_SHADER
	OP_GL_CREATE_PROGRAM
	OP_GL_CREATE_SHADER
	OP_GL_DELETE_PROGRAM
	OP_GL_DELETE_SHADER
	OP_GL_DETACH_SHADER
	OP_GL_ENABLE_VERTEX_ATTRIB_ARRAY
	OP_GL_GET_ATTRIB_LOCATION
	OP_GL_GET_PROGRAMIV
	OP_GL_GET_PROGRAM_INFO_LOG
	OP_GL_GET_SHADERIV
	OP_GL_GET_SHADER_INFO_LOG
	OP_GL_GET_UNIFORM_LOCATION
	OP_GL_LINK_PROGRAM
	OP_GL_SHADER_SOURCE
	OP_GL_USE_PROGRAM
	OP_GL_UNIFORM_1F
	OP_GL_UNIFORM_2F
	OP_GL_UNIFORM_3F
	OP_GL_UNIFORM_4F
	OP_GL_UNIFORM_1I
	OP_GL_UNIFORM_2I
	OP_GL_UNIFORM_3I
	OP_GL_UNIFORM_4I
	OP_GL_UNIFORM_1FV
	OP_GL_UNIFORM_2FV
	OP_GL_UNIFORM_3FV
	OP_GL_UNIFORM_4FV
	OP_GL_UNIFORM_1IV
	OP_GL_UNIFORM_2IV
	OP_GL_UNIFORM_3IV
	OP_GL_UNIFORM_4IV
	OP_GL_UNIFORM_MATRIX_2FV
	OP_GL_UNIFORM_MATRIX_3FV
	OP_GL_UNIFORM_MATRIX_4FV
	OP_GL_UNIFORM_V4F
	OP_GL_UNIFORM_M44F
	OP_GL_UNIFORM_M44FV
	OP_GL_VERTEX_ATTRIB_POINTER
	OP_GL_VERTEX_ATTRIB_POINTER_I32

	// gl_3_0
	OP_GL_CLEAR_BUFFER_I
	OP_GL_CLEAR_BUFFER_UI
	OP_GL_CLEAR_BUFFER_F
	OP_GL_BIND_RENDERBUFFER
	OP_GL_DELETE_RENDERBUFFERS
	OP_GL_GEN_RENDERBUFFERS
	OP_GL_RENDERBUFFER_STORAGE
	OP_GL_BIND_FRAMEBUFFER
	OP_GL_DELETE_FRAMEBUFFERS
	OP_GL_GEN_FRAMEBUFFERS
	OP_GL_CHECK_FRAMEBUFFER_STATUS
	OP_GL_FRAMEBUFFER_TEXTURE_2D
	OP_GL_FRAMEBUFFER_RENDERBUFFER
	OP_GL_GENERATE_MIPMAP
	OP_GL_BIND_VERTEX_ARRAY
	OP_GL_BIND_VERTEX_ARRAY_CORE
	OP_GL_DELETE_VERTEX_ARRAYS
	OP_GL_DELETE_VERTEX_ARRAYS_CORE
	OP_GL_GEN_VERTEX_ARRAYS
	OP_GL_GEN_VERTEX_ARRAYS_CORE

	// goglfw
	OP_GLFW_FULLSCREEN

	// glfw
	OP_GLFW_INIT
	OP_GLFW_WINDOW_HINT
	OP_GLFW_CREATE_WINDOW
	OP_GLFW_MAKE_CONTEXT_CURRENT
	OP_GLFW_SHOULD_CLOSE
	OP_GLFW_SET_SHOULD_CLOSE
	OP_GLFW_POLL_EVENTS
	OP_GLFW_SWAP_BUFFERS
	OP_GLFW_GET_FRAMEBUFFER_SIZE
	OP_GLFW_GET_WINDOW_POS
	OP_GLFW_GET_WINDOW_SIZE
	OP_GLFW_GET_TIME
	OP_GLFW_SWAP_INTERVAL
	OP_GLFW_SET_START_CALLBACK
	OP_GLFW_SET_STOP_CALLBACK
	OP_GLFW_SET_KEYBOARD_CALLBACK
	OP_GLFW_SET_MOUSE_CALLBACK
	OP_GLFW_SET_FRAMEBUFFER_SIZE_CALLBACK
	OP_GLFW_SET_WINDOW_POS_CALLBACK
	OP_GLFW_SET_WINDOW_SIZE_CALLBACK
	OP_GLFW_SET_KEY_CALLBACK
	OP_GLFW_SET_MOUSE_BUTTON_CALLBACK
	OP_GLFW_SET_CURSOR_POS_CALLBACK
	OP_GLFW_GET_CURSOR_POS
	OP_GLFW_SET_INPUT_MODE
	OP_GLFW_SET_WINDOW_POS
	OP_GLFW_GET_KEY
	OP_GLFW_FUNC_I32_I32
	OP_GLFW_CALL_I32_I32
	OP_GLFW_GET_WINDOW_CONTENT_SCALE
	OP_GLFW_GET_MONITOR_CONTENT_SCALE

	// gltext
	OP_GLTEXT_LOAD_TRUE_TYPE
	OP_GLTEXT_LOAD_TRUE_TYPE_CORE
	OP_GLTEXT_PRINTF
	OP_GLTEXT_METRICS
	OP_GLTEXT_TEXTURE
	OP_GLTEXT_NEXT_GLYPH
	OP_GLTEXT_GLYPH_BOUNDS
	OP_GLTEXT_GLYPH_METRICS
	OP_GLTEXT_GLYPH_INFO

	// goal
	OP_AL_LOAD_WAV

	// openal
	OP_AL_CLOSE_DEVICE
	OP_AL_DELETE_BUFFERS
	OP_AL_DELETE_SOURCES
	OP_AL_DEVICE_ERROR
	OP_AL_ERROR
	OP_AL_EXTENSIONS
	OP_AL_OPEN_DEVICE
	OP_AL_PAUSE_SOURCES
	OP_AL_PLAY_SOURCES
	OP_AL_RENDERER
	OP_AL_REWIND_SOURCES
	OP_AL_STOP_SOURCES
	OP_AL_VENDOR
	OP_AL_VERSION
	OP_AL_GEN_BUFFERS
	OP_AL_BUFFER_DATA
	OP_AL_GEN_SOURCES
	OP_AL_SOURCE_BUFFERS_PROCESSED
	OP_AL_SOURCE_BUFFERS_QUEUED
	OP_AL_SOURCE_QUEUE_BUFFERS
	OP_AL_SOURCE_STATE
	OP_AL_SOURCE_UNQUEUE_BUFFERS
)

func init() {
	// gogl
	ast.Op_V2(OP_GL_INIT, "gl.Init", opGlInit, nil, nil)
	ast.Op_V2(OP_GL_DESTROY, "gl.Destroy", opGlDestroy, nil, nil)
	ast.Op_V2(OP_GL_STRS, "gl.Strs", opGlStrs, In(ASTR, ASTR), nil)
	ast.Op_V2(OP_GL_FREE, "gl.Free", opGlFree, In(ASTR), nil)
	ast.Op_V2(OP_GL_NEW_TEXTURE, "gl.NewTexture", opGlNewTexture, In(ASTR), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_NEW_TEXTURE_CUBE, "gl.NewTextureCube", opGlNewTextureCube, In(ASTR, ASTR), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_NEW_GIF, "gl.NewGIF", opGlNewGIF, In(ASTR), Out(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32))
	ast.Op_V2(OP_GL_FREE_GIF, "gl.FreeGIF", opGlFreeGIF, In(ASTR), nil)
	ast.Op_V2(OP_GL_GIF_FRAME_TO_TEXTURE, "gl.GIFFrameToTexture", opGlGIFFrameToTexture, In(ASTR, ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_I32, ConstCxArg_I32))
	ast.Op_V2(OP_GL_UPLOAD_IMAGE_TO_TEXTURE, "gl.UploadImageToTexture", opGlUploadImageToTexture, In(ASTR, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_BOOL), nil)
	ast.Op_V2(OP_CX_RELEASE_TEXTURE, "gl.CxReleaseTexture", opCxReleaseTexture, In(ASTR), nil)
	ast.Op_V2(OP_CX_TEXTURE_GET_PIXEL, "gl.CxTextureGetPixel", opCxTextureGetPixel, In(ASTR, ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32))

	ast.Op_V2(OP_GL_APPEND_F32, "gl.AppendF32", opGlAppend, In(ast.Slice(constants.TYPE_UI8), ConstCxArg_F32), Out(ast.Slice(constants.TYPE_UI8)))
	ast.Op_V2(OP_GL_APPEND_UI16, "gl.AppendUI16", opGlAppend, In(ast.Slice(constants.TYPE_UI8), ConstCxArg_UI16), Out(ast.Slice(constants.TYPE_UI8)))
	ast.Op_V2(OP_GL_APPEND_UI32, "gl.AppendUI32", opGlAppend, In(ast.Slice(constants.TYPE_UI8), ConstCxArg_UI32), Out(ast.Slice(constants.TYPE_UI8)))

	// gl_0.0
	ast.Op_V2(OP_GL_MATRIX_MODE, "gl.MatrixMode", opGlMatrixMode, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_ROTATEF, "gl.Rotatef", opGlRotatef, In(ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_TRANSLATEF, "gl.Translatef", opGlTranslatef, In(ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_LOAD_IDENTITY, "gl.LoadIdentity", opGlLoadIdentity, nil, nil)
	ast.Op_V2(OP_GL_PUSH_MATRIX, "gl.PushMatrix", opGlPushMatrix, nil, nil)
	ast.Op_V2(OP_GL_POP_MATRIX, "gl.PopMatrix", opGlPopMatrix, nil, nil)
	ast.Op_V2(OP_GL_ENABLE_CLIENT_STATE, "gl.EnableClientState", opGlEnableClientState, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_COLOR3F, "gl.Color3f", opGlColor3f, In(ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_COLOR4F, "gl.Color4f", opGlColor4f, In(ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_BEGIN, "gl.Begin", opGlBegin, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_END, "gl.End", opGlEnd, nil, nil)
	ast.Op_V2(OP_GL_NORMAL3F, "gl.Normal3f", opGlNormal3f, In(ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_VERTEX_2F, "gl.Vertex2f", opGlVertex2f, In(ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_VERTEX_3F, "gl.Vertex3f", opGlVertex3f, In(ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_LIGHTFV, "gl.Lightfv", opGlLightfv, In(ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_FRUSTUM, "gl.Frustum", opGlFrustum, In(ConstCxArg_F64, ConstCxArg_F64, ConstCxArg_F64, ConstCxArg_F64, ConstCxArg_F64, ConstCxArg_F64), nil)
	ast.Op_V2(OP_GL_TEX_ENVI, "gl.TexEnvi", opGlTexEnvi, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_ORTHO, "gl.Ortho", opGlOrtho, In(ConstCxArg_F64, ConstCxArg_F64, ConstCxArg_F64, ConstCxArg_F64, ConstCxArg_F64, ConstCxArg_F64), nil)
	ast.Op_V2(OP_GL_SCALEF, "gl.Scalef", opGlScalef, In(ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_TEX_COORD_2D, "gl.TexCoord2d", opGlTexCoord2d, In(ConstCxArg_F64, ConstCxArg_F64), nil)
	ast.Op_V2(OP_GL_TEX_COORD_2F, "gl.TexCoord2f", opGlTexCoord2f, In(ConstCxArg_F32, ConstCxArg_F32), nil)

	// gl_1_0
	ast.Op_V2(OP_GL_CULL_FACE, "gl.CullFace", opGlCullFace, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_FRONT_FACE, "gl.FrontFace", opGlFrontFace, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_HINT, "gl.Hint", opGlHint, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_SCISSOR, "gl.Scissor", opGlScissor, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_TEX_PARAMETERI, "gl.TexParameteri", opGlTexParameteri, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_TEX_IMAGE_2D, "gl.TexImage2D", opGlTexImage2D, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ast.Slice(constants.TYPE_UI32)), nil)
	ast.Op_V2(OP_GL_CLEAR, "gl.Clear", opGlClear, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_CLEAR_COLOR, "gl.ClearColor", opGlClearColor, In(ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_CLEAR_STENCIL, "gl.ClearStencil", opGlClearStencil, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_CLEAR_DEPTH, "gl.ClearDepth", opGlClearDepth, In(ConstCxArg_F64), nil)
	ast.Op_V2(OP_GL_STENCIL_MASK, "gl.StencilMask", opGlStencilMask, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_COLOR_MASK, "gl.ColorMask", opGlColorMask, In(ConstCxArg_BOOL, ConstCxArg_BOOL, ConstCxArg_BOOL, ConstCxArg_BOOL), nil)
	ast.Op_V2(OP_GL_DEPTH_MASK, "gl.DepthMask", opGlDepthMask, In(ConstCxArg_BOOL), nil)
	ast.Op_V2(OP_GL_DISABLE, "gl.Disable", opGlDisable, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_ENABLE, "gl.Enable", opGlEnable, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_BLEND_FUNC, "gl.BlendFunc", opGlBlendFunc, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_STENCIL_FUNC, "gl.StencilFunc", opGlStencilFunc, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_STENCIL_OP, "gl.StencilOp", opGlStencilOp, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_DEPTH_FUNC, "gl.DepthFunc", opGlDepthFunc, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_GET_ERROR, "gl.GetError", opGlGetError, nil, Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_GET_TEX_LEVEL_PARAMETERIV, "gl.GetTexLevelParameteriv", opGlGetTexLevelParameteriv, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_DEPTH_RANGE, "gl.DepthRange", opGlDepthRange, In(ConstCxArg_F64, ConstCxArg_F64), nil)
	ast.Op_V2(OP_GL_VIEWPORT, "gl.Viewport", opGlViewport, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)

	// gl_1_1
	ast.Op_V2(OP_GL_DRAW_ARRAYS, "gl.DrawArrays", opGlDrawArrays, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_DRAW_ELEMENTS, "gl.DrawElements", opGlDrawElements, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_BIND_TEXTURE, "gl.BindTexture", opGlBindTexture, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_DELETE_TEXTURES, "gl.DeleteTextures", opGlDeleteTextures, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_GEN_TEXTURES, "gl.GenTextures", opGlGenTextures, In(ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_I32))

	// gl_1_3
	ast.Op_V2(OP_GL_ACTIVE_TEXTURE, "gl.ActiveTexture", opGlActiveTexture, In(ConstCxArg_I32), nil)

	// gl_1_4
	ast.Op_V2(OP_GL_BLEND_FUNC_SEPARATE, "gl.BlendFuncSeparate", opGlBlendFuncSeparate, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)

	// gl_1_5
	ast.Op_V2(OP_GL_BIND_BUFFER, "gl.BindBuffer", opGlBindBuffer, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_DELETE_BUFFERS, "gl.DeleteBuffers", opGlDeleteBuffers, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_GEN_BUFFERS, "gl.GenBuffers", opGlGenBuffers, In(ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_BUFFER_DATA, "gl.BufferData", opGlBufferData, In(ConstCxArg_I32, ConstCxArg_I32, ast.Slice(constants.TYPE_UI8), ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_BUFFER_SUB_DATA, "gl.BufferSubData", opGlBufferSubData, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ast.Slice(constants.TYPE_UI8)), nil)

	//gl_2_0
	ast.Op_V2(OP_GL_DRAW_BUFFERS, "gl.DrawBuffers", opGlDrawBuffers, In(ConstCxArg_I32, ast.Slice(constants.TYPE_UI32)), nil)
	ast.Op_V2(OP_GL_STENCIL_OP_SEPARATE, "gl.StencilOpSeparate", opGlStencilOpSeparate, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_STENCIL_FUNC_SEPARATE, "gl.StencilFuncSeparate", opGlStencilFuncSeparate, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_STENCIL_MASK_SEPARATE, "gl.StencilMaskSeparate", opGlStencilMaskSeparate, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_ATTACH_SHADER, "gl.AttachShader", opGlAttachShader, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_BIND_ATTRIB_LOCATION, "gl.BindAttribLocation", opGlBindAttribLocation, In(ConstCxArg_I32, ConstCxArg_I32, ASTR), nil)
	ast.Op_V2(OP_GL_COMPILE_SHADER, "gl.CompileShader", opGlCompileShader, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_CREATE_PROGRAM, "gl.CreateProgram", opGlCreateProgram, nil, Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_CREATE_SHADER, "gl.CreateShader", opGlCreateShader, In(ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_DELETE_PROGRAM, "gl.DeleteProgram", opGlDeleteProgram, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_DELETE_SHADER, "gl.DeleteShader", opGlDeleteShader, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_DETACH_SHADER, "gl.DetachShader", opGlDetachShader, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_ENABLE_VERTEX_ATTRIB_ARRAY, "gl.EnableVertexAttribArray", opGlEnableVertexAttribArray, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_GET_ATTRIB_LOCATION, "gl.GetAttribLocation", opGlGetAttribLocation, In(ConstCxArg_I32, ASTR), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_GET_PROGRAMIV, "gl.GetProgramiv", opGlGetProgramiv, In(ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_GET_PROGRAM_INFO_LOG, "gl.GetProgramInfoLog", opGlGetProgramInfoLog, In(ConstCxArg_I32, ConstCxArg_I32), Out(ASTR))
	ast.Op_V2(OP_GL_GET_SHADERIV, "gl.GetShaderiv", opGlGetShaderiv, In(ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_GET_SHADER_INFO_LOG, "gl.GetShaderInfoLog", opGlGetShaderInfoLog, In(ConstCxArg_I32, ConstCxArg_I32), Out(ASTR))
	ast.Op_V2(OP_GL_GET_UNIFORM_LOCATION, "gl.GetUniformLocation", opGlGetUniformLocation, In(ConstCxArg_I32, ASTR), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_LINK_PROGRAM, "gl.LinkProgram", opGlLinkProgram, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_SHADER_SOURCE, "gl.ShaderSource", opGlShaderSource, In(ConstCxArg_I32, ConstCxArg_I32, ASTR), nil)
	ast.Op_V2(OP_GL_USE_PROGRAM, "gl.UseProgram", opGlUseProgram, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_UNIFORM_1F, "gl.Uniform1f", opGlUniform1f, In(ConstCxArg_I32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_UNIFORM_2F, "gl.Uniform2f", opGlUniform2f, In(ConstCxArg_I32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_UNIFORM_3F, "gl.Uniform3f", opGlUniform3f, In(ConstCxArg_I32, ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_UNIFORM_4F, "gl.Uniform4f", opGlUniform4f, In(ConstCxArg_I32, ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_UNIFORM_1I, "gl.Uniform1i", opGlUniform1i, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_UNIFORM_2I, "gl.Uniform2i", opGlUniform2i, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_UNIFORM_3I, "gl.Uniform3i", opGlUniform3i, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_UNIFORM_4I, "gl.Uniform4i", opGlUniform4i, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_UNIFORM_1FV, "gl.Uniform1fv", opGlUniform1fv, In(ConstCxArg_I32, ConstCxArg_F32, ast.Slice(constants.TYPE_F32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_2FV, "gl.Uniform2fv", opGlUniform2fv, In(ConstCxArg_I32, ConstCxArg_F32, ast.Slice(constants.TYPE_F32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_3FV, "gl.Uniform3fv", opGlUniform3fv, In(ConstCxArg_I32, ConstCxArg_I32, ast.Slice(constants.TYPE_F32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_4FV, "gl.Uniform4fv", opGlUniform4fv, In(ConstCxArg_I32, ConstCxArg_I32, ast.Slice(constants.TYPE_F32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_1IV, "gl.Uniform1iv", opGlUniform1iv, In(ConstCxArg_I32, ConstCxArg_I32, ast.Slice(constants.TYPE_I32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_2IV, "gl.Uniform2iv", opGlUniform2iv, In(ConstCxArg_I32, ConstCxArg_I32, ast.Slice(constants.TYPE_I32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_3IV, "gl.Uniform3iv", opGlUniform3iv, In(ConstCxArg_I32, ConstCxArg_I32, ast.Slice(constants.TYPE_I32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_4IV, "gl.Uniform4iv", opGlUniform4iv, In(ConstCxArg_I32, ConstCxArg_I32, ast.Slice(constants.TYPE_I32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_MATRIX_2FV, "gl.UniformMatrix2fv", opGlUniformMatrix2fv, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_BOOL, ast.Slice(constants.TYPE_F32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_MATRIX_3FV, "gl.UniformMatrix3fv", opGlUniformMatrix3fv, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_BOOL, ast.Slice(constants.TYPE_F32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_MATRIX_4FV, "gl.UniformMatrix4fv", opGlUniformMatrix4fv, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_BOOL, ast.Slice(constants.TYPE_F32)), nil)
	ast.Op_V2(OP_GL_UNIFORM_V4F, "gl.UniformV4F", opGlUniformV4F, In(ConstCxArg_I32, AUND), nil)
	ast.Op_V2(OP_GL_UNIFORM_M44F, "gl.UniformM44F", opGlUniformM44F, In(ConstCxArg_I32, ConstCxArg_BOOL, AUND), nil)
	ast.Op_V2(OP_GL_UNIFORM_M44FV, "gl.UniformM44FV", opGlUniformM44FV, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_BOOL, ast.Slice(constants.TYPE_UNDEFINED)), nil)
	ast.Op_V2(OP_GL_VERTEX_ATTRIB_POINTER, "gl.VertexAttribPointer", opGlVertexAttribPointer, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_BOOL, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_VERTEX_ATTRIB_POINTER_I32, "gl.VertexAttribPointerI32", opGlVertexAttribPointerI32, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_BOOL, ConstCxArg_I32, ConstCxArg_I32), nil)

	// gl_3_0
	ast.Op_V2(OP_GL_CLEAR_BUFFER_I, "gl.ClearBufferI", opGlClearBufferI, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_CLEAR_BUFFER_UI, "gl.ClearBufferUI", opGlClearBufferUI, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_UI32, ConstCxArg_UI32, ConstCxArg_UI32, ConstCxArg_UI32), nil)
	ast.Op_V2(OP_GL_CLEAR_BUFFER_F, "gl.ClearBufferF", opGlClearBufferF, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32, ConstCxArg_F32), nil)
	ast.Op_V2(OP_GL_BIND_RENDERBUFFER, "gl.BindRenderbuffer", opGlBindRenderbuffer, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_DELETE_RENDERBUFFERS, "gl.DeleteRenderbuffers", opGlDeleteRenderbuffers, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_GEN_RENDERBUFFERS, "gl.GenRenderbuffers", opGlGenRenderbuffers, In(ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_RENDERBUFFER_STORAGE, "gl.RenderbufferStorage", opGlRenderbufferStorage, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_BIND_FRAMEBUFFER, "gl.BindFramebuffer", opGlBindFramebuffer, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_DELETE_FRAMEBUFFERS, "gl.DeleteFramebuffers", opGlDeleteFramebuffers, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_GEN_FRAMEBUFFERS, "gl.GenFramebuffers", opGlGenFramebuffers, In(ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_CHECK_FRAMEBUFFER_STATUS, "gl.CheckFramebufferStatus", opGlCheckFramebufferStatus, In(ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_FRAMEBUFFER_TEXTURE_2D, "gl.FramebufferTexture2D", opGlFramebufferTexture2D, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_FRAMEBUFFER_RENDERBUFFER, "gl.FramebufferRenderbuffer", opGlFramebufferRenderbuffer, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_GENERATE_MIPMAP, "gl.GenerateMipmap", opGlGenerateMipmap, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_BIND_VERTEX_ARRAY, "gl.BindVertexArray", opGlBindVertexArray, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_BIND_VERTEX_ARRAY_CORE, "gl.BindVertexArrayCore", opGlBindVertexArrayCore, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_DELETE_VERTEX_ARRAYS, "gl.DeleteVertexArrays", opGlDeleteVertexArrays, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_DELETE_VERTEX_ARRAYS_CORE, "gl.DeleteVertexArraysCore", opGlDeleteVertexArraysCore, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GL_GEN_VERTEX_ARRAYS, "gl.GenVertexArrays", opGlGenVertexArrays, In(ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GL_GEN_VERTEX_ARRAYS_CORE, "gl.GenVertexArraysCore", opGlGenVertexArraysCore, In(ConstCxArg_I32, ConstCxArg_I32), Out(ConstCxArg_I32))

	// goglfw
	ast.Op_V2(OP_GLFW_FULLSCREEN, "glfw.Fullscreen", opGlfwFullscreen, In(ASTR, ConstCxArg_BOOL, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)

	// glfw
	ast.Op_V2(OP_GLFW_INIT, "glfw.Init", opGlfwInit, nil, nil)
	ast.Op_V2(OP_GLFW_WINDOW_HINT, "glfw.WindowHint", opGlfwWindowHint, In(ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GLFW_CREATE_WINDOW, "glfw.CreateWindow", opGlfwCreateWindow, In(ASTR, ConstCxArg_I32, ConstCxArg_I32, ASTR), nil)
	ast.Op_V2(OP_GLFW_MAKE_CONTEXT_CURRENT, "glfw.MakeContextCurrent", opGlfwMakeContextCurrent, In(ASTR), nil)
	ast.Op_V2(OP_GLFW_SHOULD_CLOSE, "glfw.ShouldClose", opGlfwShouldClose, In(ASTR), Out(ConstCxArg_BOOL))
	ast.Op_V2(OP_GLFW_SET_SHOULD_CLOSE, "glfw.SetShouldClose", opGlfwSetShouldClose, In(ASTR, ConstCxArg_BOOL), nil)
	ast.Op_V2(OP_GLFW_POLL_EVENTS, "glfw.PollEvents", opGlfwPollEvents, nil, nil)
	ast.Op_V2(OP_GLFW_SWAP_BUFFERS, "glfw.SwapBuffers", opGlfwSwapBuffers, In(ASTR), nil)
	ast.Op_V2(OP_GLFW_GET_FRAMEBUFFER_SIZE, "glfw.GetFramebufferSize", opGlfwGetFramebufferSize, In(ASTR), Out(ConstCxArg_I32, ConstCxArg_I32))
	ast.Op_V2(OP_GLFW_GET_WINDOW_POS, "glfw.GetWindowPos", opGlfwGetWindowPos, In(ASTR), Out(ConstCxArg_I32, ConstCxArg_I32))
	ast.Op_V2(OP_GLFW_GET_WINDOW_SIZE, "glfw.GetWindowSize", opGlfwGetWindowSize, In(ASTR), Out(ConstCxArg_I32, ConstCxArg_I32))
	ast.Op_V2(OP_GLFW_GET_TIME, "glfw.GetTime", opGlfwGetTime, nil, Out(ConstCxArg_F64))
	ast.Op_V2(OP_GLFW_SWAP_INTERVAL, "glfw.SwapInterval", opGlfwSwapInterval, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_GLFW_SET_START_CALLBACK, "glfw.SetStartCallback", opGlfwSetStartCallback, In(ASTR, ASTR, ASTR), nil)
	ast.Op_V2(OP_GLFW_SET_STOP_CALLBACK, "glfw.SetStopCallback", opGlfwSetStopCallback, In(ASTR, ASTR, ASTR), nil)
	ast.Op_V2(OP_GLFW_SET_KEYBOARD_CALLBACK, "glfw.SetKeyboardCallback", opGlfwSetKeyboardCallback, In(ASTR, ASTR, ASTR), nil)
	ast.Op_V2(OP_GLFW_SET_MOUSE_CALLBACK, "glfw.SetMouseCallback", opGlfwSetMouseCallback, In(ASTR, ASTR, ASTR), nil)
	ast.Op_V2(OP_GLFW_SET_FRAMEBUFFER_SIZE_CALLBACK, "glfw.SetFramebufferSizeCallback", opGlfwSetFramebufferSizeCallback, In(ASTR, ASTR, ASTR), nil)
	ast.Op_V2(OP_GLFW_SET_WINDOW_POS_CALLBACK, "glfw.SetWindowPosCallback", opGlfwSetWindowPosCallback, In(ASTR, ASTR, ASTR), nil)
	ast.Op_V2(OP_GLFW_SET_WINDOW_SIZE_CALLBACK, "glfw.SetWindowSizeCallback", opGlfwSetWindowSizeCallback, In(ASTR, ASTR, ASTR), nil)
	ast.Op_V2(OP_GLFW_SET_KEY_CALLBACK, "glfw.SetKeyCallback", opGlfwSetKeyCallback, In(ASTR, ASTR), nil)                          // TODO : to deprecate
	ast.Op_V2(OP_GLFW_SET_MOUSE_BUTTON_CALLBACK, "glfw.SetMouseButtonCallback", opGlfwSetMouseButtonCallback, In(ASTR, ASTR), nil) // TODO : to deprecate
	ast.Op_V2(OP_GLFW_SET_CURSOR_POS_CALLBACK, "glfw.SetCursorPosCallback", opGlfwSetCursorPosCallback, In(ASTR, ASTR), nil)       // TODO : to deprecate
	ast.Op_V2(OP_GLFW_GET_CURSOR_POS, "glfw.GetCursorPos", opGlfwGetCursorPos, In(ASTR), Out(ConstCxArg_F64, ConstCxArg_F64))
	ast.Op_V2(OP_GLFW_SET_INPUT_MODE, "glfw.SetInputMode", opGlfwSetInputMode, In(ASTR, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GLFW_SET_WINDOW_POS, "glfw.SetWindowPos", opGlfwSetWindowPos, In(ASTR, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GLFW_GET_KEY, "glfw.GetKey", opGlfwGetKey, In(ASTR, ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GLFW_FUNC_I32_I32, "glfw.func_i32_i32", opGlfwFuncI32I32, In(ASTR, ASTR), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GLFW_CALL_I32_I32, "glfw.call_i32_i32", opGlfwCallI32I32, In(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GLFW_GET_WINDOW_CONTENT_SCALE, "glfw.GetWindowContentScale", opGlfwGetWindowContentScale, In(ASTR), Out(ConstCxArg_I32, ConstCxArg_I32))
	ast.Op_V2(OP_GLFW_GET_MONITOR_CONTENT_SCALE, "glfw.GetMonitorContentScale", opGlfwGetMonitorContentScale, nil, Out(ConstCxArg_F32, ConstCxArg_F32))

	// gltext
	ast.Op_V2(OP_GLTEXT_LOAD_TRUE_TYPE, "gltext.LoadTrueType", opGltextLoadTrueType, In(ConstCxArg_I32, ASTR, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GLTEXT_LOAD_TRUE_TYPE_CORE, "gltext.LoadTrueTypeCore", opGltextLoadTrueTypeCore, In(ConstCxArg_I32, ASTR, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32), nil)
	ast.Op_V2(OP_GLTEXT_PRINTF, "gltext.Printf", opGltextPrintf, In(ASTR, ConstCxArg_F32, ConstCxArg_F32, ASTR), nil)
	ast.Op_V2(OP_GLTEXT_METRICS, "gltext.Metrics", opGltextMetrics, In(ASTR, ASTR), Out(ConstCxArg_I32, ConstCxArg_I32))
	ast.Op_V2(OP_GLTEXT_TEXTURE, "gltext.Texture", opGltextTexture, In(ASTR), Out(ConstCxArg_I32))
	ast.Op_V2(OP_GLTEXT_NEXT_GLYPH, "gltext.NextGlyph", opGltextNextGlyph, In(ASTR, ASTR, ConstCxArg_I32), Out(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32))
	ast.Op_V2(OP_GLTEXT_GLYPH_BOUNDS, "gltext.GlyphBounds", opGltextGlyphBounds, nil, Out(ConstCxArg_I32, ConstCxArg_I32))
	ast.Op_V2(OP_GLTEXT_GLYPH_METRICS, "gltext.GlyphMetrics", opGltextGlyphMetrics, In(ASTR, ConstCxArg_I32), Out(ConstCxArg_I32, ConstCxArg_I32))
	ast.Op_V2(OP_GLTEXT_GLYPH_INFO, "gltext.GlyphInfo", opGltextGlyphInfo, In(ASTR, ConstCxArg_I32), Out(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32))

	// goal
	ast.Op_V2(OP_AL_LOAD_WAV, "al.LoadWav", opAlLoadWav, In(ASTR), Out(ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I32, ConstCxArg_I64, ast.Slice(constants.TYPE_UI8)))

	// openal
	ast.Op_V2(OP_AL_CLOSE_DEVICE, "al.CloseDevice", opAlCloseDevice, nil, nil)
	ast.Op_V2(OP_AL_DELETE_BUFFERS, "al.DeleteBuffers", opAlDeleteBuffers, In(ConstCxArg_I32), nil)
	ast.Op_V2(OP_AL_DELETE_SOURCES, "al.DeleteSources", opAlDeleteSources, In(ast.Slice(constants.TYPE_I32)), nil)
	ast.Op_V2(OP_AL_DEVICE_ERROR, "al.DeviceError", opAlDeviceError, nil, Out(ConstCxArg_I32))
	ast.Op_V2(OP_AL_ERROR, "al.ProgramError", opAlError, nil, Out(ConstCxArg_I32))
	ast.Op_V2(OP_AL_EXTENSIONS, "al.Extensions", opAlExtensions, nil, Out(ASTR))
	ast.Op_V2(OP_AL_OPEN_DEVICE, "al.OpenDevice", opAlOpenDevice, nil, nil)
	ast.Op_V2(OP_AL_PAUSE_SOURCES, "al.PauseSources", opAlPauseSources, In(ast.Slice(constants.TYPE_I32)), nil)
	ast.Op_V2(OP_AL_PLAY_SOURCES, "al.PlaySources", opAlPlaySources, In(ast.Slice(constants.TYPE_I32)), nil)
	ast.Op_V2(OP_AL_RENDERER, "al.Renderer", opAlRenderer, nil, Out(ASTR))
	ast.Op_V2(OP_AL_REWIND_SOURCES, "al.RewindSources", opAlRewindSources, In(ast.Slice(constants.TYPE_I32)), nil)
	ast.Op_V2(OP_AL_STOP_SOURCES, "al.StopSources", opAlStopSources, In(ast.Slice(constants.TYPE_I32)), nil)
	ast.Op_V2(OP_AL_VENDOR, "al.Vendor", opAlVendor, nil, Out(ASTR))
	ast.Op_V2(OP_AL_VERSION, "al.Version", opAlVersion, nil, Out(ASTR))
	ast.Op_V2(OP_AL_GEN_BUFFERS, "al.GenBuffers", opAlGenBuffers, In(ConstCxArg_I32), Out(ast.Slice(constants.TYPE_I32)))
	ast.Op_V2(OP_AL_BUFFER_DATA, "al.BufferData", opAlBufferData, In(ConstCxArg_I32, ConstCxArg_I32, AUND, ConstCxArg_I32), nil)
	ast.Op_V2(OP_AL_GEN_SOURCES, "al.GenSources", opAlGenSources, In(ConstCxArg_I32), Out(ast.Slice(constants.TYPE_I32)))
	ast.Op_V2(OP_AL_SOURCE_BUFFERS_PROCESSED, "al.SourceBuffersProcessed", opAlSourceBuffersProcessed, In(ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_AL_SOURCE_BUFFERS_QUEUED, "al.SourceBuffersQueued", opAlSourceBuffersQueued, In(ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_AL_SOURCE_QUEUE_BUFFERS, "al.SourceQueueBuffers", opAlSourceQueueBuffers, In(ConstCxArg_I32, ast.Slice(constants.TYPE_I32)), nil)
	ast.Op_V2(OP_AL_SOURCE_STATE, "al.SourceState", opAlSourceState, In(ConstCxArg_I32), Out(ConstCxArg_I32))
	ast.Op_V2(OP_AL_SOURCE_UNQUEUE_BUFFERS, "al.SourceUnqueueBuffers", opAlSourceUnqueueBuffers, In(ConstCxArg_I32, ast.Slice(constants.TYPE_I32)), nil)
}
