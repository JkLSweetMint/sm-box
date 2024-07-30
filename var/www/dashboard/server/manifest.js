const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png"]),
	mimeTypes: {".png":"image/png"},
	_: {
		client: {"start":"_app/immutable/entry/start.BsiFymlP.js","app":"_app/immutable/entry/app.DEzL2fqw.js","imports":["_app/immutable/entry/start.BsiFymlP.js","_app/immutable/chunks/entry.B6-GTex4.js","_app/immutable/chunks/scheduler.C73wj8LU.js","_app/immutable/chunks/index.CzHE4dtV.js","_app/immutable/entry/app.DEzL2fqw.js","_app/immutable/chunks/scheduler.C73wj8LU.js","_app/immutable/chunks/index.DcrltIXL.js"],"stylesheets":[],"fonts":[],"uses_env_dynamic_public":false},
		nodes: [
			__memo(() => import('./chunks/0-DJWYmp1d.js')),
			__memo(() => import('./chunks/1-BNGPp0KP.js')),
			__memo(() => import('./chunks/2-CS0f-vFl.js')),
			__memo(() => import('./chunks/3-Br6eDLYc.js')),
			__memo(() => import('./chunks/4-DABL1rvC.js')),
			__memo(() => import('./chunks/5-CfCvmLw8.js')),
			__memo(() => import('./chunks/6-CIHRUtsH.js')),
			__memo(() => import('./chunks/7-Dm7W5xDG.js')),
			__memo(() => import('./chunks/8-DWLD7m36.js')),
			__memo(() => import('./chunks/9-Dra34D3N.js'))
		],
		routes: [
			{
				id: "/(app)",
				pattern: /^\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 3 },
				endpoint: null
			},
			{
				id: "/auth",
				pattern: /^\/auth\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 5 },
				endpoint: null
			},
			{
				id: "/errors/403",
				pattern: /^\/errors\/403\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 6 },
				endpoint: null
			},
			{
				id: "/errors/404",
				pattern: /^\/errors\/404\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 7 },
				endpoint: null
			},
			{
				id: "/errors/50x",
				pattern: /^\/errors\/50x\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 8 },
				endpoint: null
			},
			{
				id: "/logout",
				pattern: /^\/logout\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 9 },
				endpoint: null
			},
			{
				id: "/(app)/system/urls",
				pattern: /^\/system\/urls\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 4 },
				endpoint: null
			}
		],
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();

const prerendered = new Set([]);

const base = "";

export { base, manifest, prerendered };
//# sourceMappingURL=manifest.js.map
