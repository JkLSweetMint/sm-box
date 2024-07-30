import { c as create_ssr_component, e as escape, f as null_to_empty, b as add_attribute } from './ssr-C-9IsUTH.js';
import { c as cn } from './index-VQC3TRid.js';

const css = {
  code: ".ds-card.svelte-1gsiull.svelte-1gsiull{position:relative;border-width:1px;--tw-border-opacity:1;border-color:rgb(var(--base3) / var(--tw-border-opacity));--tw-bg-opacity:1;background-color:rgb(var(--base1) / var(--tw-bg-opacity));--tw-shadow:0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);--tw-shadow-colored:0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);border-radius:var(--rounded-box)\n}.ds-card.no-border.svelte-1gsiull.svelte-1gsiull{border-style:none\n}.ds-card.no-shadow.svelte-1gsiull.svelte-1gsiull{--tw-shadow:0 0 #0000;--tw-shadow-colored:0 0 #0000;box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)\n}.ds-card.svelte-1gsiull>.header.svelte-1gsiull{display:flex;flex-direction:column;gap:0.5rem;padding-left:1rem;padding-right:1rem;padding-top:1.5rem\n}.ds-card.svelte-1gsiull>.header.svelte-1gsiull:empty{padding:0px;display:none\n}.ds-card.svelte-1gsiull>.body.svelte-1gsiull{display:flex;flex-direction:column;overflow-y:auto;padding:1rem\n}.ds-card.svelte-1gsiull>.body.svelte-1gsiull:empty{padding:0px;display:none\n}.ds-card.svelte-1gsiull>.footer.svelte-1gsiull{display:flex;flex-direction:row;justify-content:flex-end;gap:0.5rem;padding-left:1rem;padding-right:1rem;padding-top:1.5rem;padding-bottom:0.75rem\n}.ds-card.svelte-1gsiull>.footer.svelte-1gsiull:empty{padding:0px;display:none\n}.ds-card.svelte-1gsiull>.fallback.svelte-1gsiull{position:relative;display:flex;height:100%;width:100%;flex-direction:column\n}.ds-card.svelte-1gsiull>.fallback.svelte-1gsiull:empty{display:none\n}.ds-card.compact.svelte-1gsiull>.header.svelte-1gsiull{padding-left:0.5rem;padding-right:0.5rem;padding-top:0.75rem\n}.ds-card.compact.svelte-1gsiull>.header.svelte-1gsiull:empty{padding:0px\n}.ds-card.compact.svelte-1gsiull>.body.svelte-1gsiull{padding:0.5rem\n}.ds-card.compact.svelte-1gsiull>.body.svelte-1gsiull:empty{padding:0px\n}.ds-card.compact.svelte-1gsiull>.footer.svelte-1gsiull{padding-left:0.5rem;padding-right:0.5rem;padding-top:0.75rem\n}.ds-card.compact.svelte-1gsiull>.footer.svelte-1gsiull:empty{padding:0px\n}",
  map: '{"version":3,"file":"Card.svelte","sources":["Card.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">export const NO_BORDER = 1;\\nexport const NO_SHADOW = 2;\\nexport const COMPACT = 4;\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { cn } from \\"@/lib/helpers\\";\\nlet className = \\"\\";\\nexport { className as class };\\nexport let flags = 0;\\nexport let style = \\"\\";\\n<\/script>\\r\\n\\r\\n<div \\r\\n    class={\\"ds-card\\" + cn(className) + cn(flags & NO_BORDER ? \\"no-border\\" : \\"\\") + cn(flags & NO_SHADOW ? \\"no-shadow\\" : \\"\\") + cn(flags & COMPACT ? \\"compact\\" : \\"\\")} \\r\\n    style={style}\\r\\n>\\r\\n    <div class=\\"header\\">\\r\\n        <slot name=\\"header\\" />\\r\\n    </div>\\r\\n    <div class=\\"body\\">\\r\\n        <slot />\\r\\n    </div>\\r\\n    <div class=\\"footer\\">\\r\\n        <slot name=\\"footer\\" />\\r\\n    </div>\\r\\n    <div class=\\"fallback\\">\\r\\n        <slot name=\\"fallback\\" />\\r\\n    </div>\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .ds-card {\\r\\n\\r\\n    position: relative;\\r\\n\\r\\n    border-width: 1px;\\r\\n\\r\\n    --tw-border-opacity: 1;\\r\\n\\r\\n    border-color: rgb(var(--base3) / var(--tw-border-opacity));\\r\\n\\r\\n    --tw-bg-opacity: 1;\\r\\n\\r\\n    background-color: rgb(var(--base1) / var(--tw-bg-opacity));\\r\\n\\r\\n    --tw-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);\\r\\n\\r\\n    --tw-shadow-colored: 0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);\\r\\n\\r\\n    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);\\r\\n\\r\\n    border-radius: var(--rounded-box)\\n}\\r\\n\\r\\n        .ds-card.no-border {\\r\\n\\r\\n    border-style: none\\n}\\r\\n\\r\\n        .ds-card.no-shadow {\\r\\n\\r\\n    --tw-shadow: 0 0 #0000;\\r\\n\\r\\n    --tw-shadow-colored: 0 0 #0000;\\r\\n\\r\\n    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)\\n}\\r\\n\\r\\n        .ds-card > .header {\\r\\n\\r\\n    display: flex;\\r\\n\\r\\n    flex-direction: column;\\r\\n\\r\\n    gap: 0.5rem;\\r\\n\\r\\n    padding-left: 1rem;\\r\\n\\r\\n    padding-right: 1rem;\\r\\n\\r\\n    padding-top: 1.5rem\\n}\\r\\n\\r\\n        .ds-card > .header:empty {\\r\\n\\r\\n    padding: 0px;\\r\\n\\r\\n    display: none\\n}\\r\\n\\r\\n        .ds-card > .body {\\r\\n\\r\\n    display: flex;\\r\\n\\r\\n    flex-direction: column;\\r\\n\\r\\n    overflow-y: auto;\\r\\n\\r\\n    padding: 1rem\\n}\\r\\n\\r\\n        .ds-card > .body:empty {\\r\\n\\r\\n    padding: 0px;\\r\\n\\r\\n    display: none\\n}\\r\\n\\r\\n        .ds-card > .footer {\\r\\n\\r\\n    display: flex;\\r\\n\\r\\n    flex-direction: row;\\r\\n\\r\\n    justify-content: flex-end;\\r\\n\\r\\n    gap: 0.5rem;\\r\\n\\r\\n    padding-left: 1rem;\\r\\n\\r\\n    padding-right: 1rem;\\r\\n\\r\\n    padding-top: 1.5rem;\\r\\n\\r\\n    padding-bottom: 0.75rem\\n}\\r\\n\\r\\n        .ds-card > .footer:empty {\\r\\n\\r\\n    padding: 0px;\\r\\n\\r\\n    display: none\\n}\\r\\n\\r\\n        .ds-card > .fallback {\\r\\n\\r\\n    position: relative;\\r\\n\\r\\n    display: flex;\\r\\n\\r\\n    height: 100%;\\r\\n\\r\\n    width: 100%;\\r\\n\\r\\n    flex-direction: column\\n}\\r\\n\\r\\n        .ds-card > .fallback:empty {\\r\\n\\r\\n    display: none\\n}\\r\\n\\r\\n        .ds-card.compact > .header {\\r\\n\\r\\n    padding-left: 0.5rem;\\r\\n\\r\\n    padding-right: 0.5rem;\\r\\n\\r\\n    padding-top: 0.75rem\\n}\\r\\n\\r\\n        .ds-card.compact > .header:empty {\\r\\n\\r\\n    padding: 0px\\n}\\r\\n\\r\\n        .ds-card.compact > .body {\\r\\n\\r\\n    padding: 0.5rem\\n}\\r\\n\\r\\n        .ds-card.compact > .body:empty {\\r\\n\\r\\n    padding: 0px\\n}\\r\\n\\r\\n        .ds-card.compact > .footer {\\r\\n\\r\\n    padding-left: 0.5rem;\\r\\n\\r\\n    padding-right: 0.5rem;\\r\\n\\r\\n    padding-top: 0.75rem\\n}\\r\\n\\r\\n        .ds-card.compact > .footer:empty {\\r\\n\\r\\n    padding: 0px\\n}\\r\\n</style>"],"names":[],"mappings":"AA+BI,sCAAS,CAET,QAAQ,CAAE,QAAQ,CAElB,YAAY,CAAE,GAAG,CAEjB,mBAAmB,CAAE,CAAC,CAEtB,YAAY,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,mBAAmB,CAAC,CAAC,CAE1D,eAAe,CAAE,CAAC,CAElB,gBAAgB,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,eAAe,CAAC,CAAC,CAE1D,WAAW,CAAE,kEAAkE,CAE/E,mBAAmB,CAAE,8EAA8E,CAEnG,UAAU,CAAE,IAAI,uBAAuB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,gBAAgB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,WAAW,CAAC,CAEvG,aAAa,CAAE,IAAI,aAAa,CAAC;AACrC,CAEQ,QAAQ,wCAAW,CAEvB,YAAY,CAAE,IAAI;AACtB,CAEQ,QAAQ,wCAAW,CAEvB,WAAW,CAAE,SAAS,CAEtB,mBAAmB,CAAE,SAAS,CAE9B,UAAU,CAAE,IAAI,uBAAuB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,gBAAgB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,WAAW,CAAC;AAC3G,CAEQ,uBAAQ,CAAG,sBAAQ,CAEvB,OAAO,CAAE,IAAI,CAEb,cAAc,CAAE,MAAM,CAEtB,GAAG,CAAE,MAAM,CAEX,YAAY,CAAE,IAAI,CAElB,aAAa,CAAE,IAAI,CAEnB,WAAW,CAAE,MAAM;AACvB,CAEQ,uBAAQ,CAAG,sBAAO,MAAO,CAE7B,OAAO,CAAE,GAAG,CAEZ,OAAO,CAAE,IAAI;AACjB,CAEQ,uBAAQ,CAAG,oBAAM,CAErB,OAAO,CAAE,IAAI,CAEb,cAAc,CAAE,MAAM,CAEtB,UAAU,CAAE,IAAI,CAEhB,OAAO,CAAE,IAAI;AACjB,CAEQ,uBAAQ,CAAG,oBAAK,MAAO,CAE3B,OAAO,CAAE,GAAG,CAEZ,OAAO,CAAE,IAAI;AACjB,CAEQ,uBAAQ,CAAG,sBAAQ,CAEvB,OAAO,CAAE,IAAI,CAEb,cAAc,CAAE,GAAG,CAEnB,eAAe,CAAE,QAAQ,CAEzB,GAAG,CAAE,MAAM,CAEX,YAAY,CAAE,IAAI,CAElB,aAAa,CAAE,IAAI,CAEnB,WAAW,CAAE,MAAM,CAEnB,cAAc,CAAE,OAAO;AAC3B,CAEQ,uBAAQ,CAAG,sBAAO,MAAO,CAE7B,OAAO,CAAE,GAAG,CAEZ,OAAO,CAAE,IAAI;AACjB,CAEQ,uBAAQ,CAAG,wBAAU,CAEzB,QAAQ,CAAE,QAAQ,CAElB,OAAO,CAAE,IAAI,CAEb,MAAM,CAAE,IAAI,CAEZ,KAAK,CAAE,IAAI,CAEX,cAAc,CAAE,MAAM;AAC1B,CAEQ,uBAAQ,CAAG,wBAAS,MAAO,CAE/B,OAAO,CAAE,IAAI;AACjB,CAEQ,QAAQ,uBAAQ,CAAG,sBAAQ,CAE/B,YAAY,CAAE,MAAM,CAEpB,aAAa,CAAE,MAAM,CAErB,WAAW,CAAE,OAAO;AACxB,CAEQ,QAAQ,uBAAQ,CAAG,sBAAO,MAAO,CAErC,OAAO,CAAE,GAAG;AAChB,CAEQ,QAAQ,uBAAQ,CAAG,oBAAM,CAE7B,OAAO,CAAE,MAAM;AACnB,CAEQ,QAAQ,uBAAQ,CAAG,oBAAK,MAAO,CAEnC,OAAO,CAAE,GAAG;AAChB,CAEQ,QAAQ,uBAAQ,CAAG,sBAAQ,CAE/B,YAAY,CAAE,MAAM,CAEpB,aAAa,CAAE,MAAM,CAErB,WAAW,CAAE,OAAO;AACxB,CAEQ,QAAQ,uBAAQ,CAAG,sBAAO,MAAO,CAErC,OAAO,CAAE,GAAG;AAChB"}'
};
const NO_BORDER = 1;
const NO_SHADOW = 2;
const COMPACT = 4;
const Card = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { class: className = "" } = $$props;
  let { flags = 0 } = $$props;
  let { style = "" } = $$props;
  if ($$props.class === void 0 && $$bindings.class && className !== void 0) $$bindings.class(className);
  if ($$props.flags === void 0 && $$bindings.flags && flags !== void 0) $$bindings.flags(flags);
  if ($$props.style === void 0 && $$bindings.style && style !== void 0) $$bindings.style(style);
  $$result.css.add(css);
  return `<div class="${escape(null_to_empty("ds-card" + cn(className) + cn(flags & NO_BORDER ? "no-border" : "") + cn(flags & NO_SHADOW ? "no-shadow" : "") + cn(flags & COMPACT ? "compact" : "")), true) + " svelte-1gsiull"}"${add_attribute("style", style, 0)}><div class="header svelte-1gsiull">${slots.header ? slots.header({}) : ``}</div> <div class="body svelte-1gsiull">${slots.default ? slots.default({}) : ``}</div> <div class="footer svelte-1gsiull">${slots.footer ? slots.footer({}) : ``}</div> <div class="fallback svelte-1gsiull">${slots.fallback ? slots.fallback({}) : ``}</div> </div>`;
});

export { COMPACT as C, Card as a };
//# sourceMappingURL=Card-BWlCOFJT.js.map
