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
		client: {"start":"_app/immutable/entry/start.C3HWTZcA.js","app":"_app/immutable/entry/app.IeF3-vgK.js","imports":["_app/immutable/entry/start.C3HWTZcA.js","_app/immutable/chunks/entry.DvJ_tHr6.js","_app/immutable/chunks/scheduler.B5rZiKyw.js","_app/immutable/chunks/index.C_NBoP9Y.js","_app/immutable/entry/app.IeF3-vgK.js","_app/immutable/chunks/scheduler.B5rZiKyw.js","_app/immutable/chunks/index.DrETqFkR.js"],"stylesheets":[],"fonts":[],"uses_env_dynamic_public":false},
		nodes: [
			__memo(() => import('./chunks/0-D_XcMbwN.js')),
			__memo(() => import('./chunks/1-oHqr12Gk.js')),
			__memo(() => import('./chunks/2-C8naDka-.js')),
			__memo(() => import('./chunks/3-DYaZWJ6s.js')),
			__memo(() => import('./chunks/4-CbihjtMP.js')),
			__memo(() => import('./chunks/5-MyY-7flQ.js')),
			__memo(() => import('./chunks/6-B1VApCKr.js')),
			__memo(() => import('./chunks/7-Cin9Kpbw.js')),
			__memo(() => import('./chunks/8-C5-sIFNB.js')),
			__memo(() => import('./chunks/9-CtETAkrJ.js'))
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
