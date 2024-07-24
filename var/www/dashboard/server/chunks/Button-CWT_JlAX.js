import { c as create_ssr_component, b as add_attribute, e as escape, f as null_to_empty } from './ssr-0mLmEaQb.js';
import { c as cn } from './index-VQC3TRid.js';

const css = {
  code: ".btn.svelte-16w947d{position:relative;display:flex;flex-direction:row;flex-wrap:wrap;align-items:center;justify-content:center;gap:0.5rem;overflow:hidden;padding-top:0.5rem;padding-bottom:0.5rem;padding-left:1rem;padding-right:1rem;border-radius:var(--rounded-btn);-webkit-user-select:none;-moz-user-select:none;user-select:none;background-color:rgb(var(--btn-background));color:rgb(var(--btn-foreground));outline:2px solid transparent;outline-offset:2px;transition-property:all;transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1);transition-duration:150ms;transition-duration:var(--animation-btn)}.btn.svelte-16w947d:active{--tw-scale-x:.95;--tw-scale-y:.95;transform:translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))}.btn.svelte-16w947d:disabled{background-color:rgb(var(--neutral) / 0.2);color:rgb(var(--basec) / 0.2)}.btn.circle.svelte-16w947d{border-radius:9999px}.btn.plain.svelte-16w947d{background-color:transparent;--tw-text-opacity:1;color:rgb(var(--basec) / var(--tw-text-opacity))}.btn.plain.svelte-16w947d:active{background-color:rgb(var(--btn-background)/.2);color:rgb(var(--btn-background))}.btn.no-animation.svelte-16w947d:active{--tw-scale-x:1;--tw-scale-y:1;transform:translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))}.btn.svelte-16w947d>.ripple{transition:all 1s ease-out;will-change:transition;pointer-events:none;position:absolute;height:0.25rem;width:0.25rem;border-radius:9999px;--tw-bg-opacity:1;background-color:rgb(255 255 255 / var(--tw-bg-opacity));opacity:0.4}.btn.svelte-16w947d>.ripple.run{height:12rem;width:12rem;--tw-translate-x:-6rem;--tw-translate-y:-6rem;transform:translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y));opacity:0.1}.btn.svelte-16w947d:focus-visible{--btn-focus-shadow:0 0 0 0.25rem rgba(var(--btn-shadow-color) / .3);box-shadow:var(--btn-focus-shadow)}.btn.primary.svelte-16w947d{--btn-background:var(--primary);--btn-foreground:var(--primary-fg);--btn-shadow-color:var(--primary)}.btn.secondary.svelte-16w947d{--btn-background:var(--secondary);--btn-foreground:var(--secondary-fg);--btn-shadow-color:var(--secondary)}.btn.neutral.svelte-16w947d{--btn-background:var(--neutral);--btn-foreground:var(--neutral-fg);--btn-shadow-color:var(--neutral)}.btn.success.svelte-16w947d{--btn-background:var(--success);--btn-foreground:var(--success-fg);--btn-shadow-color:var(--success)}.btn.info.svelte-16w947d{--btn-background:var(--info);--btn-foreground:var(--info-fg);--btn-shadow-color:var(--info)}.btn.warning.svelte-16w947d{--btn-background:var(--warning);--btn-foreground:var(--warning-fg);--btn-shadow-color:var(--warning)}.btn.error.svelte-16w947d{--btn-background:var(--error);--btn-foreground:var(--error-fg);--btn-shadow-color:var(--error)}.btn.smoke.svelte-16w947d{--btn-background:var(--smoke);--btn-foreground:var(--smoke-fg);--btn-shadow-color:var(--smoke)}",
  map: '{"version":3,"file":"Button.svelte","sources":["Button.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">export const NO_ANIMATION = 1;\\nexport const NO_RIPPLE = 2;\\nexport const PLAIN = 4;\\nexport const SUBMIT = 8;\\nexport const RESET = 16;\\nexport const CIRCLE = 32;\\nexport const PRIMARY = \\"primary\\";\\nexport const SECONDARY = \\"secondary\\";\\nexport const NEUTRAL = \\"neutral\\";\\nexport const SUCCESS = \\"success\\";\\nexport const INFO = \\"info\\";\\nexport const WARNING = \\"warning\\";\\nexport const ERROR = \\"error\\";\\nexport const SMOKE = \\"smoke\\";\\nexport const GHOST = \\"\\";\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { cn } from \\"@/lib/helpers\\";\\nimport { onMount } from \\"svelte\\";\\nlet root;\\nconst SpawnRipple = function(e) {\\n  if (flags & NO_RIPPLE || !root) return;\\n  const rippleEl = document.createElement(\\"span\\");\\n  rippleEl.classList.add(\\"ripple\\");\\n  const x = e.offsetX;\\n  const y = e.offsetY;\\n  rippleEl.style.left = `${x}px`;\\n  rippleEl.style.top = `${y}px`;\\n  root.appendChild(rippleEl);\\n  setTimeout(() => {\\n    rippleEl.classList.add(\\"run\\");\\n    rippleEl.addEventListener(\\"transitionend\\", () => rippleEl.remove());\\n  });\\n};\\nconst GetButtonType = function(flags2) {\\n  if (flags2 & SUBMIT) return \\"submit\\";\\n  if (flags2 & RESET) return \\"reset\\";\\n  return \\"button\\";\\n};\\nlet className = \\"\\";\\nexport { className as class };\\nexport let flags = 0;\\nexport let palette = PRIMARY;\\nexport let style = \\"\\";\\nexport let disabled = false;\\nexport let OnClick = () => {\\n};\\nonMount(() => {\\n  if (!root) return;\\n  root.removeEventListener(\\"mousedown\\", SpawnRipple);\\n  root.addEventListener(\\"mousedown\\", SpawnRipple);\\n});\\n<\/script>\\r\\n\\r\\n<button \\r\\n    on:click={OnClick} \\r\\n    type={GetButtonType(flags)} \\r\\n    class={\\"btn\\" + cn(flags & NO_ANIMATION ? \\"no-animation\\" : \\"\\") + cn(palette) + cn(flags & PLAIN ? \\"plain\\" : \\"\\") + cn(flags & CIRCLE ? \\"circle\\" : \\"\\") + cn(className)} \\r\\n    disabled={disabled}\\r\\n    style={style} \\r\\n    bind:this={root}\\r\\n>\\r\\n    <slot />\\r\\n</button>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .btn {\\r\\n            position: relative;\\r\\n            display: flex;\\r\\n            flex-direction: row;\\r\\n            flex-wrap: wrap;\\r\\n            align-items: center;\\r\\n            justify-content: center;\\r\\n            gap: 0.5rem;\\r\\n            overflow: hidden;\\r\\n            padding-top: 0.5rem;\\r\\n            padding-bottom: 0.5rem;\\r\\n            padding-left: 1rem;\\r\\n            padding-right: 1rem;\\r\\n            border-radius: var(--rounded-btn);\\r\\n            -webkit-user-select: none;\\r\\n               -moz-user-select: none;\\r\\n                    user-select: none;\\r\\n            background-color: rgb(var(--btn-background));\\r\\n            color: rgb(var(--btn-foreground));\\r\\n            outline: 2px solid transparent;\\r\\n            outline-offset: 2px;\\r\\n            transition-property: all;\\r\\n            transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\\r\\n            transition-duration: 150ms;\\r\\n            transition-duration: var(--animation-btn);\\r\\n}\\r\\n    .btn:active {\\r\\n            --tw-scale-x: .95;\\r\\n            --tw-scale-y: .95;\\r\\n            transform: translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y));\\r\\n}\\r\\n\\r\\n        .btn:disabled {\\r\\n            background-color: rgb(var(--neutral) / 0.2);\\r\\n            color: rgb(var(--basec) / 0.2);\\r\\n}\\r\\n\\r\\n        .btn.circle {\\r\\n            border-radius: 9999px;\\r\\n}\\r\\n\\r\\n        .btn.plain {\\r\\n            background-color: transparent;\\r\\n            --tw-text-opacity: 1;\\r\\n            color: rgb(var(--basec) / var(--tw-text-opacity));\\r\\n}\\r\\n\\r\\n        .btn.plain:active {\\r\\n            background-color: rgb(var(--btn-background)/.2);\\r\\n            color: rgb(var(--btn-background));\\r\\n}\\r\\n\\r\\n        .btn.no-animation:active {\\r\\n            --tw-scale-x: 1;\\r\\n            --tw-scale-y: 1;\\r\\n            transform: translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y));\\r\\n}\\r\\n\\r\\n        .btn > :global(.ripple) {\\r\\n            transition: all 1s ease-out;\\r\\n            will-change: transition;\\r\\n            pointer-events: none;\\r\\n            position: absolute;\\r\\n            height: 0.25rem;\\r\\n            width: 0.25rem;\\r\\n            border-radius: 9999px;\\r\\n            --tw-bg-opacity: 1;\\r\\n            background-color: rgb(255 255 255 / var(--tw-bg-opacity));\\r\\n            opacity: 0.4;\\r\\n        }\\r\\n\\r\\n        .btn > :global(.ripple.run) {\\r\\n            height: 12rem;\\r\\n            width: 12rem;\\r\\n            --tw-translate-x: -6rem;\\r\\n            --tw-translate-y: -6rem;\\r\\n            transform: translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y));\\r\\n            opacity: 0.1;\\r\\n}\\r\\n\\r\\n        .btn:focus-visible {\\r\\n            --btn-focus-shadow: 0 0 0 0.25rem rgba(var(--btn-shadow-color) / .3);\\r\\n\\r\\n            box-shadow: var(--btn-focus-shadow);\\r\\n        }\\r\\n\\r\\n        .btn.primary {\\r\\n            --btn-background: var(--primary);\\r\\n            --btn-foreground: var(--primary-fg);\\r\\n            --btn-shadow-color: var(--primary);\\r\\n        }\\r\\n\\r\\n        .btn.secondary {\\r\\n            --btn-background: var(--secondary);\\r\\n            --btn-foreground: var(--secondary-fg);\\r\\n            --btn-shadow-color: var(--secondary);\\r\\n        }\\r\\n\\r\\n        .btn.neutral {\\r\\n            --btn-background: var(--neutral);\\r\\n            --btn-foreground: var(--neutral-fg);\\r\\n            --btn-shadow-color: var(--neutral);\\r\\n        }\\r\\n\\r\\n        .btn.success {\\r\\n            --btn-background: var(--success);\\r\\n            --btn-foreground: var(--success-fg);\\r\\n            --btn-shadow-color: var(--success);\\r\\n        }\\r\\n\\r\\n        .btn.info {\\r\\n            --btn-background: var(--info);\\r\\n            --btn-foreground: var(--info-fg);\\r\\n            --btn-shadow-color: var(--info);\\r\\n        }\\r\\n\\r\\n        .btn.warning {\\r\\n            --btn-background: var(--warning);\\r\\n            --btn-foreground: var(--warning-fg);\\r\\n            --btn-shadow-color: var(--warning);\\r\\n        }\\r\\n\\r\\n        .btn.error {\\r\\n            --btn-background: var(--error);\\r\\n            --btn-foreground: var(--error-fg);\\r\\n            --btn-shadow-color: var(--error);\\r\\n        }\\r\\n\\r\\n        .btn.smoke {\\r\\n            --btn-background: var(--smoke);\\r\\n            --btn-foreground: var(--smoke-fg);\\r\\n            --btn-shadow-color: var(--smoke);\\r\\n        }\\r\\n</style>"],"names":[],"mappings":"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"}'
};
const NO_ANIMATION = 1;
const NO_RIPPLE = 2;
const PLAIN = 4;
const SUBMIT = 8;
const RESET = 16;
const CIRCLE = 32;
const PRIMARY = "primary";
const SMOKE = "smoke";
const GHOST = "";
const Button = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let root;
  const GetButtonType = function(flags2) {
    if (flags2 & SUBMIT) return "submit";
    if (flags2 & RESET) return "reset";
    return "button";
  };
  let { class: className = "" } = $$props;
  let { flags = 0 } = $$props;
  let { palette = PRIMARY } = $$props;
  let { style = "" } = $$props;
  let { disabled = false } = $$props;
  let { OnClick = () => {
  } } = $$props;
  if ($$props.class === void 0 && $$bindings.class && className !== void 0) $$bindings.class(className);
  if ($$props.flags === void 0 && $$bindings.flags && flags !== void 0) $$bindings.flags(flags);
  if ($$props.palette === void 0 && $$bindings.palette && palette !== void 0) $$bindings.palette(palette);
  if ($$props.style === void 0 && $$bindings.style && style !== void 0) $$bindings.style(style);
  if ($$props.disabled === void 0 && $$bindings.disabled && disabled !== void 0) $$bindings.disabled(disabled);
  if ($$props.OnClick === void 0 && $$bindings.OnClick && OnClick !== void 0) $$bindings.OnClick(OnClick);
  $$result.css.add(css);
  return `<button${add_attribute("type", GetButtonType(flags), 0)} class="${escape(null_to_empty("btn" + cn(flags & NO_ANIMATION ? "no-animation" : "") + cn(palette) + cn(flags & PLAIN ? "plain" : "") + cn(flags & CIRCLE ? "circle" : "") + cn(className)), true) + " svelte-16w947d"}" ${disabled ? "disabled" : ""}${add_attribute("style", style, 0)}${add_attribute("this", root, 0)}>${slots.default ? slots.default({}) : ``} </button>`;
});

export { Button as B, CIRCLE as C, GHOST as G, NO_RIPPLE as N, PLAIN as P, SUBMIT as S, NO_ANIMATION as a, SMOKE as b };
//# sourceMappingURL=Button-CWT_JlAX.js.map
