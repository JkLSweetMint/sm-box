import { c as create_ssr_component, g as getContext, v as validate_component, h as spread, i as escape_object, b as add_attribute, d as each, e as escape, s as subscribe, j as escape_attribute_value } from './ssr-C-9IsUTH.js';
import { t as toWritableStores, o as omit, a as overridable, w as withGet, g as generateIds, m as makeElement, i as isBrowser, s as styleToString, p as portalAttr, l as tick, q as isElement, e as executeCallbacks, b as addMeltEventListener, k as kbd, u as usePortal, r as dequal, v as isFunction, c as isHTMLElement, j as createElHelpers, n as noop } from './create-DPwALkVX.js';
import { d as derivedVisible, e as effect, f as usePopper, g as getPortalDestination, r as removeScroll, s as sleep } from './action-DplsO8dc.js';
import { w as writable } from './index2-DbZx0BBT.js';
import { c as cn } from './index-VQC3TRid.js';
import { B as Button, N as NO_RIPPLE, P as PLAIN, C as CIRCLE, a as NO_ANIMATION, G as GHOST } from './Button-atP5DRSs.js';

async function handleFocus(args) {
  const { prop, defaultEl } = args;
  await Promise.all([sleep(1), tick]);
  if (prop === void 0) {
    defaultEl?.focus();
    return;
  }
  const returned = isFunction(prop) ? prop(defaultEl) : prop;
  if (typeof returned === "string") {
    const el = document.querySelector(returned);
    if (!isHTMLElement(el))
      return;
    el.focus();
  } else if (isHTMLElement(returned)) {
    returned.focus();
  }
}
function melt(node, params) {
  throw new Error("[MELTUI ERROR]: The `use:melt` action cannot be used without MeltUI's Preprocessor. See: https://www.melt-ui.com/docs/preprocessor");
}
const defaults = {
  positioning: {
    placement: "bottom"
  },
  arrowSize: 8,
  defaultOpen: false,
  disableFocusTrap: false,
  escapeBehavior: "close",
  preventScroll: false,
  onOpenChange: void 0,
  closeOnOutsideClick: true,
  portal: "body",
  forceVisible: false,
  openFocus: void 0,
  closeFocus: void 0,
  onOutsideClick: void 0,
  preventTextSelectionOverflow: true
};
const { name } = createElHelpers("popover");
const popoverIdParts = ["trigger", "content"];
function createPopover(args) {
  const withDefaults = { ...defaults, ...args };
  const options = toWritableStores(omit(withDefaults, "open", "ids"));
  const { positioning, arrowSize, disableFocusTrap, preventScroll, escapeBehavior, closeOnOutsideClick, portal, forceVisible, openFocus, closeFocus, onOutsideClick, preventTextSelectionOverflow } = options;
  const openWritable = withDefaults.open ?? writable(withDefaults.defaultOpen);
  const open = overridable(openWritable, withDefaults?.onOpenChange);
  const activeTrigger = withGet.writable(null);
  const ids = toWritableStores({ ...generateIds(popoverIdParts), ...withDefaults.ids });
  function handleClose() {
    open.set(false);
  }
  const isVisible = derivedVisible({ open, activeTrigger, forceVisible });
  const content = makeElement(name("content"), {
    stores: [isVisible, open, activeTrigger, portal, ids.content],
    returned: ([$isVisible, $open, $activeTrigger, $portal, $contentId]) => {
      return {
        hidden: $isVisible && isBrowser ? void 0 : true,
        tabindex: -1,
        style: $isVisible ? void 0 : styleToString({ display: "none" }),
        id: $contentId,
        "data-state": $open && $activeTrigger ? "open" : "closed",
        "data-portal": portalAttr($portal)
      };
    },
    action: (node) => {
      let unsubPopper = noop;
      const unsubDerived = effect([isVisible, activeTrigger, positioning, disableFocusTrap, closeOnOutsideClick, portal], ([$isVisible, $activeTrigger, $positioning, $disableFocusTrap, $closeOnOutsideClick, $portal]) => {
        unsubPopper();
        if (!$isVisible || !$activeTrigger)
          return;
        tick().then(() => {
          unsubPopper();
          unsubPopper = usePopper(node, {
            anchorElement: $activeTrigger,
            open,
            options: {
              floating: $positioning,
              focusTrap: $disableFocusTrap ? null : void 0,
              modal: {
                shouldCloseOnInteractOutside,
                onClose: handleClose,
                closeOnInteractOutside: $closeOnOutsideClick
              },
              escapeKeydown: { behaviorType: escapeBehavior },
              portal: getPortalDestination(node, $portal),
              preventTextSelectionOverflow: { enabled: preventTextSelectionOverflow }
            }
          }).destroy;
        });
      });
      return {
        destroy() {
          unsubDerived();
          unsubPopper();
        }
      };
    }
  });
  async function toggleOpen() {
    open.update((prev) => !prev);
  }
  function shouldCloseOnInteractOutside(e) {
    onOutsideClick.get()?.(e);
    if (e.defaultPrevented)
      return false;
    const target = e.target;
    const triggerEl = document.getElementById(ids.trigger.get());
    if (triggerEl && isElement(target)) {
      if (target === triggerEl || triggerEl.contains(target))
        return false;
    }
    return true;
  }
  const trigger = makeElement(name("trigger"), {
    stores: [isVisible, ids.content, ids.trigger],
    returned: ([$isVisible, $contentId, $triggerId]) => {
      return {
        role: "button",
        "aria-haspopup": "dialog",
        "aria-expanded": $isVisible ? "true" : "false",
        "data-state": stateAttr($isVisible),
        "aria-controls": $contentId,
        id: $triggerId
      };
    },
    action: (node) => {
      activeTrigger.set(node);
      const unsub = executeCallbacks(addMeltEventListener(node, "click", toggleOpen), addMeltEventListener(node, "keydown", (e) => {
        if (e.key !== kbd.ENTER && e.key !== kbd.SPACE)
          return;
        e.preventDefault();
        toggleOpen();
      }));
      return {
        destroy() {
          activeTrigger.set(null);
          unsub();
        }
      };
    }
  });
  const overlay = makeElement(name("overlay"), {
    stores: [isVisible],
    returned: ([$isVisible]) => {
      return {
        hidden: $isVisible ? void 0 : true,
        tabindex: -1,
        style: styleToString({
          display: $isVisible ? void 0 : "none"
        }),
        "aria-hidden": "true",
        "data-state": stateAttr($isVisible)
      };
    },
    action: (node) => {
      let unsubDerived = noop;
      let unsubPortal = noop;
      unsubDerived = effect([portal], ([$portal]) => {
        unsubPortal();
        if ($portal === null)
          return;
        const portalDestination = getPortalDestination(node, $portal);
        if (portalDestination === null)
          return;
        unsubPortal = usePortal(node, portalDestination).destroy;
      });
      return {
        destroy() {
          unsubDerived();
          unsubPortal();
        }
      };
    }
  });
  const arrow = makeElement(name("arrow"), {
    stores: arrowSize,
    returned: ($arrowSize) => ({
      "data-arrow": true,
      style: styleToString({
        position: "absolute",
        width: `var(--arrow-size, ${$arrowSize}px)`,
        height: `var(--arrow-size, ${$arrowSize}px)`
      })
    })
  });
  const close = makeElement(name("close"), {
    returned: () => ({
      type: "button"
    }),
    action: (node) => {
      const unsub = executeCallbacks(addMeltEventListener(node, "click", (e) => {
        if (e.defaultPrevented)
          return;
        handleClose();
      }), addMeltEventListener(node, "keydown", (e) => {
        if (e.defaultPrevented)
          return;
        if (e.key !== kbd.ENTER && e.key !== kbd.SPACE)
          return;
        e.preventDefault();
        toggleOpen();
      }));
      return {
        destroy: unsub
      };
    }
  });
  effect([open, activeTrigger, preventScroll], ([$open, $activeTrigger, $preventScroll]) => {
    if (!isBrowser || !$open)
      return;
    const unsubs = [];
    if ($preventScroll) {
      unsubs.push(removeScroll());
    }
    handleFocus({ prop: openFocus.get(), defaultEl: $activeTrigger });
    return () => {
      unsubs.forEach((unsub) => unsub());
    };
  });
  effect(open, ($open) => {
    if (!isBrowser || $open)
      return;
    const triggerEl = document.getElementById(ids.trigger.get());
    handleFocus({ prop: closeFocus.get(), defaultEl: triggerEl });
  }, { skipFirstRun: true });
  return {
    ids,
    elements: {
      trigger,
      content,
      arrow,
      close,
      overlay
    },
    states: {
      open
    },
    options
  };
}
function stateAttr(open) {
  return open ? "open" : "closed";
}
function keys(obj) {
  return Object.keys(obj);
}
function createSync(stores) {
  let setters = {};
  keys(stores).forEach((key) => {
    const store = stores[key];
    effect(store, (value) => {
      if (key in setters) {
        setters[key]?.(value);
      }
    });
  });
  return keys(stores).reduce((acc, key) => {
    return {
      ...acc,
      [key]: function sync(value, setter) {
        stores[key].update((p) => {
          if (dequal(p, value))
            return p;
          return value;
        });
        if (setter) {
          setters = { ...setters, [key]: setter };
        }
      }
    };
  }, {});
}
const css = {
  code: ".popover.svelte-4wje2y{z-index:var(--dropdown-z-index);border-width:1px;--tw-border-opacity:1;border-color:rgb(var(--base3) / var(--tw-border-opacity));--tw-bg-opacity:1;background-color:rgb(var(--base1) / var(--tw-bg-opacity));padding:0.5rem;--tw-shadow:0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);--tw-shadow-colored:0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);border-radius:var(--rounded-box)\n}.popover.no-border.svelte-4wje2y{border-style:none\n}.popover.no-shadow.svelte-4wje2y{--tw-shadow:0 0 #0000;--tw-shadow-colored:0 0 #0000;box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)\n}",
  map: '{"version":3,"file":"Popover.svelte","sources":["Popover.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">export const DISABLE_CLOSE_BTN = 1;\\nexport const NO_BORDER = 2;\\nexport const NO_SHADOW = 4;\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { scale, fade, fly } from \\"svelte/transition\\";\\nimport { createPopover, createSync, melt } from \\"@melt-ui/svelte\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nlet className = \\"\\";\\nexport { className as class };\\nexport let open = false;\\nexport let flags = 0;\\nconst {\\n  elements: { trigger, content, close },\\n  states\\n} = createPopover({\\n  forceVisible: true\\n});\\nconst sync = createSync(states);\\n$: sync.open(open, (v) => open = v);\\n<\/script>\\r\\n\\r\\n<slot name=\\"trigger\\" melt={melt} trigger={$trigger}>\\r\\n\\r\\n</slot>\\r\\n\\r\\n{#if open}\\r\\n  <div\\r\\n    {...$content} use:$content.action\\r\\n    transition:fly={{ duration: 300, y: 100 }}\\r\\n    class={\\"popover\\" + cn(className) + cn(flags & NO_BORDER ? \\"no-border\\" : \\"\\") + cn(flags & NO_SHADOW ? \\"no-shadow\\" : \\"\\")}\\r\\n  >\\r\\n    <slot />\\r\\n\\r\\n    {#if !(flags & DISABLE_CLOSE_BTN)}\\r\\n        <button type=\\"button\\" class=\\"close\\" {...$close} use:$close.action>\\r\\n            <i class=\\"fa-solid fa-xmark\\"></i>\\r\\n        </button>\\r\\n    {/if}\\r\\n  </div>\\r\\n{/if}\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .popover {\\r\\n\\r\\n    z-index: var(--dropdown-z-index);\\r\\n\\r\\n    border-width: 1px;\\r\\n\\r\\n    --tw-border-opacity: 1;\\r\\n\\r\\n    border-color: rgb(var(--base3) / var(--tw-border-opacity));\\r\\n\\r\\n    --tw-bg-opacity: 1;\\r\\n\\r\\n    background-color: rgb(var(--base1) / var(--tw-bg-opacity));\\r\\n\\r\\n    padding: 0.5rem;\\r\\n\\r\\n    --tw-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);\\r\\n\\r\\n    --tw-shadow-colored: 0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);\\r\\n\\r\\n    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);\\r\\n\\r\\n    border-radius: var(--rounded-box)\\n}\\r\\n\\r\\n        .popover.no-border {\\r\\n\\r\\n    border-style: none\\n}\\r\\n\\r\\n        .popover.no-shadow {\\r\\n\\r\\n    --tw-shadow: 0 0 #0000;\\r\\n\\r\\n    --tw-shadow-colored: 0 0 #0000;\\r\\n\\r\\n    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)\\n}\\r\\n</style>"],"names":[],"mappings":"AA2CI,sBAAS,CAET,OAAO,CAAE,IAAI,kBAAkB,CAAC,CAEhC,YAAY,CAAE,GAAG,CAEjB,mBAAmB,CAAE,CAAC,CAEtB,YAAY,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,mBAAmB,CAAC,CAAC,CAE1D,eAAe,CAAE,CAAC,CAElB,gBAAgB,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,eAAe,CAAC,CAAC,CAE1D,OAAO,CAAE,MAAM,CAEf,WAAW,CAAE,kEAAkE,CAE/E,mBAAmB,CAAE,8EAA8E,CAEnG,UAAU,CAAE,IAAI,uBAAuB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,gBAAgB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,WAAW,CAAC,CAEvG,aAAa,CAAE,IAAI,aAAa,CAAC;AACrC,CAEQ,QAAQ,wBAAW,CAEvB,YAAY,CAAE,IAAI;AACtB,CAEQ,QAAQ,wBAAW,CAEvB,WAAW,CAAE,SAAS,CAEtB,mBAAmB,CAAE,SAAS,CAE9B,UAAU,CAAE,IAAI,uBAAuB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,gBAAgB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,WAAW,CAAC;AAC3G"}'
};
const DISABLE_CLOSE_BTN = 1;
const NO_BORDER = 2;
const NO_SHADOW = 4;
const Popover = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $trigger, $$unsubscribe_trigger;
  let $content, $$unsubscribe_content;
  let $close, $$unsubscribe_close;
  let { class: className = "" } = $$props;
  let { open = false } = $$props;
  let { flags = 0 } = $$props;
  const { elements: { trigger, content, close }, states } = createPopover({ forceVisible: true });
  $$unsubscribe_trigger = subscribe(trigger, (value) => $trigger = value);
  $$unsubscribe_content = subscribe(content, (value) => $content = value);
  $$unsubscribe_close = subscribe(close, (value) => $close = value);
  const sync = createSync(states);
  if ($$props.class === void 0 && $$bindings.class && className !== void 0) $$bindings.class(className);
  if ($$props.open === void 0 && $$bindings.open && open !== void 0) $$bindings.open(open);
  if ($$props.flags === void 0 && $$bindings.flags && flags !== void 0) $$bindings.flags(flags);
  $$result.css.add(css);
  {
    sync.open(open, (v) => open = v);
  }
  $$unsubscribe_trigger();
  $$unsubscribe_content();
  $$unsubscribe_close();
  return `${slots.trigger ? slots.trigger({ melt, trigger: $trigger }) : ` `} ${open ? `<div${spread(
    [
      escape_object($content),
      {
        class: escape_attribute_value("popover" + cn(className) + cn(flags & NO_BORDER ? "no-border" : "") + cn(flags & NO_SHADOW ? "no-shadow" : ""))
      }
    ],
    { classes: "svelte-4wje2y" }
  )}>${slots.default ? slots.default({}) : ``} ${!(flags & DISABLE_CLOSE_BTN) ? `<button${spread([{ type: "button" }, { class: "close" }, escape_object($close)], { classes: "svelte-4wje2y" })}><i class="fa-solid fa-xmark"></i></button>` : ``}</div>` : ``}`;
});
const LanguageSelect = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let countries;
  let current = "";
  const callServiceMethod = getContext("CallServiceMethod");
  const ChangeLanguage = async function(code) {
    if (code == current) return;
    let resp = await callServiceMethod({
      service: "i18n",
      url: "/languages/set",
      headers: {
        "Content-Type": "application/json;charset=utf-8"
      },
      method: "POST",
      data: { code },
      showToast: true
    });
    if (resp && resp.code == 200) window.location.reload();
  };
  let { class: className = "" } = $$props;
  let { size = "24px" } = $$props;
  const GetCountryCode = function(code) {
    return (code.split("-").slice(-1)[0] ?? "").toLowerCase();
  };
  if ($$props.class === void 0 && $$bindings.class && className !== void 0) $$bindings.class(className);
  if ($$props.size === void 0 && $$bindings.size && size !== void 0) $$bindings.size(size);
  return `${validate_component(Popover, "Popover").$$render(
    $$result,
    {
      class: "!p-0 overflow-hidden",
      flags: DISABLE_CLOSE_BTN
    },
    {},
    {
      trigger: ({ melt: melt2, trigger }) => {
        return `<div${spread([{ slot: "trigger" }, escape_object(trigger)], {})}>${validate_component(Button, "Button").$$render(
          $$result,
          {
            class: "!p-1 w-max h-max" + cn(className),
            flags: NO_RIPPLE | PLAIN | CIRCLE
          },
          {},
          {
            default: () => {
              return `${`<div class="rounded bg-black/20 animate-pulse duration-500 m-1"${add_attribute("style", `width: ${size}; height: ${size};`, 0)}></div>`}`;
            }
          }
        )}</div>`;
      },
      default: () => {
        return `<div class="max-h-80 overflow-y-auto"><div class="flex flex-col w-40">${each(countries, (country, idx) => {
          let cc = GetCountryCode(country.code);
          return ` ${validate_component(Button, "Button").$$render(
            $$result,
            {
              class: "hover:!bg-black/5 focus-visible:!bg-black/5 !rounded-none",
              flags: NO_ANIMATION | NO_RIPPLE,
              palette: GHOST,
              disabled: !country.active,
              OnClick: () => ChangeLanguage(country.code)
            },
            {},
            {
              default: () => {
                return `<span${add_attribute("class", `fi fi-${cc} fis rounded !w-5 !h-5 saturate-[1.25]` + cn(!country.active ? "saturate-[.25]" : ""), 0)}></span> <span class="mr-auto max-w-24 break-words">${escape(country.name)}</span> `;
              }
            }
          )}`;
        })}</div></div>`;
      }
    }
  )}`;
});

export { DISABLE_CLOSE_BTN as D, LanguageSelect as L, Popover as P, handleFocus as h };
//# sourceMappingURL=LanguageSelect-BU0ksdbn.js.map
