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
		client: {"start":"_app/immutable/entry/start.DOpwBfWz.js","app":"_app/immutable/entry/app.CkJx5IWF.js","imports":["_app/immutable/entry/start.DOpwBfWz.js","_app/immutable/chunks/entry.CYIHfCA2.js","_app/immutable/chunks/scheduler.B4NQwY3t.js","_app/immutable/chunks/index.DtvkMZOO.js","_app/immutable/entry/app.CkJx5IWF.js","_app/immutable/chunks/index.DQqmjWy_.js","_app/immutable/chunks/scheduler.B4NQwY3t.js","_app/immutable/chunks/index.gMR2UOS6.js"],"stylesheets":[],"fonts":[],"uses_env_dynamic_public":false},
		nodes: [
			__memo(() => import('./chunks/0-D7K4aicj.js')),
			__memo(() => import('./chunks/1-Dfp_R93q.js')),
			__memo(() => import('./chunks/2-0ldhntCK.js')),
			__memo(() => import('./chunks/3-B-p69skC.js')),
			__memo(() => import('./chunks/4-CQDBkmAr.js')),
			__memo(() => import('./chunks/5-Ta2f3pwu.js')),
			__memo(() => import('./chunks/6-ovAyFb28.js')),
			__memo(() => import('./chunks/7-Lcud5ouQ.js')),
			__memo(() => import('./chunks/8-B8r9p4Un.js'))
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
				id: "/errors/404",
				pattern: /^\/errors\/404\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 6 },
				endpoint: null
			},
			{
				id: "/errors/50x",
				pattern: /^\/errors\/50x\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 7 },
				endpoint: null
			},
			{
				id: "/logout",
				pattern: /^\/logout\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 8 },
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
