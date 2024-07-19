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
		client: {"start":"_app/immutable/entry/start.qgv7Hyoe.js","app":"_app/immutable/entry/app.CGDcnyLv.js","imports":["_app/immutable/entry/start.qgv7Hyoe.js","_app/immutable/chunks/entry.DWVxLkBy.js","_app/immutable/chunks/scheduler.VUuLSjLj.js","_app/immutable/chunks/index.oxanXEPY.js","_app/immutable/entry/app.CGDcnyLv.js","_app/immutable/chunks/index.DQqmjWy_.js","_app/immutable/chunks/scheduler.VUuLSjLj.js","_app/immutable/chunks/index.YcojCQiS.js"],"stylesheets":[],"fonts":[],"uses_env_dynamic_public":false},
		nodes: [
			__memo(() => import('./chunks/0-BeIHb2z4.js')),
			__memo(() => import('./chunks/1-CCde3qI-.js')),
			__memo(() => import('./chunks/2-CqqzDA2L.js')),
			__memo(() => import('./chunks/3-Cc6i96AG.js')),
			__memo(() => import('./chunks/4-BEr6-E6P.js')),
			__memo(() => import('./chunks/5-DMn2NR_G.js')),
			__memo(() => import('./chunks/6-CJ6bcRXt.js')),
			__memo(() => import('./chunks/7-DRLelazJ.js'))
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
				page: { layouts: [0,], errors: [1,], leaf: 4 },
				endpoint: null
			},
			{
				id: "/errors/403",
				pattern: /^\/errors\/403\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 5 },
				endpoint: null
			},
			{
				id: "/errors/50x",
				pattern: /^\/errors\/50x\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 6 },
				endpoint: null
			},
			{
				id: "/logout",
				pattern: /^\/logout\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 7 },
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
