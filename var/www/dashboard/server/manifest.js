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
		client: {"start":"_app/immutable/entry/start.BNJSDRIh.js","app":"_app/immutable/entry/app.DVOrWyDu.js","imports":["_app/immutable/entry/start.BNJSDRIh.js","_app/immutable/chunks/entry.BtcyXRKn.js","_app/immutable/chunks/scheduler.VUuLSjLj.js","_app/immutable/chunks/index.oxanXEPY.js","_app/immutable/entry/app.DVOrWyDu.js","_app/immutable/chunks/scheduler.VUuLSjLj.js","_app/immutable/chunks/index.YcojCQiS.js"],"stylesheets":[],"fonts":[],"uses_env_dynamic_public":false},
		nodes: [
			__memo(() => import('./chunks/0-CyFsXgZt.js')),
			__memo(() => import('./chunks/1-BX5mAPtW.js')),
			__memo(() => import('./chunks/2-CR6WVYfP.js')),
			__memo(() => import('./chunks/3-DO7jw3Gt.js')),
			__memo(() => import('./chunks/4-C5dAoAwW.js')),
			__memo(() => import('./chunks/5-DBjvdFKm.js')),
			__memo(() => import('./chunks/6-CplzULsX.js'))
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
