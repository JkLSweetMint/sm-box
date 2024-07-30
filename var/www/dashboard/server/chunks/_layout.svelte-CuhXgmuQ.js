import { c as create_ssr_component, v as validate_component, e as escape, s as subscribe, g as getContext, a as setContext, b as add_attribute, d as each, f as null_to_empty, h as spread, i as escape_object, j as escape_attribute_value, k as set_store_value, l as compute_rest_props, m as get_store_value, o as onDestroy } from './ssr-C-9IsUTH.js';
import { r as readable, w as writable, d as derived, a as readonly } from './index2-DbZx0BBT.js';
import { c as cn } from './index-VQC3TRid.js';
import { t as toWritableStores$2, o as omit$2, w as withGet, g as generateIds$1, a as overridable$1, m as makeElement, e as executeCallbacks$2, b as addMeltEventListener$1, k as kbd$2, s as styleToString$3, p as portalAttr, n as noop$1, u as usePortal$1, i as isBrowser$2, d as disabledAttr, c as isHTMLElement$2, f as getDirectionalKeys, h as nanoid, j as createElHelpers$1, l as tick } from './create-CrsTuLZT.js';
import { L as LanguageSelect, P as Popover, D as DISABLE_CLOSE_BTN, e as effect$2, u as useModal, a as useEscapeKeydown$1, b as useFocusTrap, g as getPortalDestination$1, r as removeScroll$1, h as handleFocus$1, n as next, p as prev, l as last$1, c as createFocusTrap$1 } from './LanguageSelect-DH6Fy-Xq.js';
import { B as Button, G as GHOST, N as NO_RIPPLE, a as NO_ANIMATION, P as PLAIN, C as CIRCLE } from './Button-atP5DRSs.js';
import { n as navigating } from './stores-D5UcsfBi.js';
import { T as Text } from './Text-DvReL212.js';
import './exports-BGi7-Rnc.js';

function getElemDirection(elem) {
  const style = window.getComputedStyle(elem);
  const direction = style.getPropertyValue("direction");
  return direction;
}
const { name: name$2 } = createElHelpers$1("dialog");
const defaults$3 = {
  preventScroll: true,
  escapeBehavior: "close",
  closeOnOutsideClick: true,
  role: "dialog",
  defaultOpen: false,
  portal: "body",
  forceVisible: false,
  openFocus: void 0,
  closeFocus: void 0,
  onOutsideClick: void 0
};
const dialogIdParts$1 = ["content", "title", "description"];
function createDialog$1(props) {
  const withDefaults = { ...defaults$3, ...props };
  const options = toWritableStores$2(omit$2(withDefaults, "ids"));
  const { preventScroll, escapeBehavior, closeOnOutsideClick, role, portal, forceVisible, openFocus, closeFocus, onOutsideClick } = options;
  const activeTrigger = withGet.writable(null);
  const ids = toWritableStores$2({
    ...generateIds$1(dialogIdParts$1),
    ...withDefaults.ids
  });
  const openWritable = withDefaults.open ?? writable(withDefaults.defaultOpen);
  const open = overridable$1(openWritable, withDefaults?.onOpenChange);
  const isVisible = derived([open, forceVisible], ([$open, $forceVisible]) => {
    return $open || $forceVisible;
  });
  let unsubScroll = noop$1;
  function handleOpen(e) {
    const el = e.currentTarget;
    const triggerEl = e.currentTarget;
    if (!isHTMLElement$2(el) || !isHTMLElement$2(triggerEl))
      return;
    open.set(true);
    activeTrigger.set(triggerEl);
  }
  function handleClose() {
    open.set(false);
  }
  const trigger = makeElement(name$2("trigger"), {
    stores: [open],
    returned: ([$open]) => {
      return {
        "aria-haspopup": "dialog",
        "aria-expanded": $open,
        type: "button"
      };
    },
    action: (node) => {
      const unsub = executeCallbacks$2(addMeltEventListener$1(node, "click", (e) => {
        handleOpen(e);
      }), addMeltEventListener$1(node, "keydown", (e) => {
        if (e.key !== kbd$2.ENTER && e.key !== kbd$2.SPACE)
          return;
        e.preventDefault();
        handleOpen(e);
      }));
      return {
        destroy: unsub
      };
    }
  });
  const overlay = makeElement(name$2("overlay"), {
    stores: [isVisible, open],
    returned: ([$isVisible, $open]) => {
      return {
        hidden: $isVisible ? void 0 : true,
        tabindex: -1,
        style: $isVisible ? void 0 : styleToString$3({ display: "none" }),
        "aria-hidden": true,
        "data-state": $open ? "open" : "closed"
      };
    }
  });
  const content = makeElement(name$2("content"), {
    stores: [isVisible, ids.content, ids.description, ids.title, open],
    returned: ([$isVisible, $contentId, $descriptionId, $titleId, $open]) => {
      return {
        id: $contentId,
        role: role.get(),
        "aria-describedby": $descriptionId,
        "aria-labelledby": $titleId,
        "aria-modal": $isVisible ? "true" : void 0,
        "data-state": $open ? "open" : "closed",
        tabindex: -1,
        hidden: $isVisible ? void 0 : true,
        style: $isVisible ? void 0 : styleToString$3({ display: "none" })
      };
    },
    action: (node) => {
      let unsubEscape = noop$1;
      let unsubModal = noop$1;
      let unsubFocusTrap = noop$1;
      const unsubDerived = effect$2([isVisible, closeOnOutsideClick], ([$isVisible, $closeOnOutsideClick]) => {
        unsubModal();
        unsubEscape();
        unsubFocusTrap();
        if (!$isVisible)
          return;
        unsubModal = useModal(node, {
          closeOnInteractOutside: $closeOnOutsideClick,
          onClose: handleClose,
          shouldCloseOnInteractOutside(e) {
            onOutsideClick.get()?.(e);
            if (e.defaultPrevented)
              return false;
            return true;
          }
        }).destroy;
        unsubEscape = useEscapeKeydown$1(node, {
          handler: handleClose,
          behaviorType: escapeBehavior
        }).destroy;
        unsubFocusTrap = useFocusTrap(node, { fallbackFocus: node }).destroy;
      });
      return {
        destroy: () => {
          unsubScroll();
          unsubDerived();
          unsubModal();
          unsubEscape();
          unsubFocusTrap();
        }
      };
    }
  });
  const portalled = makeElement(name$2("portalled"), {
    stores: portal,
    returned: ($portal) => ({
      "data-portal": portalAttr($portal)
    }),
    action: (node) => {
      const unsubPortal = effect$2([portal], ([$portal]) => {
        if ($portal === null)
          return noop$1;
        const portalDestination = getPortalDestination$1(node, $portal);
        if (portalDestination === null)
          return noop$1;
        return usePortal$1(node, portalDestination).destroy;
      });
      return {
        destroy() {
          unsubPortal();
        }
      };
    }
  });
  const title = makeElement(name$2("title"), {
    stores: [ids.title],
    returned: ([$titleId]) => ({
      id: $titleId
    })
  });
  const description = makeElement(name$2("description"), {
    stores: [ids.description],
    returned: ([$descriptionId]) => ({
      id: $descriptionId
    })
  });
  const close = makeElement(name$2("close"), {
    returned: () => ({
      type: "button"
    }),
    action: (node) => {
      const unsub = executeCallbacks$2(addMeltEventListener$1(node, "click", () => {
        handleClose();
      }), addMeltEventListener$1(node, "keydown", (e) => {
        if (e.key !== kbd$2.SPACE && e.key !== kbd$2.ENTER)
          return;
        e.preventDefault();
        handleClose();
      }));
      return {
        destroy: unsub
      };
    }
  });
  effect$2([open, preventScroll], ([$open, $preventScroll]) => {
    if (!isBrowser$2)
      return;
    if ($preventScroll && $open)
      unsubScroll = removeScroll$1();
    if ($open) {
      const contentEl = document.getElementById(ids.content.get());
      handleFocus$1({ prop: openFocus.get(), defaultEl: contentEl });
    }
    return () => {
      if (!forceVisible.get()) {
        unsubScroll();
      }
    };
  });
  effect$2(open, ($open) => {
    if (!isBrowser$2 || $open)
      return;
    handleFocus$1({
      prop: closeFocus.get(),
      defaultEl: activeTrigger.get()
    });
  }, { skipFirstRun: true });
  return {
    ids,
    elements: {
      content,
      trigger,
      title,
      description,
      overlay,
      close,
      portalled
    },
    states: {
      open
    },
    options
  };
}
const defaults$2 = {
  orientation: "horizontal",
  activateOnFocus: true,
  loop: true,
  autoSet: true
};
const { name: name$1, selector } = createElHelpers$1("tabs");
function createTabs(props) {
  const withDefaults = { ...defaults$2, ...props };
  const options = toWritableStores$2(omit$2(withDefaults, "defaultValue", "value", "onValueChange", "autoSet"));
  const { orientation, activateOnFocus, loop } = options;
  const valueWritable = withDefaults.value ?? writable(withDefaults.defaultValue);
  const value = overridable$1(valueWritable, withDefaults?.onValueChange);
  let ssrValue = withDefaults.defaultValue ?? value.get();
  const root = makeElement(name$1(), {
    stores: orientation,
    returned: ($orientation) => {
      return {
        "data-orientation": $orientation
      };
    }
  });
  const list = makeElement(name$1("list"), {
    stores: orientation,
    returned: ($orientation) => {
      return {
        role: "tablist",
        "aria-orientation": $orientation,
        "data-orientation": $orientation
      };
    }
  });
  const parseTriggerProps = (props2) => {
    if (typeof props2 === "string") {
      return { value: props2 };
    } else {
      return props2;
    }
  };
  const trigger = makeElement(name$1("trigger"), {
    stores: [value, orientation],
    returned: ([$value, $orientation]) => {
      return (props2) => {
        const { value: tabValue, disabled } = parseTriggerProps(props2);
        if (!$value && !ssrValue && withDefaults.autoSet) {
          ssrValue = tabValue;
          $value = tabValue;
          value.set(tabValue);
        }
        const sourceOfTruth = isBrowser$2 ? $value : ssrValue;
        const isActive = sourceOfTruth === tabValue;
        return {
          type: "button",
          role: "tab",
          "data-state": isActive ? "active" : "inactive",
          tabindex: isActive ? 0 : -1,
          "data-value": tabValue,
          "data-orientation": $orientation,
          "data-disabled": disabledAttr(disabled),
          disabled: disabledAttr(disabled)
        };
      };
    },
    action: (node) => {
      const unsub = executeCallbacks$2(addMeltEventListener$1(node, "focus", () => {
        const disabled = node.dataset.disabled === "true";
        const tabValue = node.dataset.value;
        if (activateOnFocus.get() && !disabled && tabValue !== void 0) {
          value.set(tabValue);
        }
      }), addMeltEventListener$1(node, "click", (e) => {
        node.focus();
        e.preventDefault();
        const disabled = node.dataset.disabled === "true";
        if (disabled)
          return;
        const tabValue = node.dataset.value;
        node.focus();
        if (tabValue !== void 0) {
          value.set(tabValue);
        }
      }), addMeltEventListener$1(node, "keydown", (e) => {
        const tabValue = node.dataset.value;
        if (!tabValue)
          return;
        const el = e.currentTarget;
        if (!isHTMLElement$2(el))
          return;
        const rootEl = el.closest(selector());
        if (!isHTMLElement$2(rootEl))
          return;
        const $loop = loop.get();
        const triggers = Array.from(rootEl.querySelectorAll('[role="tab"]')).filter((trigger2) => isHTMLElement$2(trigger2));
        const enabledTriggers = triggers.filter((el2) => !el2.hasAttribute("data-disabled"));
        const triggerIdx = enabledTriggers.findIndex((el2) => el2 === e.target);
        const dir = getElemDirection(rootEl);
        const { nextKey, prevKey } = getDirectionalKeys(dir, orientation.get());
        if (e.key === nextKey) {
          e.preventDefault();
          const nextEl = next(enabledTriggers, triggerIdx, $loop);
          nextEl.focus();
        } else if (e.key === prevKey) {
          e.preventDefault();
          const prevEl = prev(enabledTriggers, triggerIdx, $loop);
          prevEl.focus();
        } else if (e.key === kbd$2.ENTER || e.key === kbd$2.SPACE) {
          e.preventDefault();
          value.set(tabValue);
        } else if (e.key === kbd$2.HOME) {
          e.preventDefault();
          const firstTrigger = enabledTriggers[0];
          firstTrigger.focus();
        } else if (e.key === kbd$2.END) {
          e.preventDefault();
          const lastTrigger = last$1(enabledTriggers);
          lastTrigger.focus();
        }
      }));
      return {
        destroy: unsub
      };
    }
  });
  const content = makeElement(name$1("content"), {
    stores: value,
    returned: ($value) => {
      return (tabValue) => {
        return {
          role: "tabpanel",
          // TODO: improve
          "aria-labelledby": tabValue,
          hidden: isBrowser$2 ? $value === tabValue ? void 0 : true : ssrValue === tabValue ? void 0 : true,
          tabindex: 0
        };
      };
    }
  });
  return {
    elements: {
      root,
      list,
      trigger,
      content
    },
    states: {
      value
    },
    options
  };
}
const css$6 = {
  code: '.page__navigation__list__link.svelte-aelmyg.svelte-aelmyg{--icon-size:25px;display:flex;width:100%;align-items:center;justify-content:center;gap:0.5rem;overflow:hidden;padding-left:0.5rem;padding-right:0.5rem;padding-top:0.75rem;padding-bottom:0.75rem}.page__navigation__list__link.nested.svelte-aelmyg.svelte-aelmyg{position:relative}.page__navigation__list__link.full.svelte-aelmyg.svelte-aelmyg{justify-content:flex-start}.page__navigation__list__link.svelte-aelmyg>.icon.svelte-aelmyg{display:flex;height:100%;max-height:var(--icon-size);width:100%;max-width:var(--icon-size);align-items:center;justify-content:center}.page__navigation__list__link.svelte-aelmyg>.indicator.svelte-aelmyg{position:absolute;top:0px;bottom:0px;left:0px;width:3px;--tw-bg-opacity:1;background-color:rgb(var(--neutral-fg) / var(--tw-bg-opacity))}.page__navigation__list__link.svelte-aelmyg>.indicator.svelte-aelmyg:where([dir="rtl"], [dir="rtl"] *){right:0px;left:auto}.page__navigation__list__link.svelte-aelmyg.svelte-aelmyg:hover{background-color:rgb(0 0 0 / 0.1)}.page__navigation__list__link.svelte-aelmyg.svelte-aelmyg:focus{background-color:rgb(0 0 0 / 0.1);outline:2px solid transparent;outline-offset:2px}',
  map: '{"version":3,"file":"NavigationLink.svelte","sources":["NavigationLink.svelte"],"sourcesContent":["<script lang=\\"ts\\">import { getContext } from \\"svelte\\";\\nimport { writable } from \\"svelte/store\\";\\nimport { fade } from \\"svelte/transition\\";\\nimport {} from \\"./Navigation.svelte\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nconst toggleMenu = getContext(\\"toggleMenu\\");\\nconst toggleSubMenu = writable(false);\\nexport let data;\\nexport let nested = false;\\n<\/script>\\r\\n\\r\\n{#if data}\\r\\n    {@const url = `/${data.href}`.replaceAll(\\"//\\", \\"/\\")}\\r\\n\\r\\n    {#if data.children}\\r\\n        <button \\r\\n            class=\\"page__navigation__list__link\\" \\r\\n            on:click={() => $toggleSubMenu = !$toggleSubMenu} \\r\\n            class:full={$toggleMenu}\\r\\n            class:nested={nested}\\r\\n        >\\r\\n            <div class=\\"icon\\">\\r\\n                <span class={data.icon ?? \\"fa-solid fa-circle\\"}></span>\\r\\n            </div>\\r\\n            {#if $toggleMenu}\\r\\n                <span class=\\"text-sm\\" in:fade>{data.title}</span>\\r\\n            {/if}\\r\\n            {#if $toggleMenu && data.children}\\r\\n                <span class={\\"fa-solid fa-chevron-left ml-auto rtl:mr-auto rtl:ml-2 transition-transform mr-2\\" + cn($toggleSubMenu ? \\"-rotate-90\\" : \\"\\")} in:fade></span>\\r\\n            {/if}\\r\\n            {#if nested}\\r\\n                <span class=\\"indicator\\" transition:fade></span>\\r\\n            {/if}\\r\\n        </button>\\r\\n        {#if $toggleSubMenu}\\r\\n            <div class=\\"flex flex-col w-full\\" transition:fade>\\r\\n                {#each data.children as child, idx (idx)}\\r\\n                    <svelte:self data={child} nested />\\r\\n                {/each}\\r\\n            </div>\\r\\n        {/if}\\r\\n    {:else}\\r\\n        <a class=\\"page__navigation__list__link\\" class:full={$toggleMenu} class:nested={nested} href={url} rel=\\"external\\">\\r\\n            <div class=\\"icon\\">\\r\\n                <span class={data.icon ?? \\"fa-solid fa-circle\\"}></span>\\r\\n            </div>\\r\\n            {#if $toggleMenu}\\r\\n                <span class=\\"text-sm\\" in:fade>{data.title}</span>\\r\\n            {/if}\\r\\n            {#if nested}\\r\\n                <span class=\\"indicator\\" transition:fade></span>\\r\\n            {/if}\\r\\n        </a>\\r\\n    {/if}\\r\\n{/if}\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .page__navigation__list__link {\\r\\n        --icon-size: 25px;\\r\\n        display: flex;\\r\\n        width: 100%;\\r\\n        align-items: center;\\r\\n        justify-content: center;\\r\\n        gap: 0.5rem;\\r\\n        overflow: hidden;\\r\\n        padding-left: 0.5rem;\\r\\n        padding-right: 0.5rem;\\r\\n        padding-top: 0.75rem;\\r\\n        padding-bottom: 0.75rem;\\r\\n    }\\r\\n\\r\\n        .page__navigation__list__link.nested {\\r\\n        position: relative;\\r\\n}\\r\\n\\r\\n        .page__navigation__list__link.full {\\r\\n        justify-content: flex-start;\\r\\n}\\r\\n\\r\\n        .page__navigation__list__link > .icon {\\r\\n        display: flex;\\r\\n        height: 100%;\\r\\n        max-height: var(--icon-size);\\r\\n        width: 100%;\\r\\n        max-width: var(--icon-size);\\r\\n        align-items: center;\\r\\n        justify-content: center;\\r\\n}\\r\\n\\r\\n        .page__navigation__list__link > .indicator {\\r\\n        position: absolute;\\r\\n        top: 0px;\\r\\n        bottom: 0px;\\r\\n        left: 0px;\\r\\n        width: 3px;\\r\\n        --tw-bg-opacity: 1;\\r\\n        background-color: rgb(var(--neutral-fg) / var(--tw-bg-opacity));\\r\\n}\\r\\n\\r\\n        .page__navigation__list__link > .indicator:where([dir=\\"rtl\\"], [dir=\\"rtl\\"] *) {\\r\\n        right: 0px;\\r\\n        left: auto;\\r\\n}\\r\\n\\r\\n        .page__navigation__list__link:hover {\\r\\n        background-color: rgb(0 0 0 / 0.1);\\r\\n}\\r\\n\\r\\n        .page__navigation__list__link:focus {\\r\\n        background-color: rgb(0 0 0 / 0.1);\\r\\n        outline: 2px solid transparent;\\r\\n        outline-offset: 2px;\\r\\n}\\r\\n</style>"],"names":[],"mappings":"AAyDI,yDAA8B,CAC1B,WAAW,CAAE,IAAI,CACjB,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,IAAI,CACX,WAAW,CAAE,MAAM,CACnB,eAAe,CAAE,MAAM,CACvB,GAAG,CAAE,MAAM,CACX,QAAQ,CAAE,MAAM,CAChB,YAAY,CAAE,MAAM,CACpB,aAAa,CAAE,MAAM,CACrB,WAAW,CAAE,OAAO,CACpB,cAAc,CAAE,OACpB,CAEI,6BAA6B,mCAAQ,CACrC,QAAQ,CAAE,QAClB,CAEQ,6BAA6B,iCAAM,CACnC,eAAe,CAAE,UACzB,CAEQ,2CAA6B,CAAG,mBAAM,CACtC,OAAO,CAAE,IAAI,CACb,MAAM,CAAE,IAAI,CACZ,UAAU,CAAE,IAAI,WAAW,CAAC,CAC5B,KAAK,CAAE,IAAI,CACX,SAAS,CAAE,IAAI,WAAW,CAAC,CAC3B,WAAW,CAAE,MAAM,CACnB,eAAe,CAAE,MACzB,CAEQ,2CAA6B,CAAG,wBAAW,CAC3C,QAAQ,CAAE,QAAQ,CAClB,GAAG,CAAE,GAAG,CACR,MAAM,CAAE,GAAG,CACX,IAAI,CAAE,GAAG,CACT,KAAK,CAAE,GAAG,CACV,eAAe,CAAE,CAAC,CAClB,gBAAgB,CAAE,IAAI,IAAI,YAAY,CAAC,CAAC,CAAC,CAAC,IAAI,eAAe,CAAC,CACtE,CAEQ,2CAA6B,CAAG,wBAAU,OAAO,CAAC,GAAG,CAAC,KAAK,CAAC,EAAE,CAAC,GAAG,CAAC,KAAK,CAAC,CAAC,CAAC,CAAE,CAC7E,KAAK,CAAE,GAAG,CACV,IAAI,CAAE,IACd,CAEQ,yDAA6B,MAAO,CACpC,gBAAgB,CAAE,IAAI,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,GAAG,CACzC,CAEQ,yDAA6B,MAAO,CACpC,gBAAgB,CAAE,IAAI,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,GAAG,CAAC,CAClC,OAAO,CAAE,GAAG,CAAC,KAAK,CAAC,WAAW,CAC9B,cAAc,CAAE,GACxB"}'
};
const NavigationLink = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $toggleSubMenu, $$unsubscribe_toggleSubMenu;
  let $toggleMenu, $$unsubscribe_toggleMenu;
  const toggleMenu = getContext("toggleMenu");
  $$unsubscribe_toggleMenu = subscribe(toggleMenu, (value) => $toggleMenu = value);
  const toggleSubMenu = writable(false);
  $$unsubscribe_toggleSubMenu = subscribe(toggleSubMenu, (value) => $toggleSubMenu = value);
  let { data } = $$props;
  let { nested = false } = $$props;
  if ($$props.data === void 0 && $$bindings.data && data !== void 0) $$bindings.data(data);
  if ($$props.nested === void 0 && $$bindings.nested && nested !== void 0) $$bindings.nested(nested);
  $$result.css.add(css$6);
  $$unsubscribe_toggleSubMenu();
  $$unsubscribe_toggleMenu();
  return `${data ? (() => {
    let url = `/${data.href}`.replaceAll("//", "/");
    return ` ${data.children ? `<button class="${[
      "page__navigation__list__link svelte-aelmyg",
      ($toggleMenu ? "full" : "") + " " + (nested ? "nested" : "")
    ].join(" ").trim()}"><div class="icon svelte-aelmyg"><span class="${escape(null_to_empty(data.icon ?? "fa-solid fa-circle"), true) + " svelte-aelmyg"}"></span></div> ${$toggleMenu ? `<span class="text-sm">${escape(data.title)}</span>` : ``} ${$toggleMenu && data.children ? `<span class="${escape(null_to_empty("fa-solid fa-chevron-left ml-auto rtl:mr-auto rtl:ml-2 transition-transform mr-2" + cn($toggleSubMenu ? "-rotate-90" : "")), true) + " svelte-aelmyg"}"></span>` : ``} ${nested ? `<span class="indicator svelte-aelmyg"></span>` : ``}</button> ${$toggleSubMenu ? `<div class="flex flex-col w-full">${each(data.children, (child, idx) => {
      return `${validate_component(NavigationLink, "svelte:self").$$render($$result, { data: child, nested: true }, {}, {})}`;
    })}</div>` : ``}` : `<a class="${[
      "page__navigation__list__link svelte-aelmyg",
      ($toggleMenu ? "full" : "") + " " + (nested ? "nested" : "")
    ].join(" ").trim()}"${add_attribute("href", url, 0)} rel="external"><div class="icon svelte-aelmyg"><span class="${escape(null_to_empty(data.icon ?? "fa-solid fa-circle"), true) + " svelte-aelmyg"}"></span></div> ${$toggleMenu ? `<span class="text-sm">${escape(data.title)}</span>` : ``} ${nested ? `<span class="indicator svelte-aelmyg"></span>` : ``}</a>`}`;
  })() : ``}`;
});
const css$5 = {
  code: ".notifications-block.svelte-1rz0x92{display:flex;min-width:300px;flex-direction:column;overflow:hidden\n}.notifications-block__header.svelte-1rz0x92{position:relative;--tw-bg-opacity:1;background-color:rgb(var(--neutral) / var(--tw-bg-opacity));padding-left:1rem;padding-right:1rem;padding-top:1rem;padding-bottom:3.5rem;--tw-text-opacity:1;color:rgb(var(--neutral-fg) / var(--tw-text-opacity));border-top-left-radius:var(--rounded-box);border-top-right-radius:var(--rounded-box)\n}.notifications-block__nav.svelte-1rz0x92{position:absolute;left:0.5rem;right:0.5rem;bottom:0px;display:flex;flex-shrink:0;overflow-x:auto\n}.notifications-block__trigger.svelte-1rz0x92{-webkit-user-select:none;-moz-user-select:none;user-select:none;border-top-left-radius:0.25rem;border-top-right-radius:0.25rem;padding-left:0.75rem;padding-right:0.75rem;padding-top:0.5rem;padding-bottom:0.5rem\n}.notifications-block__trigger.svelte-1rz0x92:focus{background-color:rgb(0 0 0 / 0.2);outline:2px solid transparent;outline-offset:2px\n}.notifications-block__trigger.svelte-1rz0x92:hover{background-color:rgb(0 0 0 / 0.2)\n}.notifications-block__trigger.active.svelte-1rz0x92{--tw-bg-opacity:1;background-color:rgb(var(--base1) / var(--tw-bg-opacity));--tw-text-opacity:1;color:rgb(var(--basec) / var(--tw-text-opacity))\n}.notifications-block__page.svelte-1rz0x92{max-width:300px;flex-grow:1;overflow:auto;--tw-bg-opacity:1;background-color:rgb(255 255 255 / var(--tw-bg-opacity));padding:1.25rem;outline:2px solid transparent;outline-offset:2px;border-bottom-left-radius:var(--rounded-box);border-bottom-right-radius:var(--rounded-box)\n}",
  map: `{"version":3,"file":"NotificationsBlock.svelte","sources":["NotificationsBlock.svelte"],"sourcesContent":["<script lang=\\"ts\\">import { createTabs, melt } from \\"@melt-ui/svelte\\";\\nimport { GetText, Text } from \\"@/components/i18n/Text\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nconst {\\n  elements: { root, list, content, trigger },\\n  states: { value }\\n} = createTabs({\\n  defaultValue: \\"all\\"\\n});\\nlet className = \\"\\";\\nexport { className as class };\\nexport let dictionary;\\nconst triggers = [\\n  { id: \\"all\\" },\\n  { id: \\"alerts\\" }\\n];\\n\\n\\t$: __MELTUI_BUILDER_1__ = $content('all');\\n\\t$: __MELTUI_BUILDER_2__ = $content('messages');\\n\\t$: __MELTUI_BUILDER_3__ = $content('alerts');\\n<\/script>\\r\\n\\r\\n<div class={\\"notifications-block\\" + cn(className)} {...$root} use:$root.action>\\r\\n    <div class=\\"notifications-block__header\\">\\r\\n        <span class=\\"text-xl font-semibold\\">\\r\\n            <Text key=\\"dashboard.navbar.notifications.title\\" source={dictionary} />\\r\\n        </span>\\r\\n        <div\\r\\n            {...$list} use:$list.action\\r\\n            class=\\"notifications-block__nav\\"\\r\\n        >\\r\\n            {#each triggers as triggerItem}\\r\\n                {@const __MELTUI_BUILDER_0__ = $trigger(triggerItem.id)}\\n                <button type=\\"button\\" {...__MELTUI_BUILDER_0__} use:__MELTUI_BUILDER_0__.action class=\\"notifications-block__trigger\\" class:active={$value === triggerItem.id}>\\r\\n                    {#if triggerItem.id == \\"all\\"}\\r\\n                        <Text key=\\"dashboard.navbar.notifications.tabs.all\\" source={dictionary} />\\r\\n                    {/if}\\r\\n                    {#if triggerItem.id == \\"messages\\"}\\r\\n                        <Text key=\\"dashboard.navbar.notifications.messages\\" source={dictionary} />\\r\\n                    {/if}\\r\\n                    {#if triggerItem.id == \\"alerts\\"}\\r\\n                        <Text key=\\"dashboard.navbar.notifications.tabs.alerts\\" source={dictionary} />\\r\\n                    {/if}\\r\\n                </button>\\r\\n            {/each}\\r\\n        </div>\\r\\n    </div>\\r\\n  \\r\\n    <div tabindex=\\"-1\\" {...__MELTUI_BUILDER_1__} use:__MELTUI_BUILDER_1__.action class=\\"notifications-block__page\\">\\r\\n    </div>\\r\\n    <div tabindex=\\"-1\\" {...__MELTUI_BUILDER_2__} use:__MELTUI_BUILDER_2__.action class=\\"notifications-block__page\\">\\r\\n    </div>\\r\\n    <div tabindex=\\"-1\\" {...__MELTUI_BUILDER_3__} use:__MELTUI_BUILDER_3__.action class=\\"notifications-block__page\\">\\r\\n    </div>\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .notifications-block {\\r\\n\\r\\n    display: flex;\\r\\n\\r\\n    min-width: 300px;\\r\\n\\r\\n    flex-direction: column;\\r\\n\\r\\n    overflow: hidden\\n}\\r\\n\\r\\n        .notifications-block__header {\\r\\n\\r\\n    position: relative;\\r\\n\\r\\n    --tw-bg-opacity: 1;\\r\\n\\r\\n    background-color: rgb(var(--neutral) / var(--tw-bg-opacity));\\r\\n\\r\\n    padding-left: 1rem;\\r\\n\\r\\n    padding-right: 1rem;\\r\\n\\r\\n    padding-top: 1rem;\\r\\n\\r\\n    padding-bottom: 3.5rem;\\r\\n\\r\\n    --tw-text-opacity: 1;\\r\\n\\r\\n    color: rgb(var(--neutral-fg) / var(--tw-text-opacity));\\r\\n\\r\\n    border-top-left-radius: var(--rounded-box);\\r\\n\\r\\n    border-top-right-radius: var(--rounded-box)\\n}\\r\\n\\r\\n        .notifications-block__nav {\\r\\n\\r\\n    position: absolute;\\r\\n\\r\\n    left: 0.5rem;\\r\\n\\r\\n    right: 0.5rem;\\r\\n\\r\\n    bottom: 0px;\\r\\n\\r\\n    display: flex;\\r\\n\\r\\n    flex-shrink: 0;\\r\\n\\r\\n    overflow-x: auto\\n}\\r\\n\\r\\n        .notifications-block__trigger {\\r\\n\\r\\n    -webkit-user-select: none;\\r\\n\\r\\n       -moz-user-select: none;\\r\\n\\r\\n            user-select: none;\\r\\n\\r\\n    border-top-left-radius: 0.25rem;\\r\\n\\r\\n    border-top-right-radius: 0.25rem;\\r\\n\\r\\n    padding-left: 0.75rem;\\r\\n\\r\\n    padding-right: 0.75rem;\\r\\n\\r\\n    padding-top: 0.5rem;\\r\\n\\r\\n    padding-bottom: 0.5rem\\n}\\r\\n\\r\\n        .notifications-block__trigger:focus {\\r\\n\\r\\n    background-color: rgb(0 0 0 / 0.2);\\r\\n\\r\\n    outline: 2px solid transparent;\\r\\n\\r\\n    outline-offset: 2px\\n}\\r\\n\\r\\n        .notifications-block__trigger:hover {\\r\\n\\r\\n    background-color: rgb(0 0 0 / 0.2)\\n}\\r\\n\\r\\n        .notifications-block__trigger.active {\\r\\n\\r\\n    --tw-bg-opacity: 1;\\r\\n\\r\\n    background-color: rgb(var(--base1) / var(--tw-bg-opacity));\\r\\n\\r\\n    --tw-text-opacity: 1;\\r\\n\\r\\n    color: rgb(var(--basec) / var(--tw-text-opacity))\\n}\\r\\n\\r\\n        .notifications-block__page {\\r\\n\\r\\n    max-width: 300px;\\r\\n\\r\\n    flex-grow: 1;\\r\\n\\r\\n    overflow: auto;\\r\\n\\r\\n    --tw-bg-opacity: 1;\\r\\n\\r\\n    background-color: rgb(255 255 255 / var(--tw-bg-opacity));\\r\\n\\r\\n    padding: 1.25rem;\\r\\n\\r\\n    outline: 2px solid transparent;\\r\\n\\r\\n    outline-offset: 2px;\\r\\n\\r\\n    border-bottom-left-radius: var(--rounded-box);\\r\\n\\r\\n    border-bottom-right-radius: var(--rounded-box)\\n}\\r\\n</style>"],"names":[],"mappings":"AAyDI,mCAAqB,CAErB,OAAO,CAAE,IAAI,CAEb,SAAS,CAAE,KAAK,CAEhB,cAAc,CAAE,MAAM,CAEtB,QAAQ,CAAE,MAAM;AACpB,CAEQ,2CAA6B,CAEjC,QAAQ,CAAE,QAAQ,CAElB,eAAe,CAAE,CAAC,CAElB,gBAAgB,CAAE,IAAI,IAAI,SAAS,CAAC,CAAC,CAAC,CAAC,IAAI,eAAe,CAAC,CAAC,CAE5D,YAAY,CAAE,IAAI,CAElB,aAAa,CAAE,IAAI,CAEnB,WAAW,CAAE,IAAI,CAEjB,cAAc,CAAE,MAAM,CAEtB,iBAAiB,CAAE,CAAC,CAEpB,KAAK,CAAE,IAAI,IAAI,YAAY,CAAC,CAAC,CAAC,CAAC,IAAI,iBAAiB,CAAC,CAAC,CAEtD,sBAAsB,CAAE,IAAI,aAAa,CAAC,CAE1C,uBAAuB,CAAE,IAAI,aAAa,CAAC;AAC/C,CAEQ,wCAA0B,CAE9B,QAAQ,CAAE,QAAQ,CAElB,IAAI,CAAE,MAAM,CAEZ,KAAK,CAAE,MAAM,CAEb,MAAM,CAAE,GAAG,CAEX,OAAO,CAAE,IAAI,CAEb,WAAW,CAAE,CAAC,CAEd,UAAU,CAAE,IAAI;AACpB,CAEQ,4CAA8B,CAElC,mBAAmB,CAAE,IAAI,CAEtB,gBAAgB,CAAE,IAAI,CAEjB,WAAW,CAAE,IAAI,CAEzB,sBAAsB,CAAE,OAAO,CAE/B,uBAAuB,CAAE,OAAO,CAEhC,YAAY,CAAE,OAAO,CAErB,aAAa,CAAE,OAAO,CAEtB,WAAW,CAAE,MAAM,CAEnB,cAAc,CAAE,MAAM;AAC1B,CAEQ,4CAA6B,MAAO,CAExC,gBAAgB,CAAE,IAAI,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,GAAG,CAAC,CAElC,OAAO,CAAE,GAAG,CAAC,KAAK,CAAC,WAAW,CAE9B,cAAc,CAAE,GAAG;AACvB,CAEQ,4CAA6B,MAAO,CAExC,gBAAgB,CAAE,IAAI,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,GAAG,CAAC;AACtC,CAEQ,6BAA6B,sBAAQ,CAEzC,eAAe,CAAE,CAAC,CAElB,gBAAgB,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,eAAe,CAAC,CAAC,CAE1D,iBAAiB,CAAE,CAAC,CAEpB,KAAK,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,iBAAiB,CAAC,CAAC;AACrD,CAEQ,yCAA2B,CAE/B,SAAS,CAAE,KAAK,CAEhB,SAAS,CAAE,CAAC,CAEZ,QAAQ,CAAE,IAAI,CAEd,eAAe,CAAE,CAAC,CAElB,gBAAgB,CAAE,IAAI,GAAG,CAAC,GAAG,CAAC,GAAG,CAAC,CAAC,CAAC,IAAI,eAAe,CAAC,CAAC,CAEzD,OAAO,CAAE,OAAO,CAEhB,OAAO,CAAE,GAAG,CAAC,KAAK,CAAC,WAAW,CAE9B,cAAc,CAAE,GAAG,CAEnB,yBAAyB,CAAE,IAAI,aAAa,CAAC,CAE7C,0BAA0B,CAAE,IAAI,aAAa,CAAC;AAClD"}`
};
const NotificationsBlock = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let __MELTUI_BUILDER_1__;
  let __MELTUI_BUILDER_2__;
  let __MELTUI_BUILDER_3__;
  let $content, $$unsubscribe_content;
  let $root, $$unsubscribe_root;
  let $list, $$unsubscribe_list;
  let $trigger, $$unsubscribe_trigger;
  let $value, $$unsubscribe_value;
  const { elements: { root, list, content, trigger }, states: { value } } = createTabs({ defaultValue: "all" });
  $$unsubscribe_root = subscribe(root, (value2) => $root = value2);
  $$unsubscribe_list = subscribe(list, (value2) => $list = value2);
  $$unsubscribe_content = subscribe(content, (value2) => $content = value2);
  $$unsubscribe_trigger = subscribe(trigger, (value2) => $trigger = value2);
  $$unsubscribe_value = subscribe(value, (value2) => $value = value2);
  let { class: className = "" } = $$props;
  let { dictionary } = $$props;
  const triggers = [{ id: "all" }, { id: "alerts" }];
  if ($$props.class === void 0 && $$bindings.class && className !== void 0) $$bindings.class(className);
  if ($$props.dictionary === void 0 && $$bindings.dictionary && dictionary !== void 0) $$bindings.dictionary(dictionary);
  $$result.css.add(css$5);
  __MELTUI_BUILDER_1__ = $content("all");
  __MELTUI_BUILDER_2__ = $content("messages");
  __MELTUI_BUILDER_3__ = $content("alerts");
  $$unsubscribe_content();
  $$unsubscribe_root();
  $$unsubscribe_list();
  $$unsubscribe_trigger();
  $$unsubscribe_value();
  return `<div${spread(
    [
      {
        class: escape_attribute_value("notifications-block" + cn(className))
      },
      escape_object($root)
    ],
    { classes: "svelte-1rz0x92" }
  )}><div class="notifications-block__header svelte-1rz0x92"><span class="text-xl font-semibold">${validate_component(Text, "Text").$$render(
    $$result,
    {
      key: "dashboard.navbar.notifications.title",
      source: dictionary
    },
    {},
    {}
  )}</span> <div${spread([escape_object($list), { class: "notifications-block__nav" }], { classes: "svelte-1rz0x92" })}>${each(triggers, (triggerItem) => {
    let __MELTUI_BUILDER_0__ = $trigger(triggerItem.id);
    return ` <button${spread(
      [
        { type: "button" },
        escape_object(__MELTUI_BUILDER_0__),
        { class: "notifications-block__trigger" }
      ],
      {
        classes: ($value === triggerItem.id ? "active" : "") + " svelte-1rz0x92"
      }
    )}>${triggerItem.id == "all" ? `${validate_component(Text, "Text").$$render(
      $$result,
      {
        key: "dashboard.navbar.notifications.tabs.all",
        source: dictionary
      },
      {},
      {}
    )}` : ``} ${triggerItem.id == "messages" ? `${validate_component(Text, "Text").$$render(
      $$result,
      {
        key: "dashboard.navbar.notifications.messages",
        source: dictionary
      },
      {},
      {}
    )}` : ``} ${triggerItem.id == "alerts" ? `${validate_component(Text, "Text").$$render(
      $$result,
      {
        key: "dashboard.navbar.notifications.tabs.alerts",
        source: dictionary
      },
      {},
      {}
    )}` : ``} </button>`;
  })}</div></div> <div${spread(
    [
      { tabindex: "-1" },
      escape_object(__MELTUI_BUILDER_1__),
      { class: "notifications-block__page" }
    ],
    { classes: "svelte-1rz0x92" }
  )}></div> <div${spread(
    [
      { tabindex: "-1" },
      escape_object(__MELTUI_BUILDER_2__),
      { class: "notifications-block__page" }
    ],
    { classes: "svelte-1rz0x92" }
  )}></div> <div${spread(
    [
      { tabindex: "-1" },
      escape_object(__MELTUI_BUILDER_3__),
      { class: "notifications-block__page" }
    ],
    { classes: "svelte-1rz0x92" }
  )}></div> </div>`;
});
const css$4 = {
  code: '.customizer-option.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x{--option-main-color:var(--primary);display:flex;cursor:pointer;flex-direction:column;align-items:center;padding:0.5rem}.customizer-option.svelte-1nf0f0x .svelte-1nf0f0x.svelte-1nf0f0x{cursor:pointer}.customizer-option.svelte-1nf0f0x>label.svelte-1nf0f0x.svelte-1nf0f0x:focus{outline:2px solid transparent !important;outline-offset:2px !important}.customizer-option.svelte-1nf0f0x>label.svelte-1nf0f0x:focus>.customizer-option__body.svelte-1nf0f0x,.customizer-option.svelte-1nf0f0x>label[aria-checked="true"].svelte-1nf0f0x>.customizer-option__body.svelte-1nf0f0x{border-color:rgb(var(--option-main-color))}.customizer-option.svelte-1nf0f0x>label.svelte-1nf0f0x:focus>.customizer-option__body.svelte-1nf0f0x{--option-focus-shadow:0 0 0 0.25rem rgba(var(--option-main-color) / .3);box-shadow:var(--option-focus-shadow)}.customizer-option__body.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x{display:flex;width:-moz-max-content;width:max-content;border-width:2px;--tw-border-opacity:1;border-color:rgb(var(--base3) / var(--tw-border-opacity));transition-property:all;transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1);transition-duration:150ms;border-radius:var(--rounded-btn)}.customizer-option__body.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x:hover{border-color:rgb(var(--option-main-color))}.customizer-option.primary.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x{--option-main-color:var(--primary)}.customizer-option.secondary.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x{--option-main-color:var(--secondary)}.customizer-option.neutral.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x{--option-main-color:var(--neutral)}.customizer-option.success.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x{--option-main-color:var(--success)}.customizer-option.info.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x{--option-main-color:var(--info)}.customizer-option.warning.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x{--option-main-color:var(--warning)}.customizer-option.error.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x{--option-main-color:var(--error)}.customizer-option.smoke.svelte-1nf0f0x.svelte-1nf0f0x.svelte-1nf0f0x{--option-main-color:var(--smoke)}',
  map: '{"version":3,"file":"CustomizerOption.svelte","sources":["CustomizerOption.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">export const PRIMARY = \\"primary\\";\\nexport const SECONDARY = \\"secondary\\";\\nexport const NEUTRAL = \\"neutral\\";\\nexport const SUCCESS = \\"success\\";\\nexport const INFO = \\"info\\";\\nexport const WARNING = \\"warning\\";\\nexport const ERROR = \\"error\\";\\nexport const SMOKE = \\"smoke\\";\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { fade } from \\"svelte/transition\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nexport let value;\\nexport let userSelect = \\"\\";\\nexport let palette = PRIMARY;\\nexport let error = null;\\n<\/script>\\r\\n\\r\\n<div class={\\"customizer-option\\" + cn(palette)}>\\r\\n    <!-- svelte-ignore a11y-no-noninteractive-tabindex -->\\r\\n    <!-- svelte-ignore a11y-no-noninteractive-element-to-interactive-role -->\\r\\n    <!-- svelte-ignore a11y-click-events-have-key-events -->\\r\\n    <label tabindex=\\"0\\" role=\\"radio\\" aria-checked={userSelect == value} on:click={() => userSelect = value}>\\r\\n        <span class=\\"customizer-option__body\\">\\r\\n            <slot name=\\"body\\" />\\r\\n        </span>\\r\\n    </label>\\r\\n    <slot name=\\"label\\" />\\r\\n\\r\\n    {#if error}\\r\\n        <span class=\\"text-error\\" transition:fade>{ error }</span>\\r\\n    {/if}\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .customizer-option {\\r\\n        --option-main-color: var(--primary);\\r\\n        display: flex;\\r\\n        cursor: pointer;\\r\\n        flex-direction: column;\\r\\n        align-items: center;\\r\\n        padding: 0.5rem;\\r\\n    }\\r\\n\\r\\n        .customizer-option * {\\r\\n        cursor: pointer;\\r\\n}\\r\\n\\r\\n        .customizer-option > label:focus {\\r\\n        outline: 2px solid transparent !important;\\r\\n        outline-offset: 2px !important;\\r\\n}\\r\\n\\r\\n        .customizer-option > label:focus > .customizer-option__body, .customizer-option > label[aria-checked=\\"true\\"] > .customizer-option__body {\\r\\n        border-color: rgb(var(--option-main-color));\\r\\n}\\r\\n\\r\\n        .customizer-option > label:focus > .customizer-option__body {\\r\\n            --option-focus-shadow: 0 0 0 0.25rem rgba(var(--option-main-color) / .3);\\r\\n            box-shadow: var(--option-focus-shadow);\\r\\n        }\\r\\n\\r\\n        .customizer-option__body {\\r\\n        display: flex;\\r\\n        width: -moz-max-content;\\r\\n        width: max-content;\\r\\n        border-width: 2px;\\r\\n        --tw-border-opacity: 1;\\r\\n        border-color: rgb(var(--base3) / var(--tw-border-opacity));\\r\\n        transition-property: all;\\r\\n        transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\\r\\n        transition-duration: 150ms;\\r\\n        border-radius: var(--rounded-btn);\\r\\n}\\r\\n\\r\\n        .customizer-option__body:hover {\\r\\n        border-color: rgb(var(--option-main-color));\\r\\n}\\r\\n\\r\\n        .customizer-option.primary {\\r\\n            --option-main-color: var(--primary);\\r\\n        }\\r\\n\\r\\n        .customizer-option.secondary {\\r\\n            --option-main-color: var(--secondary);\\r\\n        }\\r\\n\\r\\n        .customizer-option.neutral {\\r\\n            --option-main-color: var(--neutral);\\r\\n        }\\r\\n\\r\\n        .customizer-option.success {\\r\\n            --option-main-color: var(--success);\\r\\n        }\\r\\n\\r\\n        .customizer-option.info {\\r\\n            --option-main-color: var(--info);\\r\\n        }\\r\\n\\r\\n        .customizer-option.warning {\\r\\n            --option-main-color: var(--warning);\\r\\n        }\\r\\n\\r\\n        .customizer-option.error {\\r\\n            --option-main-color: var(--error);\\r\\n        }\\r\\n\\r\\n        .customizer-option.smoke {\\r\\n            --option-main-color: var(--smoke);\\r\\n        }\\r\\n</style>"],"names":[],"mappings":"AAmCI,+DAAmB,CACf,mBAAmB,CAAE,cAAc,CACnC,OAAO,CAAE,IAAI,CACb,MAAM,CAAE,OAAO,CACf,cAAc,CAAE,MAAM,CACtB,WAAW,CAAE,MAAM,CACnB,OAAO,CAAE,MACb,CAEI,iCAAkB,CAAC,8BAAE,CACrB,MAAM,CAAE,OAChB,CAEQ,iCAAkB,CAAG,mCAAK,MAAO,CACjC,OAAO,CAAE,GAAG,CAAC,KAAK,CAAC,WAAW,CAAC,UAAU,CACzC,cAAc,CAAE,GAAG,CAAC,UAC5B,CAEQ,iCAAkB,CAAG,oBAAK,MAAM,CAAG,uCAAwB,CAAE,iCAAkB,CAAG,KAAK,CAAC,YAAY,CAAC,MAAM,gBAAC,CAAG,uCAAyB,CACxI,YAAY,CAAE,IAAI,IAAI,mBAAmB,CAAC,CAClD,CAEQ,iCAAkB,CAAG,oBAAK,MAAM,CAAG,uCAAyB,CACxD,qBAAqB,CAAE,iDAAiD,CACxE,UAAU,CAAE,IAAI,qBAAqB,CACzC,CAEA,qEAAyB,CACzB,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,gBAAgB,CACvB,KAAK,CAAE,WAAW,CAClB,YAAY,CAAE,GAAG,CACjB,mBAAmB,CAAE,CAAC,CACtB,YAAY,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,mBAAmB,CAAC,CAAC,CAC1D,mBAAmB,CAAE,GAAG,CACxB,0BAA0B,CAAE,aAAa,GAAG,CAAC,CAAC,CAAC,CAAC,CAAC,GAAG,CAAC,CAAC,CAAC,CAAC,CACxD,mBAAmB,CAAE,KAAK,CAC1B,aAAa,CAAE,IAAI,aAAa,CACxC,CAEQ,qEAAwB,MAAO,CAC/B,YAAY,CAAE,IAAI,IAAI,mBAAmB,CAAC,CAClD,CAEQ,kBAAkB,qDAAS,CACvB,mBAAmB,CAAE,cACzB,CAEA,kBAAkB,uDAAW,CACzB,mBAAmB,CAAE,gBACzB,CAEA,kBAAkB,qDAAS,CACvB,mBAAmB,CAAE,cACzB,CAEA,kBAAkB,qDAAS,CACvB,mBAAmB,CAAE,cACzB,CAEA,kBAAkB,kDAAM,CACpB,mBAAmB,CAAE,WACzB,CAEA,kBAAkB,qDAAS,CACvB,mBAAmB,CAAE,cACzB,CAEA,kBAAkB,mDAAO,CACrB,mBAAmB,CAAE,YACzB,CAEA,kBAAkB,mDAAO,CACrB,mBAAmB,CAAE,YACzB"}'
};
const PRIMARY = "primary";
const CustomizerOption = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { value } = $$props;
  let { userSelect = "" } = $$props;
  let { palette = PRIMARY } = $$props;
  let { error = null } = $$props;
  if ($$props.value === void 0 && $$bindings.value && value !== void 0) $$bindings.value(value);
  if ($$props.userSelect === void 0 && $$bindings.userSelect && userSelect !== void 0) $$bindings.userSelect(userSelect);
  if ($$props.palette === void 0 && $$bindings.palette && palette !== void 0) $$bindings.palette(palette);
  if ($$props.error === void 0 && $$bindings.error && error !== void 0) $$bindings.error(error);
  $$result.css.add(css$4);
  return `<div class="${escape(null_to_empty("customizer-option" + cn(palette)), true) + " svelte-1nf0f0x"}">   <label tabindex="0" role="radio"${add_attribute("aria-checked", userSelect == value, 0)} class="svelte-1nf0f0x"><span class="customizer-option__body svelte-1nf0f0x">${slots.body ? slots.body({}) : ``}</span></label> ${slots.label ? slots.label({}) : ``} ${error ? `<span class="text-error svelte-1nf0f0x">${escape(error)}</span>` : ``} </div>`;
});
const SCORE_CONTINUE_MATCH = 1, SCORE_SPACE_WORD_JUMP = 0.9, SCORE_NON_SPACE_WORD_JUMP = 0.8, SCORE_CHARACTER_JUMP = 0.17, SCORE_TRANSPOSITION = 0.1, PENALTY_SKIPPED = 0.999, PENALTY_CASE_MISMATCH = 0.9999, PENALTY_NOT_COMPLETE = 0.99;
const IS_GAP_REGEXP = /[\\/_+.#"@[({&]/, COUNT_GAPS_REGEXP = /[\\/_+.#"@[({&]/g, IS_SPACE_REGEXP = /[\s-]/, COUNT_SPACE_REGEXP = /[\s-]/g;
function commandScoreInner(string, abbreviation, lowerString, lowerAbbreviation, stringIndex, abbreviationIndex, memoizedResults) {
  if (abbreviationIndex === abbreviation.length) {
    if (stringIndex === string.length) {
      return SCORE_CONTINUE_MATCH;
    }
    return PENALTY_NOT_COMPLETE;
  }
  const memoizeKey = `${stringIndex},${abbreviationIndex}`;
  if (memoizedResults[memoizeKey] !== void 0) {
    return memoizedResults[memoizeKey];
  }
  const abbreviationChar = lowerAbbreviation.charAt(abbreviationIndex);
  let index = lowerString.indexOf(abbreviationChar, stringIndex);
  let highScore = 0;
  let score, transposedScore, wordBreaks, spaceBreaks;
  while (index >= 0) {
    score = commandScoreInner(string, abbreviation, lowerString, lowerAbbreviation, index + 1, abbreviationIndex + 1, memoizedResults);
    if (score > highScore) {
      if (index === stringIndex) {
        score *= SCORE_CONTINUE_MATCH;
      } else if (IS_GAP_REGEXP.test(string.charAt(index - 1))) {
        score *= SCORE_NON_SPACE_WORD_JUMP;
        wordBreaks = string.slice(stringIndex, index - 1).match(COUNT_GAPS_REGEXP);
        if (wordBreaks && stringIndex > 0) {
          score *= Math.pow(PENALTY_SKIPPED, wordBreaks.length);
        }
      } else if (IS_SPACE_REGEXP.test(string.charAt(index - 1))) {
        score *= SCORE_SPACE_WORD_JUMP;
        spaceBreaks = string.slice(stringIndex, index - 1).match(COUNT_SPACE_REGEXP);
        if (spaceBreaks && stringIndex > 0) {
          score *= Math.pow(PENALTY_SKIPPED, spaceBreaks.length);
        }
      } else {
        score *= SCORE_CHARACTER_JUMP;
        if (stringIndex > 0) {
          score *= Math.pow(PENALTY_SKIPPED, index - stringIndex);
        }
      }
      if (string.charAt(index) !== abbreviation.charAt(abbreviationIndex)) {
        score *= PENALTY_CASE_MISMATCH;
      }
    }
    if (score < SCORE_TRANSPOSITION && lowerString.charAt(index - 1) === lowerAbbreviation.charAt(abbreviationIndex + 1) || lowerAbbreviation.charAt(abbreviationIndex + 1) === lowerAbbreviation.charAt(abbreviationIndex) && // allow duplicate letters. Ref #7428
    lowerString.charAt(index - 1) !== lowerAbbreviation.charAt(abbreviationIndex)) {
      transposedScore = commandScoreInner(string, abbreviation, lowerString, lowerAbbreviation, index + 1, abbreviationIndex + 2, memoizedResults);
      if (transposedScore * SCORE_TRANSPOSITION > score) {
        score = transposedScore * SCORE_TRANSPOSITION;
      }
    }
    if (score > highScore) {
      highScore = score;
    }
    index = lowerString.indexOf(abbreviationChar, index + 1);
  }
  memoizedResults[memoizeKey] = highScore;
  return highScore;
}
function formatInput(string) {
  return string.toLowerCase().replace(COUNT_SPACE_REGEXP, " ");
}
function commandScore(string, abbreviation) {
  return commandScoreInner(string, abbreviation, formatInput(string), formatInput(abbreviation), 0, 0, {});
}
const isBrowser$1 = typeof document !== "undefined";
function isHTMLElement$1(element) {
  return element instanceof HTMLElement;
}
function isUndefined(value) {
  return value === void 0;
}
function generateId$1() {
  return nanoid(10);
}
const kbd$1 = {
  ALT: "Alt",
  ARROW_DOWN: "ArrowDown",
  ARROW_LEFT: "ArrowLeft",
  ARROW_RIGHT: "ArrowRight",
  ARROW_UP: "ArrowUp",
  BACKSPACE: "Backspace",
  CAPS_LOCK: "CapsLock",
  CONTROL: "Control",
  DELETE: "Delete",
  END: "End",
  ENTER: "Enter",
  ESCAPE: "Escape",
  F1: "F1",
  F10: "F10",
  F11: "F11",
  F12: "F12",
  F2: "F2",
  F3: "F3",
  F4: "F4",
  F5: "F5",
  F6: "F6",
  F7: "F7",
  F8: "F8",
  F9: "F9",
  HOME: "Home",
  META: "Meta",
  PAGE_DOWN: "PageDown",
  PAGE_UP: "PageUp",
  SHIFT: "Shift",
  SPACE: " ",
  TAB: "Tab",
  CTRL: "Control",
  ASTERISK: "*"
};
function omit$1(obj, ...keys) {
  const result = {};
  for (const key of Object.keys(obj)) {
    if (!keys.includes(key)) {
      result[key] = obj[key];
    }
  }
  return result;
}
function removeUndefined$1(obj) {
  const result = {};
  for (const key in obj) {
    const value = obj[key];
    if (value !== void 0) {
      result[key] = value;
    }
  }
  return result;
}
function toWritableStores$1(properties) {
  const result = {};
  Object.keys(properties).forEach((key) => {
    const propertyKey = key;
    const value = properties[propertyKey];
    result[propertyKey] = writable(value);
  });
  return result;
}
function effect$1(stores, fn) {
  const unsub = derivedWithUnsubscribe$1(stores, (stores2, onUnsubscribe) => {
    return {
      stores: stores2,
      onUnsubscribe
    };
  }).subscribe(({ stores: stores2, onUnsubscribe }) => {
    const returned = fn(stores2);
    if (returned) {
      onUnsubscribe(returned);
    }
  });
  onDestroy(unsub);
  return unsub;
}
function derivedWithUnsubscribe$1(stores, fn) {
  let unsubscribers = [];
  const onUnsubscribe = (cb) => {
    unsubscribers.push(cb);
  };
  const unsubscribe = () => {
    unsubscribers.forEach((fn2) => fn2());
    unsubscribers = [];
  };
  const derivedStore = derived(stores, ($storeValues) => {
    unsubscribe();
    return fn($storeValues, onUnsubscribe);
  });
  onDestroy(unsubscribe);
  const subscribe2 = (...args) => {
    const unsub = derivedStore.subscribe(...args);
    return () => {
      unsub();
      unsubscribe();
    };
  };
  return {
    ...derivedStore,
    subscribe: subscribe2
  };
}
function styleToString$2(style) {
  return Object.keys(style).reduce((str, key) => {
    if (style[key] === void 0)
      return str;
    return str + `${key}:${style[key]};`;
  }, "");
}
const srOnlyStyles = {
  position: "absolute",
  width: "1px",
  height: "1px",
  padding: "0",
  margin: "-1px",
  overflow: "hidden",
  clip: "rect(0, 0, 0, 0)",
  whiteSpace: "nowrap",
  borderWidth: "0"
};
function addEventListener$1(target, event, handler, options) {
  const events = Array.isArray(event) ? event : [event];
  events.forEach((_event) => target.addEventListener(_event, handler, options));
  return () => {
    events.forEach((_event) => target.removeEventListener(_event, handler, options));
  };
}
function executeCallbacks$1(...callbacks) {
  return (...args) => {
    for (const callback of callbacks) {
      if (typeof callback === "function") {
        callback(...args);
      }
    }
  };
}
const NAME$m = "Command";
const STATE_NAME = "CommandState";
const GROUP_NAME = "CommandGroup";
const LIST_SELECTOR = `[data-cmdk-list-sizer]`;
const GROUP_SELECTOR = `[data-cmdk-group]`;
const GROUP_ITEMS_SELECTOR = `[data-cmdk-group-items]`;
const GROUP_HEADING_SELECTOR = `[data-cmdk-group-heading]`;
const ITEM_SELECTOR = `[data-cmdk-item]`;
const VALID_ITEM_SELECTOR = `${ITEM_SELECTOR}:not([aria-disabled="true"])`;
const VALUE_ATTR = `data-value`;
const defaultFilter = (value, search) => commandScore(value, search);
function getCtx$1() {
  return getContext(NAME$m);
}
function getState() {
  return getContext(STATE_NAME);
}
function createGroup(alwaysRender) {
  const id = generateId$1();
  setContext(GROUP_NAME, {
    id,
    alwaysRender: isUndefined(alwaysRender) ? false : alwaysRender
  });
  return { id };
}
function getGroup() {
  const context = getContext(GROUP_NAME);
  if (!context)
    return void 0;
  return context;
}
function createState(initialValues) {
  const defaultState = {
    search: "",
    value: "",
    filtered: {
      count: 0,
      items: /* @__PURE__ */ new Map(),
      groups: /* @__PURE__ */ new Set()
    }
  };
  const state = writable(initialValues ? { ...defaultState, ...removeUndefined$1(initialValues) } : defaultState);
  return state;
}
const defaults$1 = {
  label: "Command menu",
  shouldFilter: true,
  loop: false,
  onValueChange: void 0,
  value: void 0,
  filter: defaultFilter,
  ids: {
    root: generateId$1(),
    list: generateId$1(),
    label: generateId$1(),
    input: generateId$1()
  }
};
function createCommand(props) {
  const ids = {
    root: generateId$1(),
    list: generateId$1(),
    label: generateId$1(),
    input: generateId$1(),
    ...props.ids
  };
  const withDefaults = {
    ...defaults$1,
    ...removeUndefined$1(props)
  };
  const state = props.state ?? createState({
    value: withDefaults.value
  });
  const allItems = writable(/* @__PURE__ */ new Set());
  const allGroups = writable(/* @__PURE__ */ new Map());
  const allIds = writable(/* @__PURE__ */ new Map());
  const commandEl = writable(null);
  const options = toWritableStores$1(omit$1(withDefaults, "value", "ids"));
  let $allItems = get_store_value(allItems);
  let $allGroups = get_store_value(allGroups);
  let $allIds = get_store_value(allIds);
  let shouldFilter = get_store_value(options.shouldFilter);
  let loop = get_store_value(options.loop);
  let label = get_store_value(options.label);
  let filter = get_store_value(options.filter);
  effect$1(options.shouldFilter, ($shouldFilter) => {
    shouldFilter = $shouldFilter;
  });
  effect$1(options.loop, ($loop) => {
    loop = $loop;
  });
  effect$1(options.filter, ($filter) => {
    filter = $filter;
  });
  effect$1(options.label, ($label) => {
    label = $label;
  });
  effect$1(allItems, (v) => {
    $allItems = v;
  });
  effect$1(allGroups, (v) => {
    $allGroups = v;
  });
  effect$1(allIds, (v) => {
    $allIds = v;
  });
  const context = {
    value: (id, value) => {
      if (value !== $allIds.get(id)) {
        allIds.update(($allIds2) => {
          $allIds2.set(id, value);
          return $allIds2;
        });
        state.update(($state) => {
          $state.filtered.items.set(id, score(value, $state.search));
          return $state;
        });
      }
    },
    // Track item lifecycle (add/remove)
    item: (id, groupId) => {
      allItems.update(($allItems2) => $allItems2.add(id));
      if (groupId) {
        allGroups.update(($allGroups2) => {
          if (!$allGroups2.has(groupId)) {
            $allGroups2.set(groupId, /* @__PURE__ */ new Set([id]));
          } else {
            $allGroups2.get(groupId)?.add(id);
          }
          return $allGroups2;
        });
      }
      state.update(($state) => {
        const filteredState = filterItems($state, shouldFilter);
        if (!filteredState.value) {
          const value = selectFirstItem();
          filteredState.value = value ?? "";
        }
        return filteredState;
      });
      return () => {
        allIds.update(($allIds2) => {
          $allIds2.delete(id);
          return $allIds2;
        });
        allItems.update(($allItems2) => {
          $allItems2.delete(id);
          return $allItems2;
        });
        state.update(($state) => {
          $state.filtered.items.delete(id);
          const selectedItem = getSelectedItem();
          const filteredState = filterItems($state);
          if (selectedItem?.getAttribute("id") === id) {
            filteredState.value = selectFirstItem() ?? "";
          }
          return $state;
        });
      };
    },
    group: (id) => {
      allGroups.update(($allGroups2) => {
        if (!$allGroups2.has(id)) {
          $allGroups2.set(id, /* @__PURE__ */ new Set());
        }
        return $allGroups2;
      });
      return () => {
        allIds.update(($allIds2) => {
          $allIds2.delete(id);
          return $allIds2;
        });
        allGroups.update(($allGroups2) => {
          $allGroups2.delete(id);
          return $allGroups2;
        });
      };
    },
    filter: () => {
      return shouldFilter;
    },
    label: label || props["aria-label"] || "",
    commandEl,
    ids,
    updateState
  };
  function updateState(key, value, preventScroll) {
    state.update((curr) => {
      if (Object.is(curr[key], value))
        return curr;
      curr[key] = value;
      if (key === "search") {
        const filteredState = filterItems(curr, shouldFilter);
        curr = filteredState;
        const sortedState = sort(curr, shouldFilter);
        curr = sortedState;
        tick().then(() => state.update((curr2) => {
          curr2.value = selectFirstItem() ?? "";
          return curr2;
        }));
      } else if (key === "value") {
        props.onValueChange?.(curr.value);
        if (!preventScroll) {
          tick().then(() => scrollSelectedIntoView());
        }
      }
      return curr;
    });
  }
  function filterItems(state2, shouldFilterVal) {
    const $shouldFilter = shouldFilterVal ?? shouldFilter;
    if (!state2.search || !$shouldFilter) {
      state2.filtered.count = $allItems.size;
      return state2;
    }
    state2.filtered.groups = /* @__PURE__ */ new Set();
    let itemCount = 0;
    for (const id of $allItems) {
      const value = $allIds.get(id);
      const rank = score(value, state2.search);
      state2.filtered.items.set(id, rank);
      if (rank > 0) {
        itemCount++;
      }
    }
    for (const [groupId, group] of $allGroups) {
      for (const itemId of group) {
        const rank = state2.filtered.items.get(itemId);
        if (rank && rank > 0) {
          state2.filtered.groups.add(groupId);
        }
      }
    }
    state2.filtered.count = itemCount;
    return state2;
  }
  function sort(state2, shouldFilterVal) {
    const $shouldFilter = shouldFilterVal ?? shouldFilter;
    if (!state2.search || !$shouldFilter) {
      return state2;
    }
    const scores = state2.filtered.items;
    const groups = [];
    for (const value of state2.filtered.groups) {
      const items = $allGroups.get(value);
      if (!items)
        continue;
      let max = 0;
      for (const item of items) {
        const score2 = scores.get(item);
        if (isUndefined(score2))
          continue;
        max = Math.max(score2, max);
      }
      groups.push([value, max]);
    }
    const rootEl = document.getElementById(ids.root);
    if (!rootEl)
      return state2;
    const list = rootEl.querySelector(LIST_SELECTOR);
    const validItems = getValidItems(rootEl).sort((a, b) => {
      const valueA = a.getAttribute(VALUE_ATTR) ?? "";
      const valueB = b.getAttribute(VALUE_ATTR) ?? "";
      return (scores.get(valueA) ?? 0) - (scores.get(valueB) ?? 0);
    });
    for (const item of validItems) {
      const group = item.closest(GROUP_ITEMS_SELECTOR);
      const closest = item.closest(`${GROUP_ITEMS_SELECTOR} > *`);
      if (group) {
        if (item.parentElement === group) {
          group.appendChild(item);
        } else {
          if (!closest)
            continue;
          group.appendChild(closest);
        }
      } else {
        if (item.parentElement === list) {
          list?.appendChild(item);
        } else {
          if (!closest)
            continue;
          list?.appendChild(closest);
        }
      }
    }
    groups.sort((a, b) => b[1] - a[1]);
    for (const group of groups) {
      const el = rootEl.querySelector(`${GROUP_SELECTOR}[${VALUE_ATTR}="${group[0]}"]`);
      el?.parentElement?.appendChild(el);
    }
    return state2;
  }
  function selectFirstItem() {
    const item = getValidItems().find((item2) => !item2.ariaDisabled);
    if (!item)
      return;
    const value = item.getAttribute(VALUE_ATTR);
    if (!value)
      return;
    return value;
  }
  function score(value, search) {
    const lowerCaseAndTrimmedValue = value?.toLowerCase().trim();
    const filterFn = filter;
    if (!filterFn) {
      return lowerCaseAndTrimmedValue ? defaultFilter(lowerCaseAndTrimmedValue, search) : 0;
    }
    return lowerCaseAndTrimmedValue ? filterFn(lowerCaseAndTrimmedValue, search) : 0;
  }
  function scrollSelectedIntoView() {
    const item = getSelectedItem();
    if (!item) {
      return;
    }
    if (item.parentElement?.firstChild === item) {
      tick().then(() => item.closest(GROUP_SELECTOR)?.querySelector(GROUP_HEADING_SELECTOR)?.scrollIntoView({
        block: "nearest"
      }));
    }
    tick().then(() => item.scrollIntoView({ block: "nearest" }));
  }
  function getValidItems(rootElement) {
    const rootEl = rootElement ?? document.getElementById(ids.root);
    if (!rootEl)
      return [];
    return Array.from(rootEl.querySelectorAll(VALID_ITEM_SELECTOR)).filter((el) => el ? true : false);
  }
  function getSelectedItem(rootElement) {
    const rootEl = document.getElementById(ids.root);
    if (!rootEl)
      return;
    const selectedEl = rootEl.querySelector(`${VALID_ITEM_SELECTOR}[aria-selected="true"]`);
    if (!selectedEl)
      return;
    return selectedEl;
  }
  function updateSelectedToIndex(index) {
    const rootEl = document.getElementById(ids.root);
    if (!rootEl)
      return;
    const items = getValidItems(rootEl);
    const item = items[index];
    if (!item)
      return;
  }
  function updateSelectedByChange(change) {
    const selected = getSelectedItem();
    const items = getValidItems();
    const index = items.findIndex((item) => item === selected);
    let newSelected = items[index + change];
    if (loop) {
      if (index + change < 0) {
        newSelected = items[items.length - 1];
      } else if (index + change === items.length) {
        newSelected = items[0];
      } else {
        newSelected = items[index + change];
      }
    }
    if (newSelected) {
      updateState("value", newSelected.getAttribute(VALUE_ATTR) ?? "");
    }
  }
  function updateSelectedToGroup(change) {
    const selected = getSelectedItem();
    let group = selected?.closest(GROUP_SELECTOR);
    let item = void 0;
    while (group && !item) {
      group = change > 0 ? findNextSibling(group, GROUP_SELECTOR) : findPreviousSibling(group, GROUP_SELECTOR);
      item = group?.querySelector(VALID_ITEM_SELECTOR);
    }
    if (item) {
      updateState("value", item.getAttribute(VALUE_ATTR) ?? "");
    } else {
      updateSelectedByChange(change);
    }
  }
  function last2() {
    return updateSelectedToIndex(getValidItems().length - 1);
  }
  function next2(e) {
    e.preventDefault();
    if (e.metaKey) {
      last2();
    } else if (e.altKey) {
      updateSelectedToGroup(1);
    } else {
      updateSelectedByChange(1);
    }
  }
  function prev2(e) {
    e.preventDefault();
    if (e.metaKey) {
      updateSelectedToIndex(0);
    } else if (e.altKey) {
      updateSelectedToGroup(-1);
    } else {
      updateSelectedByChange(-1);
    }
  }
  function handleRootKeydown(e) {
    switch (e.key) {
      case kbd$1.ARROW_DOWN:
        next2(e);
        break;
      case kbd$1.ARROW_UP:
        prev2(e);
        break;
      case kbd$1.HOME:
        e.preventDefault();
        updateSelectedToIndex(0);
        break;
      case kbd$1.END:
        e.preventDefault();
        last2();
        break;
      case kbd$1.ENTER: {
        e.preventDefault();
        const item = getSelectedItem();
        if (item) {
          item?.click();
        }
      }
    }
  }
  setContext(NAME$m, context);
  const stateStore = {
    subscribe: state.subscribe,
    update: state.update,
    set: state.set,
    updateState
  };
  setContext(STATE_NAME, stateStore);
  return {
    state: stateStore,
    handleRootKeydown,
    commandEl,
    ids
  };
}
function findNextSibling(el, selector2) {
  let sibling = el.nextElementSibling;
  while (sibling) {
    if (sibling.matches(selector2))
      return sibling;
    sibling = sibling.nextElementSibling;
  }
}
function findPreviousSibling(el, selector2) {
  let sibling = el.previousElementSibling;
  while (sibling) {
    if (sibling.matches(selector2))
      return sibling;
    sibling = sibling.previousElementSibling;
  }
}
const Command = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let slotProps;
  let $$restProps = compute_rest_props($$props, [
    "label",
    "shouldFilter",
    "filter",
    "value",
    "onValueChange",
    "loop",
    "onKeydown",
    "state",
    "ids",
    "asChild"
  ]);
  let $stateStore, $$unsubscribe_stateStore;
  let { label = void 0 } = $$props;
  let { shouldFilter = true } = $$props;
  let { filter = void 0 } = $$props;
  let { value = void 0 } = $$props;
  let { onValueChange = void 0 } = $$props;
  let { loop = void 0 } = $$props;
  let { onKeydown = void 0 } = $$props;
  let { state = void 0 } = $$props;
  let { ids = void 0 } = $$props;
  let { asChild = false } = $$props;
  const { commandEl, handleRootKeydown, ids: commandIds, state: stateStore } = createCommand({
    label,
    shouldFilter,
    filter,
    value,
    onValueChange: (next2) => {
      if (next2 !== value) {
        value = next2;
        onValueChange?.(next2);
      }
    },
    loop,
    state,
    ids
  });
  $$unsubscribe_stateStore = subscribe(stateStore, (value2) => $stateStore = value2);
  function syncValueAndState(value2) {
    if (value2 && value2 !== $stateStore.value) {
      set_store_value(stateStore, $stateStore.value = value2, $stateStore);
    }
  }
  function rootAction(node) {
    commandEl.set(node);
    const unsubEvents = executeCallbacks$1(addEventListener$1(node, "keydown", handleKeydown));
    return { destroy: unsubEvents };
  }
  const rootAttrs = {
    role: "application",
    id: commandIds.root,
    "data-cmdk-root": ""
  };
  const labelAttrs = {
    "data-cmdk-label": "",
    for: commandIds.input,
    id: commandIds.label,
    style: styleToString$2(srOnlyStyles)
  };
  function handleKeydown(e) {
    onKeydown?.(e);
    if (e.defaultPrevented) return;
    handleRootKeydown(e);
  }
  const root = { action: rootAction, attrs: rootAttrs };
  if ($$props.label === void 0 && $$bindings.label && label !== void 0) $$bindings.label(label);
  if ($$props.shouldFilter === void 0 && $$bindings.shouldFilter && shouldFilter !== void 0) $$bindings.shouldFilter(shouldFilter);
  if ($$props.filter === void 0 && $$bindings.filter && filter !== void 0) $$bindings.filter(filter);
  if ($$props.value === void 0 && $$bindings.value && value !== void 0) $$bindings.value(value);
  if ($$props.onValueChange === void 0 && $$bindings.onValueChange && onValueChange !== void 0) $$bindings.onValueChange(onValueChange);
  if ($$props.loop === void 0 && $$bindings.loop && loop !== void 0) $$bindings.loop(loop);
  if ($$props.onKeydown === void 0 && $$bindings.onKeydown && onKeydown !== void 0) $$bindings.onKeydown(onKeydown);
  if ($$props.state === void 0 && $$bindings.state && state !== void 0) $$bindings.state(state);
  if ($$props.ids === void 0 && $$bindings.ids && ids !== void 0) $$bindings.ids(ids);
  if ($$props.asChild === void 0 && $$bindings.asChild && asChild !== void 0) $$bindings.asChild(asChild);
  {
    syncValueAndState(value);
  }
  slotProps = {
    root,
    label: { attrs: labelAttrs },
    stateStore,
    state: $stateStore
  };
  $$unsubscribe_stateStore();
  return `${asChild ? `${slots.default ? slots.default({ ...slotProps }) : ``}` : `<div${spread([escape_object(rootAttrs), escape_object($$restProps)], {})}> <label${spread([escape_object(labelAttrs)], {})}>${escape(label ?? "")}</label> ${slots.default ? slots.default({ ...slotProps }) : ``}</div>`}`;
});
function last(array) {
  return array[array.length - 1];
}
function styleToString$1(style) {
  return Object.keys(style).reduce((str, key) => {
    if (style[key] === void 0)
      return str;
    return str + `${key}:${style[key]};`;
  }, "");
}
({
  type: "hidden",
  "aria-hidden": true,
  hidden: true,
  tabIndex: -1,
  style: styleToString$1({
    position: "absolute",
    opacity: 0,
    "pointer-events": "none",
    margin: 0,
    transform: "translateX(-100%)"
  })
});
function lightable(value) {
  function subscribe2(run) {
    run(value);
    return () => {
    };
  }
  return { subscribe: subscribe2 };
}
const hiddenAction = (obj) => {
  return new Proxy(obj, {
    get(target, prop, receiver) {
      return Reflect.get(target, prop, receiver);
    },
    ownKeys(target) {
      return Reflect.ownKeys(target).filter((key) => key !== "action");
    }
  });
};
const isFunctionWithParams = (fn) => {
  return typeof fn === "function";
};
function builder(name2, args) {
  const { stores, action, returned } = args ?? {};
  const derivedStore = (() => {
    if (stores && returned) {
      return derived(stores, (values) => {
        const result = returned(values);
        if (isFunctionWithParams(result)) {
          const fn = (...args2) => {
            return hiddenAction({
              ...result(...args2),
              [`data-melt-${name2}`]: "",
              action: action ?? noop
            });
          };
          fn.action = action ?? noop;
          return fn;
        }
        return hiddenAction({
          ...result,
          [`data-melt-${name2}`]: "",
          action: action ?? noop
        });
      });
    } else {
      const returnedFn = returned;
      const result = returnedFn?.();
      if (isFunctionWithParams(result)) {
        const resultFn = (...args2) => {
          return hiddenAction({
            ...result(...args2),
            [`data-melt-${name2}`]: "",
            action: action ?? noop
          });
        };
        resultFn.action = action ?? noop;
        return lightable(resultFn);
      }
      return lightable(hiddenAction({
        ...result,
        [`data-melt-${name2}`]: "",
        action: action ?? noop
      }));
    }
  })();
  const actionFn = action ?? (() => {
  });
  actionFn.subscribe = derivedStore.subscribe;
  return actionFn;
}
function createElHelpers(prefix) {
  const name2 = (part) => part ? `${prefix}-${part}` : prefix;
  const attribute = (part) => `data-melt-${prefix}${part ? `-${part}` : ""}`;
  const selector2 = (part) => `[data-melt-${prefix}${part ? `-${part}` : ""}]`;
  const getEl = (part) => document.querySelector(selector2(part));
  return {
    name: name2,
    attribute,
    selector: selector2,
    getEl
  };
}
const isBrowser = typeof document !== "undefined";
const isFunction = (v) => typeof v === "function";
function isHTMLElement(element) {
  return element instanceof HTMLElement;
}
function executeCallbacks(...callbacks) {
  return (...args) => {
    for (const callback of callbacks) {
      if (typeof callback === "function") {
        callback(...args);
      }
    }
  };
}
function noop() {
}
function addEventListener(target, event, handler, options) {
  const events = Array.isArray(event) ? event : [event];
  events.forEach((_event) => target.addEventListener(_event, handler, options));
  return () => {
    events.forEach((_event) => target.removeEventListener(_event, handler, options));
  };
}
function addMeltEventListener(target, event, handler, options) {
  const events = Array.isArray(event) ? event : [event];
  if (typeof handler === "function") {
    const handlerWithMelt = withMelt((_event) => handler(_event));
    events.forEach((_event) => target.addEventListener(_event, handlerWithMelt, options));
    return () => {
      events.forEach((_event) => target.removeEventListener(_event, handlerWithMelt, options));
    };
  }
  return () => noop();
}
function dispatchMeltEvent(originalEvent) {
  const node = originalEvent.currentTarget;
  if (!isHTMLElement(node))
    return null;
  const customMeltEvent = new CustomEvent(`m-${originalEvent.type}`, {
    detail: {
      originalEvent
    },
    cancelable: true
  });
  node.dispatchEvent(customMeltEvent);
  return customMeltEvent;
}
function withMelt(handler) {
  return (event) => {
    const customEvent = dispatchMeltEvent(event);
    if (customEvent?.defaultPrevented)
      return;
    return handler(event);
  };
}
function omit(obj, ...keys) {
  const result = {};
  for (const key of Object.keys(obj)) {
    if (!keys.includes(key)) {
      result[key] = obj[key];
    }
  }
  return result;
}
const overridable = (store, onChange) => {
  const update = (updater, sideEffect) => {
    store.update((curr) => {
      const next2 = updater(curr);
      let res = next2;
      if (onChange) {
        res = onChange({ curr, next: next2 });
      }
      sideEffect?.(res);
      return res;
    });
  };
  const set = (curr) => {
    update(() => curr);
  };
  return {
    ...store,
    update,
    set
  };
};
function sleep$1(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}
function generateId() {
  return nanoid(10);
}
function generateIds(args) {
  return args.reduce((acc, curr) => {
    acc[curr] = generateId();
    return acc;
  }, {});
}
const kbd = {
  ALT: "Alt",
  ARROW_DOWN: "ArrowDown",
  ARROW_LEFT: "ArrowLeft",
  ARROW_RIGHT: "ArrowRight",
  ARROW_UP: "ArrowUp",
  BACKSPACE: "Backspace",
  CAPS_LOCK: "CapsLock",
  CONTROL: "Control",
  DELETE: "Delete",
  END: "End",
  ENTER: "Enter",
  ESCAPE: "Escape",
  F1: "F1",
  F10: "F10",
  F11: "F11",
  F12: "F12",
  F2: "F2",
  F3: "F3",
  F4: "F4",
  F5: "F5",
  F6: "F6",
  F7: "F7",
  F8: "F8",
  F9: "F9",
  HOME: "Home",
  META: "Meta",
  PAGE_DOWN: "PageDown",
  PAGE_UP: "PageUp",
  SHIFT: "Shift",
  SPACE: " ",
  TAB: "Tab",
  CTRL: "Control",
  ASTERISK: "*",
  A: "a",
  P: "p"
};
const isDom = () => typeof window !== "undefined";
function getPlatform() {
  const agent = navigator.userAgentData;
  return agent?.platform ?? navigator.platform;
}
const pt = (v) => isDom() && v.test(getPlatform().toLowerCase());
const isTouchDevice = () => isDom() && !!navigator.maxTouchPoints;
const isMac = () => pt(/^mac/) && !isTouchDevice();
const isApple = () => pt(/mac|iphone|ipad|ipod/i);
const isIos = () => isApple() && !isMac();
const LOCK_CLASSNAME = "data-melt-scroll-lock";
function assignStyle(el, style) {
  if (!el)
    return;
  const previousStyle = el.style.cssText;
  Object.assign(el.style, style);
  return () => {
    el.style.cssText = previousStyle;
  };
}
function setCSSProperty(el, property, value) {
  if (!el)
    return;
  const previousValue = el.style.getPropertyValue(property);
  el.style.setProperty(property, value);
  return () => {
    if (previousValue) {
      el.style.setProperty(property, previousValue);
    } else {
      el.style.removeProperty(property);
    }
  };
}
function getPaddingProperty(documentElement) {
  const documentLeft = documentElement.getBoundingClientRect().left;
  const scrollbarX = Math.round(documentLeft) + documentElement.scrollLeft;
  return scrollbarX ? "paddingLeft" : "paddingRight";
}
function removeScroll(_document) {
  const doc = document;
  const win = doc.defaultView ?? window;
  const { documentElement, body } = doc;
  const locked = body.hasAttribute(LOCK_CLASSNAME);
  if (locked)
    return noop;
  body.setAttribute(LOCK_CLASSNAME, "");
  const scrollbarWidth = win.innerWidth - documentElement.clientWidth;
  const setScrollbarWidthProperty = () => setCSSProperty(documentElement, "--scrollbar-width", `${scrollbarWidth}px`);
  const paddingProperty = getPaddingProperty(documentElement);
  const scrollbarSidePadding = win.getComputedStyle(body)[paddingProperty];
  const setStyle = () => assignStyle(body, {
    overflow: "hidden",
    [paddingProperty]: `calc(${scrollbarSidePadding} + ${scrollbarWidth}px)`
  });
  const setIOSStyle = () => {
    const { scrollX, scrollY, visualViewport } = win;
    const offsetLeft = visualViewport?.offsetLeft ?? 0;
    const offsetTop = visualViewport?.offsetTop ?? 0;
    const restoreStyle = assignStyle(body, {
      position: "fixed",
      overflow: "hidden",
      top: `${-(scrollY - Math.floor(offsetTop))}px`,
      left: `${-(scrollX - Math.floor(offsetLeft))}px`,
      right: "0",
      [paddingProperty]: `calc(${scrollbarSidePadding} + ${scrollbarWidth}px)`
    });
    return () => {
      restoreStyle?.();
      win.scrollTo(scrollX, scrollY);
    };
  };
  const cleanups = [setScrollbarWidthProperty(), isIos() ? setIOSStyle() : setStyle()];
  return () => {
    cleanups.forEach((fn) => fn?.());
    body.removeAttribute(LOCK_CLASSNAME);
  };
}
function derivedWithUnsubscribe(stores, fn) {
  let unsubscribers = [];
  const onUnsubscribe = (cb) => {
    unsubscribers.push(cb);
  };
  const unsubscribe = () => {
    unsubscribers.forEach((fn2) => fn2());
    unsubscribers = [];
  };
  const derivedStore = derived(stores, ($storeValues) => {
    unsubscribe();
    return fn($storeValues, onUnsubscribe);
  });
  onDestroy(unsubscribe);
  const subscribe2 = (...args) => {
    const unsub = derivedStore.subscribe(...args);
    return () => {
      unsub();
      unsubscribe();
    };
  };
  return {
    ...derivedStore,
    subscribe: subscribe2
  };
}
function effect(stores, fn) {
  const unsub = derivedWithUnsubscribe(stores, (stores2, onUnsubscribe) => {
    return {
      stores: stores2,
      onUnsubscribe
    };
  }).subscribe(({ stores: stores2, onUnsubscribe }) => {
    const returned = fn(stores2);
    if (returned) {
      onUnsubscribe(returned);
    }
  });
  onDestroy(unsub);
  return unsub;
}
function toWritableStores(properties) {
  const result = {};
  Object.keys(properties).forEach((key) => {
    const propertyKey = key;
    const value = properties[propertyKey];
    result[propertyKey] = writable(value);
  });
  return result;
}
function getPortalParent(node) {
  let parent = node.parentElement;
  while (isHTMLElement(parent) && !parent.hasAttribute("data-portal")) {
    parent = parent.parentElement;
  }
  return parent || "body";
}
function getPortalDestination(node, portalProp) {
  const portalParent = getPortalParent(node);
  if (portalProp !== void 0)
    return portalProp;
  if (portalParent === "body")
    return document.body;
  return null;
}
function handleFocus(args) {
  const { prop, defaultEl } = args;
  sleep$1(1).then(() => {
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
  });
}
const documentClickStore = readable(void 0, (set) => {
  function clicked(event) {
    set(event);
    set(void 0);
  }
  const unsubscribe = addEventListener(document, "pointerup", clicked, {
    passive: false,
    capture: true
  });
  return unsubscribe;
});
const useClickOutside = (node, config = {}) => {
  let options = { enabled: true, ...config };
  function isEnabled() {
    return typeof options.enabled === "boolean" ? options.enabled : get_store_value(options.enabled);
  }
  const unsubscribe = documentClickStore.subscribe((e) => {
    if (!isEnabled() || !e || e.target === node) {
      return;
    }
    const composedPath = e.composedPath();
    if (composedPath.includes(node))
      return;
    if (options.ignore) {
      if (isFunction(options.ignore)) {
        if (options.ignore(e))
          return;
      } else if (Array.isArray(options.ignore)) {
        if (options.ignore.length > 0 && options.ignore.some((ignoreEl) => {
          return ignoreEl && (e.target === ignoreEl || composedPath.includes(ignoreEl));
        }))
          return;
      }
    }
    options.handler?.(e);
  });
  return {
    update(params) {
      options = { ...options, ...params };
    },
    destroy() {
      unsubscribe();
    }
  };
};
const documentEscapeKeyStore = readable(void 0, (set) => {
  function keydown(event) {
    if (event && event.key === kbd.ESCAPE) {
      set(event);
    }
    set(void 0);
  }
  const unsubscribe = addEventListener(document, "keydown", keydown, {
    passive: false,
    capture: true
  });
  return unsubscribe;
});
const useEscapeKeydown = (node, config = {}) => {
  node.dataset.escapee = "";
  let options = { enabled: true, ...config };
  function isEnabled() {
    return typeof options.enabled === "boolean" ? options.enabled : get_store_value(options.enabled);
  }
  const unsubscribe = documentEscapeKeyStore.subscribe((e) => {
    if (!e || !isEnabled())
      return;
    const target = e.target;
    if (!isHTMLElement(target) || target.closest("[data-escapee]") !== node) {
      return;
    }
    e.preventDefault();
    if (options.ignore) {
      if (isFunction(options.ignore)) {
        if (options.ignore(e))
          return;
      } else if (Array.isArray(options.ignore)) {
        if (options.ignore.length > 0 && options.ignore.some((ignoreEl) => {
          return ignoreEl && target === ignoreEl;
        }))
          return;
      }
    }
    options.handler?.(e);
  });
  return {
    update(params) {
      options = { ...options, ...params };
    },
    destroy() {
      node.removeAttribute("data-escapee");
      unsubscribe();
    }
  };
};
function createFocusTrap(config = {}) {
  let trap;
  const { immediate, ...focusTrapOptions } = config;
  const hasFocus = writable(false);
  const isPaused = writable(false);
  const activate = (opts) => trap?.activate(opts);
  const deactivate = (opts) => {
    trap?.deactivate(opts);
  };
  const pause = () => {
    if (trap) {
      trap.pause();
      isPaused.set(true);
    }
  };
  const unpause = () => {
    if (trap) {
      trap.unpause();
      isPaused.set(false);
    }
  };
  const useFocusTrap2 = (node) => {
    trap = createFocusTrap$1(node, {
      ...focusTrapOptions,
      onActivate() {
        hasFocus.set(true);
        config.onActivate?.();
      },
      onDeactivate() {
        hasFocus.set(false);
        config.onDeactivate?.();
      }
    });
    if (immediate) {
      activate();
    }
    return {
      destroy() {
        deactivate();
        trap = void 0;
      }
    };
  };
  return {
    useFocusTrap: useFocusTrap2,
    hasFocus: readonly(hasFocus),
    isPaused: readonly(isPaused),
    activate,
    deactivate,
    pause,
    unpause
  };
}
const usePortal = (el, target = "body") => {
  let targetEl;
  if (!isHTMLElement(target) && typeof target !== "string") {
    return {
      destroy: noop
    };
  }
  async function update(newTarget) {
    target = newTarget;
    if (typeof target === "string") {
      targetEl = document.querySelector(target);
      if (targetEl === null) {
        await tick();
        targetEl = document.querySelector(target);
      }
      if (targetEl === null) {
        throw new Error(`No element found matching css selector: "${target}"`);
      }
    } else if (target instanceof HTMLElement) {
      targetEl = target;
    } else {
      throw new TypeError(`Unknown portal target type: ${target === null ? "null" : typeof target}. Allowed types: string (CSS selector) or HTMLElement.`);
    }
    el.dataset.portal = "";
    targetEl.appendChild(el);
    el.hidden = false;
  }
  function destroy() {
    el.remove();
  }
  update(target);
  return {
    update,
    destroy
  };
};
const { name } = createElHelpers("dialog");
const defaults = {
  preventScroll: true,
  closeOnEscape: true,
  closeOnOutsideClick: true,
  role: "dialog",
  defaultOpen: false,
  portal: "body",
  forceVisible: false,
  openFocus: void 0,
  closeFocus: void 0
};
const openDialogIds = writable([]);
const dialogIdParts = ["content", "title", "description"];
function createDialog(props) {
  const withDefaults = { ...defaults, ...props };
  const options = toWritableStores(omit(withDefaults, "ids"));
  const { preventScroll, closeOnEscape, closeOnOutsideClick, role, portal, forceVisible, openFocus, closeFocus } = options;
  const activeTrigger = writable(null);
  const ids = toWritableStores({
    ...generateIds(dialogIdParts),
    ...withDefaults.ids
  });
  const openWritable = withDefaults.open ?? writable(withDefaults.defaultOpen);
  const open = overridable(openWritable, withDefaults?.onOpenChange);
  const isVisible = derived([open, forceVisible], ([$open, $forceVisible]) => {
    return $open || $forceVisible;
  });
  let unsubScroll = noop;
  function handleOpen(e) {
    const el = e.currentTarget;
    const triggerEl = e.currentTarget;
    if (!isHTMLElement(el) || !isHTMLElement(triggerEl))
      return;
    open.set(true);
    activeTrigger.set(triggerEl);
  }
  function handleClose() {
    open.set(false);
    handleFocus({
      prop: get_store_value(closeFocus),
      defaultEl: get_store_value(activeTrigger)
    });
  }
  effect([open], ([$open]) => {
    sleep$1(100).then(() => {
      if ($open) {
        openDialogIds.update((prev2) => {
          prev2.push(get_store_value(ids.content));
          return prev2;
        });
      } else {
        openDialogIds.update((prev2) => prev2.filter((id) => id !== get_store_value(ids.content)));
      }
    });
  });
  const trigger = builder(name("trigger"), {
    stores: [open, ids.content],
    returned: ([$open, $contentId]) => {
      return {
        "aria-haspopup": "dialog",
        "aria-expanded": $open,
        "aria-controls": $contentId,
        type: "button"
      };
    },
    action: (node) => {
      const unsub = executeCallbacks(addMeltEventListener(node, "click", (e) => {
        handleOpen(e);
      }), addMeltEventListener(node, "keydown", (e) => {
        if (e.key !== kbd.ENTER && e.key !== kbd.SPACE)
          return;
        e.preventDefault();
        handleOpen(e);
      }));
      return {
        destroy: unsub
      };
    }
  });
  const overlay = builder(name("overlay"), {
    stores: [isVisible],
    returned: ([$isVisible]) => {
      return {
        hidden: $isVisible ? void 0 : true,
        tabindex: -1,
        style: styleToString$1({
          display: $isVisible ? void 0 : "none"
        }),
        "aria-hidden": true,
        "data-state": $isVisible ? "open" : "closed"
      };
    },
    action: (node) => {
      let unsubEscapeKeydown = noop;
      if (get_store_value(closeOnEscape)) {
        const escapeKeydown = useEscapeKeydown(node, {
          handler: () => {
            handleClose();
          }
        });
        if (escapeKeydown && escapeKeydown.destroy) {
          unsubEscapeKeydown = escapeKeydown.destroy;
        }
      }
      return {
        destroy() {
          unsubEscapeKeydown();
        }
      };
    }
  });
  const content = builder(name("content"), {
    stores: [isVisible, ids.content, ids.description, ids.title],
    returned: ([$isVisible, $contentId, $descriptionId, $titleId]) => {
      return {
        id: $contentId,
        role: get_store_value(role),
        "aria-describedby": $descriptionId,
        "aria-labelledby": $titleId,
        "data-state": $isVisible ? "open" : "closed",
        tabindex: -1,
        hidden: $isVisible ? void 0 : true,
        style: styleToString$1({
          display: $isVisible ? void 0 : "none"
        })
      };
    },
    action: (node) => {
      let activate = noop;
      let deactivate = noop;
      const destroy = executeCallbacks(effect([open], ([$open]) => {
        if (!$open)
          return;
        const focusTrap = createFocusTrap({
          immediate: false,
          escapeDeactivates: false,
          returnFocusOnDeactivate: false,
          fallbackFocus: node
        });
        activate = focusTrap.activate;
        deactivate = focusTrap.deactivate;
        const ac = focusTrap.useFocusTrap(node);
        if (ac && ac.destroy) {
          return ac.destroy;
        } else {
          return focusTrap.deactivate;
        }
      }), effect([closeOnOutsideClick, open], ([$closeOnOutsideClick, $open]) => {
        return useClickOutside(node, {
          enabled: $open,
          handler: (e) => {
            if (e.defaultPrevented)
              return;
            const $openDialogIds = get_store_value(openDialogIds);
            const isLast = last($openDialogIds) === get_store_value(ids.content);
            if ($closeOnOutsideClick && isLast) {
              handleClose();
            }
          }
        }).destroy;
      }), effect([closeOnEscape], ([$closeOnEscape]) => {
        if (!$closeOnEscape)
          return noop;
        const escapeKeydown = useEscapeKeydown(node, {
          handler: () => {
            handleClose();
          }
        });
        if (escapeKeydown && escapeKeydown.destroy) {
          return escapeKeydown.destroy;
        }
        return noop;
      }), effect([isVisible], ([$isVisible]) => {
        tick().then(() => {
          if (!$isVisible) {
            deactivate();
          } else {
            activate();
          }
        });
      }));
      return {
        destroy: () => {
          unsubScroll();
          destroy();
        }
      };
    }
  });
  const portalled = builder(name("portalled"), {
    stores: portal,
    returned: ($portal) => ({
      "data-portal": $portal ? "" : void 0
    }),
    action: (node) => {
      const unsubPortal = effect([portal], ([$portal]) => {
        if (!$portal)
          return noop;
        const portalDestination = getPortalDestination(node, $portal);
        if (portalDestination === null)
          return noop;
        const portalAction = usePortal(node, portalDestination);
        if (portalAction && portalAction.destroy) {
          return portalAction.destroy;
        } else {
          return noop;
        }
      });
      return {
        destroy() {
          unsubPortal();
        }
      };
    }
  });
  const title = builder(name("title"), {
    stores: [ids.title],
    returned: ([$titleId]) => ({
      id: $titleId
    })
  });
  const description = builder(name("description"), {
    stores: [ids.description],
    returned: ([$descriptionId]) => ({
      id: $descriptionId
    })
  });
  const close = builder(name("close"), {
    returned: () => ({
      type: "button"
    }),
    action: (node) => {
      const unsub = executeCallbacks(addMeltEventListener(node, "click", () => {
        handleClose();
      }), addMeltEventListener(node, "keydown", (e) => {
        if (e.key !== kbd.SPACE && e.key !== kbd.ENTER)
          return;
        e.preventDefault();
        handleClose();
      }));
      return {
        destroy: unsub
      };
    }
  });
  effect([open, preventScroll], ([$open, $preventScroll]) => {
    if (!isBrowser)
      return;
    if ($preventScroll && $open)
      unsubScroll = removeScroll();
    if ($open) {
      const contentEl = document.getElementById(get_store_value(ids.content));
      handleFocus({ prop: get_store_value(openFocus), defaultEl: contentEl });
    }
    return () => {
      if (!get_store_value(forceVisible)) {
        unsubScroll();
      }
    };
  });
  return {
    ids,
    elements: {
      content,
      trigger,
      title,
      description,
      overlay,
      close,
      portalled
    },
    states: {
      open
    },
    options
  };
}
function createBitAttrs(bit, parts) {
  const attrs = {};
  parts.forEach((part) => {
    attrs[part] = {
      [`data-bits-${bit}-${part}`]: ""
    };
  });
  return (part) => attrs[part];
}
function removeUndefined(obj) {
  const result = {};
  for (const key in obj) {
    const value = obj[key];
    if (value !== void 0) {
      result[key] = value;
    }
  }
  return result;
}
function styleToString(style) {
  return Object.keys(style).reduce((str, key) => {
    if (style[key] === void 0)
      return str;
    return str + `${key}:${style[key]};`;
  }, "");
}
styleToString({
  position: "absolute",
  width: "1px",
  height: "1px",
  padding: "0",
  margin: "-1px",
  overflow: "hidden",
  clip: "rect(0, 0, 0, 0)",
  whiteSpace: "nowrap",
  borderWidth: "0"
});
styleToString({
  position: "absolute",
  width: "25px",
  height: "25px",
  opacity: "0",
  margin: "0px",
  pointerEvents: "none",
  transform: "translateX(-100%)"
});
function getOptionUpdater(options) {
  return function(key, value) {
    if (value === void 0)
      return;
    const store = options[key];
    if (store) {
      store.set(value);
    }
  };
}
const NAME$l = "accordion";
const PARTS$l = ["root", "content", "header", "item", "trigger"];
createBitAttrs(NAME$l, PARTS$l);
const NAME$k = "alert-dialog";
const PARTS$k = [
  "action",
  "cancel",
  "content",
  "description",
  "overlay",
  "portal",
  "title",
  "trigger"
];
createBitAttrs(NAME$k, PARTS$k);
const NAME$j = "avatar";
const PARTS$j = ["root", "image", "fallback"];
createBitAttrs(NAME$j, PARTS$j);
const NAME$i = "checkbox";
const PARTS$i = ["root", "input", "indicator"];
createBitAttrs(NAME$i, PARTS$i);
const NAME$h = "collapsible";
const PARTS$h = ["root", "content", "trigger"];
createBitAttrs(NAME$h, PARTS$h);
const NAME$g = "context-menu";
const PARTS$g = [
  "arrow",
  "checkbox-indicator",
  "checkbox-item",
  "content",
  "group",
  "item",
  "label",
  "radio-group",
  "radio-item",
  "separator",
  "sub-content",
  "sub-trigger",
  "trigger"
];
createBitAttrs(NAME$g, PARTS$g);
const NAME$f = "dialog";
const PARTS$f = ["close", "content", "description", "overlay", "portal", "title", "trigger"];
const getAttrs = createBitAttrs(NAME$f, PARTS$f);
function setCtx(props) {
  const dialog = createDialog({ ...removeUndefined(props), role: "dialog" });
  setContext(NAME$f, dialog);
  return {
    ...dialog,
    updateOption: getOptionUpdater(dialog.options)
  };
}
function getCtx() {
  return getContext(NAME$f);
}
const Dialog = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $idValues, $$unsubscribe_idValues;
  let { preventScroll = void 0 } = $$props;
  let { closeOnEscape = void 0 } = $$props;
  let { closeOnOutsideClick = void 0 } = $$props;
  let { portal = void 0 } = $$props;
  let { forceVisible = true } = $$props;
  let { open = void 0 } = $$props;
  let { onOpenChange = void 0 } = $$props;
  let { openFocus = void 0 } = $$props;
  let { closeFocus = void 0 } = $$props;
  const { states: { open: localOpen }, updateOption, ids } = setCtx({
    closeOnEscape,
    preventScroll,
    closeOnOutsideClick,
    portal,
    forceVisible,
    defaultOpen: open,
    openFocus,
    closeFocus,
    onOpenChange: ({ next: next2 }) => {
      if (open !== next2) {
        onOpenChange?.(next2);
        open = next2;
      }
      return next2;
    }
  });
  const idValues = derived([ids.content, ids.description, ids.title], ([$contentId, $descriptionId, $titleId]) => ({
    content: $contentId,
    description: $descriptionId,
    title: $titleId
  }));
  $$unsubscribe_idValues = subscribe(idValues, (value) => $idValues = value);
  if ($$props.preventScroll === void 0 && $$bindings.preventScroll && preventScroll !== void 0) $$bindings.preventScroll(preventScroll);
  if ($$props.closeOnEscape === void 0 && $$bindings.closeOnEscape && closeOnEscape !== void 0) $$bindings.closeOnEscape(closeOnEscape);
  if ($$props.closeOnOutsideClick === void 0 && $$bindings.closeOnOutsideClick && closeOnOutsideClick !== void 0) $$bindings.closeOnOutsideClick(closeOnOutsideClick);
  if ($$props.portal === void 0 && $$bindings.portal && portal !== void 0) $$bindings.portal(portal);
  if ($$props.forceVisible === void 0 && $$bindings.forceVisible && forceVisible !== void 0) $$bindings.forceVisible(forceVisible);
  if ($$props.open === void 0 && $$bindings.open && open !== void 0) $$bindings.open(open);
  if ($$props.onOpenChange === void 0 && $$bindings.onOpenChange && onOpenChange !== void 0) $$bindings.onOpenChange(onOpenChange);
  if ($$props.openFocus === void 0 && $$bindings.openFocus && openFocus !== void 0) $$bindings.openFocus(openFocus);
  if ($$props.closeFocus === void 0 && $$bindings.closeFocus && closeFocus !== void 0) $$bindings.closeFocus(closeFocus);
  open !== void 0 && localOpen.set(open);
  {
    updateOption("preventScroll", preventScroll);
  }
  {
    updateOption("closeOnEscape", closeOnEscape);
  }
  {
    updateOption("closeOnOutsideClick", closeOnOutsideClick);
  }
  {
    updateOption("portal", portal);
  }
  {
    updateOption("forceVisible", forceVisible);
  }
  {
    updateOption("openFocus", openFocus);
  }
  {
    updateOption("closeFocus", closeFocus);
  }
  $$unsubscribe_idValues();
  return `${slots.default ? slots.default({ ids: $idValues }) : ``}`;
});
const DialogPortal = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let builder2;
  let $$restProps = compute_rest_props($$props, ["asChild"]);
  let $portalled, $$unsubscribe_portalled;
  let { asChild = false } = $$props;
  const { elements: { portalled } } = getCtx();
  $$unsubscribe_portalled = subscribe(portalled, (value) => $portalled = value);
  const attrs = getAttrs("portal");
  if ($$props.asChild === void 0 && $$bindings.asChild && asChild !== void 0) $$bindings.asChild(asChild);
  builder2 = $portalled;
  $$unsubscribe_portalled();
  return `${asChild ? `${slots.default ? slots.default({ builder: builder2, attrs }) : ``}` : `<div${spread([escape_object(builder2), escape_object($$restProps), escape_object(attrs)], {})}>${slots.default ? slots.default({ builder: builder2, attrs }) : ``}</div>`}`;
});
const DialogContent = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let builder2;
  let $$restProps = compute_rest_props($$props, [
    "transition",
    "transitionConfig",
    "inTransition",
    "inTransitionConfig",
    "outTransition",
    "outTransitionConfig",
    "asChild",
    "id"
  ]);
  let $content, $$unsubscribe_content;
  let $open, $$unsubscribe_open;
  let { transition = void 0 } = $$props;
  let { transitionConfig = void 0 } = $$props;
  let { inTransition = void 0 } = $$props;
  let { inTransitionConfig = void 0 } = $$props;
  let { outTransition = void 0 } = $$props;
  let { outTransitionConfig = void 0 } = $$props;
  let { asChild = false } = $$props;
  let { id = void 0 } = $$props;
  const { elements: { content }, states: { open }, ids } = getCtx();
  $$unsubscribe_content = subscribe(content, (value) => $content = value);
  $$unsubscribe_open = subscribe(open, (value) => $open = value);
  const attrs = getAttrs("content");
  if ($$props.transition === void 0 && $$bindings.transition && transition !== void 0) $$bindings.transition(transition);
  if ($$props.transitionConfig === void 0 && $$bindings.transitionConfig && transitionConfig !== void 0) $$bindings.transitionConfig(transitionConfig);
  if ($$props.inTransition === void 0 && $$bindings.inTransition && inTransition !== void 0) $$bindings.inTransition(inTransition);
  if ($$props.inTransitionConfig === void 0 && $$bindings.inTransitionConfig && inTransitionConfig !== void 0) $$bindings.inTransitionConfig(inTransitionConfig);
  if ($$props.outTransition === void 0 && $$bindings.outTransition && outTransition !== void 0) $$bindings.outTransition(outTransition);
  if ($$props.outTransitionConfig === void 0 && $$bindings.outTransitionConfig && outTransitionConfig !== void 0) $$bindings.outTransitionConfig(outTransitionConfig);
  if ($$props.asChild === void 0 && $$bindings.asChild && asChild !== void 0) $$bindings.asChild(asChild);
  if ($$props.id === void 0 && $$bindings.id && id !== void 0) $$bindings.id(id);
  {
    if (id) {
      ids.content.set(id);
    }
  }
  builder2 = $content;
  $$unsubscribe_content();
  $$unsubscribe_open();
  return `${asChild && $open ? `${slots.default ? slots.default({ builder: builder2, attrs }) : ``}` : `${transition && $open ? `<div${spread([escape_object(builder2), escape_object($$restProps), escape_object(attrs)], {})}>${slots.default ? slots.default({ builder: builder2, attrs }) : ``}</div>` : `${inTransition && outTransition && $open ? `<div${spread([escape_object(builder2), escape_object($$restProps), escape_object(attrs)], {})}>${slots.default ? slots.default({ builder: builder2, attrs }) : ``}</div>` : `${inTransition && $open ? `<div${spread(
    [
      escape_object(builder2),
      escape_object($$restProps),
      escape_object(attrs)
    ],
    {}
  )}>${slots.default ? slots.default({ builder: builder2, attrs }) : ``}</div>` : `${outTransition && $open ? `<div${spread(
    [
      escape_object(builder2),
      escape_object($$restProps),
      escape_object(attrs)
    ],
    {}
  )}>${slots.default ? slots.default({ builder: builder2, attrs }) : ``}</div>` : `${$open ? `<div${spread(
    [
      escape_object(builder2),
      escape_object($$restProps),
      escape_object(attrs)
    ],
    {}
  )}>${slots.default ? slots.default({ builder: builder2, attrs }) : ``}</div>` : ``}`}`}`}`}`}`;
});
const DialogOverlay = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let builder2;
  let $$restProps = compute_rest_props($$props, [
    "transition",
    "transitionConfig",
    "inTransition",
    "inTransitionConfig",
    "outTransition",
    "outTransitionConfig",
    "asChild"
  ]);
  let $overlay, $$unsubscribe_overlay;
  let $open, $$unsubscribe_open;
  let { transition = void 0 } = $$props;
  let { transitionConfig = void 0 } = $$props;
  let { inTransition = void 0 } = $$props;
  let { inTransitionConfig = void 0 } = $$props;
  let { outTransition = void 0 } = $$props;
  let { outTransitionConfig = void 0 } = $$props;
  let { asChild = false } = $$props;
  const { elements: { overlay }, states: { open } } = getCtx();
  $$unsubscribe_overlay = subscribe(overlay, (value) => $overlay = value);
  $$unsubscribe_open = subscribe(open, (value) => $open = value);
  const attrs = getAttrs("overlay");
  if ($$props.transition === void 0 && $$bindings.transition && transition !== void 0) $$bindings.transition(transition);
  if ($$props.transitionConfig === void 0 && $$bindings.transitionConfig && transitionConfig !== void 0) $$bindings.transitionConfig(transitionConfig);
  if ($$props.inTransition === void 0 && $$bindings.inTransition && inTransition !== void 0) $$bindings.inTransition(inTransition);
  if ($$props.inTransitionConfig === void 0 && $$bindings.inTransitionConfig && inTransitionConfig !== void 0) $$bindings.inTransitionConfig(inTransitionConfig);
  if ($$props.outTransition === void 0 && $$bindings.outTransition && outTransition !== void 0) $$bindings.outTransition(outTransition);
  if ($$props.outTransitionConfig === void 0 && $$bindings.outTransitionConfig && outTransitionConfig !== void 0) $$bindings.outTransitionConfig(outTransitionConfig);
  if ($$props.asChild === void 0 && $$bindings.asChild && asChild !== void 0) $$bindings.asChild(asChild);
  builder2 = $overlay;
  $$unsubscribe_overlay();
  $$unsubscribe_open();
  return `${asChild && $open ? `${slots.default ? slots.default({ builder: builder2, attrs }) : ``}` : `${transition && $open ? `<div${spread([escape_object(builder2), escape_object($$restProps), escape_object(attrs)], {})}></div>` : `${inTransition && outTransition && $open ? `<div${spread([escape_object(builder2), escape_object($$restProps), escape_object(attrs)], {})}></div>` : `${inTransition && $open ? `<div${spread(
    [
      escape_object(builder2),
      escape_object($$restProps),
      escape_object(attrs)
    ],
    {}
  )}></div>` : `${outTransition && $open ? `<div${spread(
    [
      escape_object(builder2),
      escape_object($$restProps),
      escape_object(attrs)
    ],
    {}
  )}></div>` : `${$open ? `<div${spread(
    [
      escape_object(builder2),
      escape_object($$restProps),
      escape_object(attrs)
    ],
    {}
  )}></div>` : ``}`}`}`}`}`}`;
});
const NAME$e = "dropdown-menu";
const PARTS$e = [
  "arrow",
  "checkbox-indicator",
  "checkbox-item",
  "content",
  "group",
  "item",
  "label",
  "radio-group",
  "radio-item",
  "separator",
  "sub-content",
  "sub-trigger",
  "trigger"
];
createBitAttrs(NAME$e, PARTS$e);
const NAME$d = "link-preview";
const PARTS$d = ["arrow", "content", "trigger"];
createBitAttrs(NAME$d, PARTS$d);
const NAME$c = "label";
const PARTS$c = ["root"];
createBitAttrs(NAME$c, PARTS$c);
const NAME$b = "menubar";
const PARTS$b = [
  "root",
  "arrow",
  "checkbox-indicator",
  "checkbox-item",
  "content",
  "group",
  "item",
  "label",
  "radio-group",
  "radio-item",
  "separator",
  "sub-content",
  "sub-trigger",
  "trigger"
];
createBitAttrs(NAME$b, PARTS$b);
const NAME$a = "popover";
const PARTS$a = ["arrow", "close", "content", "trigger"];
createBitAttrs(NAME$a, PARTS$a);
const NAME$9 = "progress";
const PARTS$9 = ["root"];
createBitAttrs(NAME$9, PARTS$9);
const NAME$8 = "radio-group";
const PARTS$8 = ["root", "item", "input"];
createBitAttrs(NAME$8, PARTS$8);
const NAME$7 = "select";
const PARTS$7 = ["arrow", "content", "group", "item", "input", "label", "trigger", "value"];
createBitAttrs(NAME$7, PARTS$7);
const NAME$6 = "separator";
const PARTS$6 = ["root"];
createBitAttrs(NAME$6, PARTS$6);
const NAME$5 = "slider";
const PARTS$5 = ["root", "input", "range", "thumb", "tick"];
createBitAttrs(NAME$5, PARTS$5);
const NAME$4 = "switch";
const PARTS$4 = ["root", "input", "thumb"];
createBitAttrs(NAME$4, PARTS$4);
const NAME$3 = "tabs";
const PARTS$3 = ["root", "content", "list", "trigger"];
createBitAttrs(NAME$3, PARTS$3);
const NAME$2 = "toggle";
const PARTS$2 = ["root", "input"];
createBitAttrs(NAME$2, PARTS$2);
const NAME$1 = "toggle-group";
const PARTS$1 = ["root", "item"];
createBitAttrs(NAME$1, PARTS$1);
const NAME = "tooltip";
const PARTS = ["arrow", "content", "trigger"];
createBitAttrs(NAME, PARTS);
const CommandDialog = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let overlayProps;
  let contentProps;
  let $$restProps = compute_rest_props($$props, [
    "open",
    "value",
    "portal",
    "overlayClasses",
    "contentClasses",
    "contentTransition",
    "contentTransitionConfig",
    "contentInTransition",
    "contentInTransitionConfig",
    "contentOutTransition",
    "contentOutTransitionConfig",
    "overlayTransition",
    "overlayTransitionConfig",
    "overlayInTransition",
    "overlayInTransitionConfig",
    "overlayOutTransition",
    "overlayOutTransitionConfig",
    "label"
  ]);
  let { open = false } = $$props;
  let { value = void 0 } = $$props;
  let { portal = void 0 } = $$props;
  let { overlayClasses = void 0 } = $$props;
  let { contentClasses = void 0 } = $$props;
  let { contentTransition = void 0 } = $$props;
  let { contentTransitionConfig = void 0 } = $$props;
  let { contentInTransition = void 0 } = $$props;
  let { contentInTransitionConfig = void 0 } = $$props;
  let { contentOutTransition = void 0 } = $$props;
  let { contentOutTransitionConfig = void 0 } = $$props;
  let { overlayTransition = void 0 } = $$props;
  let { overlayTransitionConfig = void 0 } = $$props;
  let { overlayInTransition = void 0 } = $$props;
  let { overlayInTransitionConfig = void 0 } = $$props;
  let { overlayOutTransition = void 0 } = $$props;
  let { overlayOutTransitionConfig = void 0 } = $$props;
  let { label = void 0 } = $$props;
  if ($$props.open === void 0 && $$bindings.open && open !== void 0) $$bindings.open(open);
  if ($$props.value === void 0 && $$bindings.value && value !== void 0) $$bindings.value(value);
  if ($$props.portal === void 0 && $$bindings.portal && portal !== void 0) $$bindings.portal(portal);
  if ($$props.overlayClasses === void 0 && $$bindings.overlayClasses && overlayClasses !== void 0) $$bindings.overlayClasses(overlayClasses);
  if ($$props.contentClasses === void 0 && $$bindings.contentClasses && contentClasses !== void 0) $$bindings.contentClasses(contentClasses);
  if ($$props.contentTransition === void 0 && $$bindings.contentTransition && contentTransition !== void 0) $$bindings.contentTransition(contentTransition);
  if ($$props.contentTransitionConfig === void 0 && $$bindings.contentTransitionConfig && contentTransitionConfig !== void 0) $$bindings.contentTransitionConfig(contentTransitionConfig);
  if ($$props.contentInTransition === void 0 && $$bindings.contentInTransition && contentInTransition !== void 0) $$bindings.contentInTransition(contentInTransition);
  if ($$props.contentInTransitionConfig === void 0 && $$bindings.contentInTransitionConfig && contentInTransitionConfig !== void 0) $$bindings.contentInTransitionConfig(contentInTransitionConfig);
  if ($$props.contentOutTransition === void 0 && $$bindings.contentOutTransition && contentOutTransition !== void 0) $$bindings.contentOutTransition(contentOutTransition);
  if ($$props.contentOutTransitionConfig === void 0 && $$bindings.contentOutTransitionConfig && contentOutTransitionConfig !== void 0) $$bindings.contentOutTransitionConfig(contentOutTransitionConfig);
  if ($$props.overlayTransition === void 0 && $$bindings.overlayTransition && overlayTransition !== void 0) $$bindings.overlayTransition(overlayTransition);
  if ($$props.overlayTransitionConfig === void 0 && $$bindings.overlayTransitionConfig && overlayTransitionConfig !== void 0) $$bindings.overlayTransitionConfig(overlayTransitionConfig);
  if ($$props.overlayInTransition === void 0 && $$bindings.overlayInTransition && overlayInTransition !== void 0) $$bindings.overlayInTransition(overlayInTransition);
  if ($$props.overlayInTransitionConfig === void 0 && $$bindings.overlayInTransitionConfig && overlayInTransitionConfig !== void 0) $$bindings.overlayInTransitionConfig(overlayInTransitionConfig);
  if ($$props.overlayOutTransition === void 0 && $$bindings.overlayOutTransition && overlayOutTransition !== void 0) $$bindings.overlayOutTransition(overlayOutTransition);
  if ($$props.overlayOutTransitionConfig === void 0 && $$bindings.overlayOutTransitionConfig && overlayOutTransitionConfig !== void 0) $$bindings.overlayOutTransitionConfig(overlayOutTransitionConfig);
  if ($$props.label === void 0 && $$bindings.label && label !== void 0) $$bindings.label(label);
  let $$settled;
  let $$rendered;
  let previous_head = $$result.head;
  do {
    $$settled = true;
    $$result.head = previous_head;
    overlayProps = {
      class: overlayClasses,
      transition: overlayTransition,
      transitionConfig: overlayTransitionConfig,
      inTransition: overlayInTransition,
      inTransitionConfig: overlayInTransitionConfig,
      outTransition: overlayOutTransition,
      outTransitionConfig: overlayOutTransitionConfig,
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      "data-cmdk-overlay": ""
    };
    contentProps = {
      class: contentClasses,
      transition: contentTransition,
      transitionConfig: contentTransitionConfig,
      inTransition: contentInTransition,
      inTransitionConfig: contentInTransitionConfig,
      outTransition: contentOutTransition,
      outTransitionConfig: contentOutTransitionConfig,
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      "data-cmdk-dialog": ""
    };
    $$rendered = `${validate_component(Dialog, "DialogPrimitive.Root").$$render(
      $$result,
      Object.assign({}, $$restProps, { open }),
      {
        open: ($$value) => {
          open = $$value;
          $$settled = false;
        }
      },
      {
        default: () => {
          return `${portal === null ? `${validate_component(DialogOverlay, "DialogPrimitive.Overlay").$$render($$result, Object.assign({}, overlayProps), {}, {})} ${validate_component(DialogContent, "DialogPrimitive.Content").$$render($$result, Object.assign({}, { "aria-label": label }, contentProps), {}, {
            default: () => {
              return `${validate_component(Command, "Command.Root").$$render(
                $$result,
                Object.assign({}, $$restProps, { value }),
                {
                  value: ($$value) => {
                    value = $$value;
                    $$settled = false;
                  }
                },
                {
                  default: () => {
                    return `${slots.default ? slots.default({}) : ``}`;
                  }
                }
              )}`;
            }
          })}` : `${validate_component(DialogPortal, "DialogPrimitive.Portal").$$render($$result, {}, {}, {
            default: () => {
              return `${validate_component(DialogOverlay, "DialogPrimitive.Overlay").$$render($$result, Object.assign({}, overlayProps), {}, {})} ${validate_component(DialogContent, "DialogPrimitive.Content").$$render($$result, Object.assign({}, { "aria-label": label }, contentProps), {}, {
                default: () => {
                  return `${validate_component(Command, "Command.Root").$$render(
                    $$result,
                    Object.assign({}, $$restProps, { value }),
                    {
                      value: ($$value) => {
                        value = $$value;
                        $$settled = false;
                      }
                    },
                    {
                      default: () => {
                        return `${slots.default ? slots.default({}) : ``}`;
                      }
                    }
                  )}`;
                }
              })}`;
            }
          })}`}`;
        }
      }
    )}`;
  } while (!$$settled);
  return $$rendered;
});
const CommandEmpty = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  compute_rest_props($$props, ["asChild"]);
  let $state, $$unsubscribe_state;
  let { asChild = false } = $$props;
  const state = getState();
  $$unsubscribe_state = subscribe(state, (value) => $state = value);
  if ($$props.asChild === void 0 && $$bindings.asChild && asChild !== void 0) $$bindings.asChild(asChild);
  $state.filtered.count === 0;
  $$unsubscribe_state();
  return `${``}`;
});
const CommandGroup = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let containerAttrs;
  let groupAttrs;
  let container;
  let group;
  let $$restProps = compute_rest_props($$props, ["heading", "value", "alwaysRender", "asChild"]);
  let $render, $$unsubscribe_render;
  let { heading = void 0 } = $$props;
  let { value = "" } = $$props;
  let { alwaysRender = false } = $$props;
  let { asChild = false } = $$props;
  const { id } = createGroup(alwaysRender);
  const context = getCtx$1();
  const state = getState();
  const headingId = generateId$1();
  const render = derived(state, ($state) => {
    if (alwaysRender) return true;
    if (context.filter() === false) return true;
    if (!$state.search) return true;
    return $state.filtered.groups.has(id);
  });
  $$unsubscribe_render = subscribe(render, (value2) => $render = value2);
  function containerAction(node) {
    if (value) {
      context.value(id, value);
      node.setAttribute(VALUE_ATTR, value);
      return;
    }
    if (heading) {
      value = heading.trim().toLowerCase();
    } else if (node.textContent) {
      value = node.textContent.trim().toLowerCase();
    }
    context.value(id, value);
    node.setAttribute(VALUE_ATTR, value);
  }
  const headingAttrs = {
    "data-cmdk-group-heading": "",
    "aria-hidden": true,
    id: headingId
  };
  if ($$props.heading === void 0 && $$bindings.heading && heading !== void 0) $$bindings.heading(heading);
  if ($$props.value === void 0 && $$bindings.value && value !== void 0) $$bindings.value(value);
  if ($$props.alwaysRender === void 0 && $$bindings.alwaysRender && alwaysRender !== void 0) $$bindings.alwaysRender(alwaysRender);
  if ($$props.asChild === void 0 && $$bindings.asChild && asChild !== void 0) $$bindings.asChild(asChild);
  containerAttrs = {
    "data-cmdk-group": "",
    role: "presentation",
    hidden: $render ? void 0 : true,
    "data-value": value
  };
  groupAttrs = {
    "data-cmdk-group-items": "",
    role: "group",
    "aria-labelledby": heading ? headingId : void 0
  };
  container = {
    action: containerAction,
    attrs: containerAttrs
  };
  group = { attrs: groupAttrs };
  $$unsubscribe_render();
  return `${asChild ? `${slots.default ? slots.default({
    container,
    group,
    heading: { attrs: headingAttrs }
  }) : ``}` : `<div${spread([escape_object(containerAttrs), escape_object($$restProps)], {})}>${heading ? `<div${spread([escape_object(headingAttrs)], {})}>${escape(heading)}</div>` : ``} <div${spread([escape_object(groupAttrs)], {})}>${slots.default ? slots.default({
    container,
    group,
    heading: { attrs: headingAttrs }
  }) : ``}</div></div>`}`;
});
function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}
const CommandInput = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $$restProps = compute_rest_props($$props, ["autofocus", "value", "asChild", "el"]);
  let $selectedItemId, $$unsubscribe_selectedItemId;
  const { ids, commandEl } = getCtx$1();
  const state = getState();
  const search = derived(state, ($state) => $state.search);
  const valueStore = derived(state, ($state) => $state.value);
  let { autofocus = void 0 } = $$props;
  let { value = get_store_value(search) } = $$props;
  let { asChild = false } = $$props;
  let { el = void 0 } = $$props;
  const selectedItemId = derived([valueStore, commandEl], ([$value, $commandEl]) => {
    if (!isBrowser$1) return void 0;
    const item = $commandEl?.querySelector(`${ITEM_SELECTOR}[${VALUE_ATTR}="${$value}"]`);
    return item?.getAttribute("id");
  });
  $$unsubscribe_selectedItemId = subscribe(selectedItemId, (value2) => $selectedItemId = value2);
  function handleValueUpdate(v) {
    state.updateState("search", v);
  }
  function action(node) {
    if (autofocus) {
      sleep(10).then(() => node.focus());
    }
    if (asChild) {
      const unsubEvents = addEventListener$1(node, "change", (e) => {
        const currTarget = e.currentTarget;
        state.updateState("search", currTarget.value);
      });
      return { destroy: unsubEvents };
    }
  }
  let attrs;
  if ($$props.autofocus === void 0 && $$bindings.autofocus && autofocus !== void 0) $$bindings.autofocus(autofocus);
  if ($$props.value === void 0 && $$bindings.value && value !== void 0) $$bindings.value(value);
  if ($$props.asChild === void 0 && $$bindings.asChild && asChild !== void 0) $$bindings.asChild(asChild);
  if ($$props.el === void 0 && $$bindings.el && el !== void 0) $$bindings.el(el);
  {
    handleValueUpdate(value);
  }
  attrs = {
    type: "text",
    "data-cmdk-input": "",
    autocomplete: "off",
    autocorrect: "off",
    spellcheck: false,
    "aria-autocomplete": "list",
    role: "combobox",
    "aria-expanded": true,
    "aria-controls": ids.list,
    "aria-labelledby": ids.label,
    "aria-activedescendant": $selectedItemId ?? void 0,
    id: ids.input
  };
  $$unsubscribe_selectedItemId();
  return `${asChild ? `${slots.default ? slots.default({ action, attrs }) : ``}` : `<input${spread([escape_object(attrs), escape_object($$restProps)], {})}${add_attribute("this", el, 0)}${add_attribute("value", value, 0)}>`}`;
});
const CommandItem = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let attrs;
  let $$restProps = compute_rest_props($$props, ["disabled", "value", "onSelect", "alwaysRender", "asChild", "id"]);
  let $selected, $$unsubscribe_selected;
  let $render, $$unsubscribe_render;
  let { disabled = false } = $$props;
  let { value = "" } = $$props;
  let { onSelect = void 0 } = $$props;
  let { alwaysRender = false } = $$props;
  let { asChild = false } = $$props;
  let { id = generateId$1() } = $$props;
  const groupContext = getGroup();
  const context = getCtx$1();
  const state = getState();
  const trueAlwaysRender = alwaysRender ?? groupContext?.alwaysRender;
  const render = derived(state, ($state) => {
    if (trueAlwaysRender || context.filter() === false || !$state.search) return true;
    const currentScore = $state.filtered.items.get(id);
    if (isUndefined(currentScore)) return false;
    return currentScore > 0;
  });
  $$unsubscribe_render = subscribe(render, (value2) => $render = value2);
  let isFirstRender = true;
  const selected = derived(state, ($state) => $state.value === value);
  $$unsubscribe_selected = subscribe(selected, (value2) => $selected = value2);
  function action(node) {
    if (!value && node.textContent) {
      value = node.textContent.trim().toLowerCase();
    }
    context.value(id, value);
    node.setAttribute(VALUE_ATTR, value);
    const unsubEvents = executeCallbacks$1(
      addEventListener$1(node, "pointermove", () => {
        if (disabled) return;
        select();
      }),
      addEventListener$1(node, "click", () => {
        if (disabled) return;
        handleItemClick();
      })
    );
    return {
      destroy() {
        unsubEvents();
      }
    };
  }
  function handleItemClick() {
    select();
    onSelect?.(value);
  }
  function select() {
    state.updateState("value", value, true);
  }
  if ($$props.disabled === void 0 && $$bindings.disabled && disabled !== void 0) $$bindings.disabled(disabled);
  if ($$props.value === void 0 && $$bindings.value && value !== void 0) $$bindings.value(value);
  if ($$props.onSelect === void 0 && $$bindings.onSelect && onSelect !== void 0) $$bindings.onSelect(onSelect);
  if ($$props.alwaysRender === void 0 && $$bindings.alwaysRender && alwaysRender !== void 0) $$bindings.alwaysRender(alwaysRender);
  if ($$props.asChild === void 0 && $$bindings.asChild && asChild !== void 0) $$bindings.asChild(asChild);
  if ($$props.id === void 0 && $$bindings.id && id !== void 0) $$bindings.id(id);
  attrs = {
    "aria-disabled": disabled ? true : void 0,
    "aria-selected": $selected ? true : void 0,
    "data-disabled": disabled ? true : void 0,
    "data-selected": $selected ? true : void 0,
    "data-cmdk-item": "",
    "data-value": value,
    role: "option",
    id
  };
  $$unsubscribe_selected();
  $$unsubscribe_render();
  return `${$render || isFirstRender ? `${asChild ? `${slots.default ? slots.default({ action, attrs }) : ``}` : `<div${spread([escape_object(attrs), escape_object($$restProps)], {})}>${slots.default ? slots.default({ action, attrs }) : ``}</div>`}` : ``}`;
});
const CommandList = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $$restProps = compute_rest_props($$props, ["el", "asChild"]);
  let $$unsubscribe_state;
  const { ids } = getCtx$1();
  const state = getState();
  $$unsubscribe_state = subscribe(state, (value) => value);
  let { el = void 0 } = $$props;
  let { asChild = false } = $$props;
  function sizerAction(node) {
    let animationFrame;
    const listEl = node.closest("[data-cmdk-list]");
    if (!isHTMLElement$1(listEl)) {
      return;
    }
    const observer = new ResizeObserver(() => {
      animationFrame = requestAnimationFrame(() => {
        const height = node.offsetHeight;
        listEl.style.setProperty("--cmdk-list-height", height.toFixed(1) + "px");
      });
    });
    observer.observe(node);
    return {
      destroy() {
        cancelAnimationFrame(animationFrame);
        observer.unobserve(node);
      }
    };
  }
  const listAttrs = {
    "data-cmdk-list": "",
    role: "listbox",
    "aria-label": "Suggestions",
    id: ids.list,
    "aria-labelledby": ids.input
  };
  const sizerAttrs = { "data-cmdk-list-sizer": "" };
  const list = { attrs: listAttrs };
  const sizer = { attrs: sizerAttrs, action: sizerAction };
  if ($$props.el === void 0 && $$bindings.el && el !== void 0) $$bindings.el(el);
  if ($$props.asChild === void 0 && $$bindings.asChild && asChild !== void 0) $$bindings.asChild(asChild);
  $$unsubscribe_state();
  return `${asChild ? `${slots.default ? slots.default({ list, sizer }) : ``}` : `<div${spread([escape_object(listAttrs), escape_object($$restProps)], {})}${add_attribute("this", el, 0)}><div${spread([escape_object(sizerAttrs)], {})}>${slots.default ? slots.default({}) : ``}</div></div>`}`;
});
const CommandSeparator = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $$restProps = compute_rest_props($$props, ["alwaysRender", "asChild"]);
  let $render, $$unsubscribe_render;
  let { alwaysRender = false } = $$props;
  let { asChild = false } = $$props;
  const state = getState();
  const render = derived(state, ($state) => !$state.search);
  $$unsubscribe_render = subscribe(render, (value) => $render = value);
  const attrs = {
    "data-cmdk-separator": "",
    role: "separator"
  };
  if ($$props.alwaysRender === void 0 && $$bindings.alwaysRender && alwaysRender !== void 0) $$bindings.alwaysRender(alwaysRender);
  if ($$props.asChild === void 0 && $$bindings.asChild && asChild !== void 0) $$bindings.asChild(asChild);
  $$unsubscribe_render();
  return `${$render || alwaysRender ? `${asChild ? `${slots.default ? slots.default({ attrs }) : ``}` : `<div${spread([escape_object(attrs), escape_object($$restProps)], {})}></div>`}` : ``}`;
});
const css$3 = {
  code: '[data-cmdk-root]{width:90%;max-width:42rem;overflow:hidden;--tw-bg-opacity:1;background-color:rgb(var(--base1) / var(--tw-bg-opacity));--tw-text-opacity:1;color:rgb(var(--basec) / var(--tw-text-opacity));--tw-shadow:0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);--tw-shadow-colored:0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);border-radius:var(--rounded-box);position:fixed;left:50%;top:50%;z-index:var(--dialog-z-index);--tw-translate-x:-50%;--tw-translate-y:-50%;transform:translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))}[data-cmdk-input]{width:100%;border-bottom-width:1px;--tw-border-opacity:1;border-bottom-color:rgb(var(--smoke) / var(--tw-border-opacity));background-color:transparent;padding:1.25rem;font-size:1.125rem;line-height:1.75rem;caret-color:rgb(var(--primary) / 1);outline:2px solid transparent;outline-offset:2px}[data-cmdk-input]::-moz-placeholder{--tw-text-opacity:1;color:rgb(var(--smoke) / var(--tw-text-opacity))}[data-cmdk-input]::placeholder{--tw-text-opacity:1;color:rgb(var(--smoke) / var(--tw-text-opacity))}[data-cmdk-list]{height:min(300px, var(--cmdk-list-height));max-height:500px;overflow:auto;overscroll-behavior:contain;transition:height 0.1s ease 0s}[data-cmdk-list]::-webkit-scrollbar{width:0.5rem}[data-cmdk-list]::-webkit-scrollbar-track{background-color:transparent}[data-cmdk-list]::-webkit-scrollbar-thumb{background-color:rgb(var(--neutral-fg) / 0.2)}[data-cmdk-list]::-webkit-scrollbar-thumb:hover{background-color:rgb(var(--neutral-fg) / 0.4);background-color:rgb(var(--neutral-fg) / 0.8)}[data-cmdk-group]{margin-top:0.5rem}[data-cmdk-group-heading]{margin-bottom:0.5rem;display:flex;-webkit-user-select:none;-moz-user-select:none;user-select:none;align-items:center;padding-left:0.75rem;padding-right:0.75rem;font-size:0.75rem;line-height:1rem;--tw-text-opacity:1;color:rgb(var(--smoke) / var(--tw-text-opacity))}[data-cmdk-item]{content-visibility:auto;will-change:background, color;transition:none 0.15s ease 0s;position:relative;display:flex;height:3rem;cursor:pointer;-webkit-user-select:none;-moz-user-select:none;user-select:none;align-items:center;gap:0.75rem;padding-left:1rem;padding-right:1rem;font-size:0.875rem;line-height:1.25rem}[data-cmdk-item][data-selected="true"]{--tw-bg-opacity:1;background-color:rgb(var(--base2) / var(--tw-bg-opacity))}[data-cmdk-item][data-selected="true"]::after{position:absolute;left:0px;height:100%;width:0.25rem;--tw-bg-opacity:1;background-color:rgb(var(--primary) / var(--tw-bg-opacity));--tw-content:"";content:var(--tw-content)}[data-cmdk-item][data-selected="true"]:where([dir="rtl"], [dir="rtl"] *)::after{right:0px;left:auto}[data-cmdk-empty]{display:flex;height:4rem;align-items:center;justify-content:center;white-space:pre-wrap;text-align:center;font-size:0.875rem;line-height:1.25rem;--tw-text-opacity:1;color:rgb(var(--smoke) / var(--tw-text-opacity))}[data-cmdk-footer]{display:flex;height:2.5rem;width:100%;align-items:center;overflow:hidden;border-top-width:1px;--tw-border-opacity:1;border-color:rgb(var(--smoke) / var(--tw-border-opacity));background-color:rgb(var(--base2) / 0.5);padding:0.5rem}',
  map: '{"version":3,"file":"Search.svelte","sources":["Search.svelte"],"sourcesContent":["<script lang=\\"ts\\">import { Command } from \\"cmdk-sv\\";\\nlet open = false;\\nlet value = \\"\\";\\nexport const Open = () => {\\n  open = true;\\n};\\nexport const Close = () => {\\n  open = false;\\n};\\nexport const Toggle = () => {\\n  open != open;\\n};\\nexport let inputPlaceholder = \\"Type something to search...\\";\\nexport let noResultsLabel = \\"No results found.\\";\\n<\/script>\\r\\n\\r\\n<Command.Dialog\\r\\n    label=\\"Command Menu\\"\\r\\n    bind:value\\r\\n    bind:open\\r\\n>\\r\\n\\t<Command.Input placeholder={ inputPlaceholder } />\\r\\n\\r\\n\\t<Command.List>\\r\\n\\t\\t<Command.Empty>{ noResultsLabel }</Command.Empty>\\r\\n\\r\\n\\t\\t<Command.Group heading=\\"Letters\\">\\r\\n\\t\\t\\t<Command.Item>a</Command.Item>\\r\\n\\t\\t\\t<Command.Item>b</Command.Item>\\r\\n\\t\\t\\t<Command.Separator />\\r\\n\\t\\t\\t<Command.Item>c</Command.Item>\\r\\n\\t\\t</Command.Group>\\r\\n\\r\\n\\t\\t<Command.Item>Apple</Command.Item>\\r\\n\\t</Command.List>\\r\\n\\r\\n    <div data-cmdk-footer>\\r\\n\\r\\n    </div>\\r\\n</Command.Dialog>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    :global([data-cmdk-root]) {\\r\\n        width: 90%;\\r\\n        max-width: 42rem;\\r\\n        overflow: hidden;\\r\\n        --tw-bg-opacity: 1;\\r\\n        background-color: rgb(var(--base1) / var(--tw-bg-opacity));\\r\\n        --tw-text-opacity: 1;\\r\\n        color: rgb(var(--basec) / var(--tw-text-opacity));\\r\\n        --tw-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);\\r\\n        --tw-shadow-colored: 0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);\\r\\n        box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);\\r\\n        border-radius: var(--rounded-box);\\r\\n        position: fixed;\\r\\n        left: 50%;\\r\\n        top: 50%;\\r\\n        z-index: var(--dialog-z-index);\\r\\n        --tw-translate-x: -50%;\\r\\n        --tw-translate-y: -50%;\\r\\n        transform: translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y));\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-input]) {\\r\\n        width: 100%;\\r\\n        border-bottom-width: 1px;\\r\\n        --tw-border-opacity: 1;\\r\\n        border-bottom-color: rgb(var(--smoke) / var(--tw-border-opacity));\\r\\n        background-color: transparent;\\r\\n        padding: 1.25rem;\\r\\n        font-size: 1.125rem;\\r\\n        line-height: 1.75rem;\\r\\n        caret-color: rgb(var(--primary) / 1);\\r\\n        outline: 2px solid transparent;\\r\\n        outline-offset: 2px;\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-input])::-moz-placeholder {\\r\\n        --tw-text-opacity: 1;\\r\\n        color: rgb(var(--smoke) / var(--tw-text-opacity));\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-input])::placeholder {\\r\\n        --tw-text-opacity: 1;\\r\\n        color: rgb(var(--smoke) / var(--tw-text-opacity));\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-list]) {\\r\\n        height: min(300px, var(--cmdk-list-height));\\r\\n        max-height: 500px;\\r\\n        overflow: auto;\\r\\n        overscroll-behavior: contain;\\r\\n        transition: height 0.1s ease 0s;\\r\\n    }\\r\\n\\r\\n    :global([data-cmdk-list])::-webkit-scrollbar {\\r\\n        width: 0.5rem;\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-list])::-webkit-scrollbar-track {\\r\\n        background-color: transparent;\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-list])::-webkit-scrollbar-thumb {\\r\\n        background-color: rgb(var(--neutral-fg) / 0.2);\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-list])::-webkit-scrollbar-thumb:hover {\\r\\n        background-color: rgb(var(--neutral-fg) / 0.4);\\r\\n        background-color: rgb(var(--neutral-fg) / 0.8);\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-group]) {\\r\\n        margin-top: 0.5rem;\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-group-heading]) {\\r\\n        margin-bottom: 0.5rem;\\r\\n        display: flex;\\r\\n        -webkit-user-select: none;\\r\\n           -moz-user-select: none;\\r\\n                user-select: none;\\r\\n        align-items: center;\\r\\n        padding-left: 0.75rem;\\r\\n        padding-right: 0.75rem;\\r\\n        font-size: 0.75rem;\\r\\n        line-height: 1rem;\\r\\n        --tw-text-opacity: 1;\\r\\n        color: rgb(var(--smoke) / var(--tw-text-opacity));\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-item]) {\\r\\n        content-visibility: auto;\\r\\n        will-change: background, color;\\r\\n        transition: none 0.15s ease 0s;\\r\\n        position: relative;\\r\\n        display: flex;\\r\\n        height: 3rem;\\r\\n        cursor: pointer;\\r\\n        -webkit-user-select: none;\\r\\n           -moz-user-select: none;\\r\\n                user-select: none;\\r\\n        align-items: center;\\r\\n        gap: 0.75rem;\\r\\n        padding-left: 1rem;\\r\\n        padding-right: 1rem;\\r\\n        font-size: 0.875rem;\\r\\n        line-height: 1.25rem;\\r\\n    }\\r\\n\\r\\n    :global([data-cmdk-item][data-selected=\\"true\\"]) {\\r\\n        --tw-bg-opacity: 1;\\r\\n        background-color: rgb(var(--base2) / var(--tw-bg-opacity));\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-item][data-selected=\\"true\\"])::after {\\r\\n        position: absolute;\\r\\n        left: 0px;\\r\\n        height: 100%;\\r\\n        width: 0.25rem;\\r\\n        --tw-bg-opacity: 1;\\r\\n        background-color: rgb(var(--primary) / var(--tw-bg-opacity));\\r\\n        --tw-content: \\"\\";\\r\\n        content: var(--tw-content);\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-item][data-selected=\\"true\\"]):where([dir=\\"rtl\\"], [dir=\\"rtl\\"] *)::after {\\r\\n        right: 0px;\\r\\n        left: auto;\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-empty]) {\\r\\n        display: flex;\\r\\n        height: 4rem;\\r\\n        align-items: center;\\r\\n        justify-content: center;\\r\\n        white-space: pre-wrap;\\r\\n        text-align: center;\\r\\n        font-size: 0.875rem;\\r\\n        line-height: 1.25rem;\\r\\n        --tw-text-opacity: 1;\\r\\n        color: rgb(var(--smoke) / var(--tw-text-opacity));\\r\\n}\\r\\n\\r\\n    :global([data-cmdk-footer]) {\\r\\n        display: flex;\\r\\n        height: 2.5rem;\\r\\n        width: 100%;\\r\\n        align-items: center;\\r\\n        overflow: hidden;\\r\\n        border-top-width: 1px;\\r\\n        --tw-border-opacity: 1;\\r\\n        border-color: rgb(var(--smoke) / var(--tw-border-opacity));\\r\\n        background-color: rgb(var(--base2) / 0.5);\\r\\n        padding: 0.5rem;\\r\\n}\\r\\n</style>"],"names":[],"mappings":"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"}'
};
const Search = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let open = false;
  let value = "";
  const Open = () => {
    open = true;
  };
  const Close = () => {
    open = false;
  };
  const Toggle = () => {
  };
  let { inputPlaceholder = "Type something to search..." } = $$props;
  let { noResultsLabel = "No results found." } = $$props;
  if ($$props.Open === void 0 && $$bindings.Open && Open !== void 0) $$bindings.Open(Open);
  if ($$props.Close === void 0 && $$bindings.Close && Close !== void 0) $$bindings.Close(Close);
  if ($$props.Toggle === void 0 && $$bindings.Toggle && Toggle !== void 0) $$bindings.Toggle(Toggle);
  if ($$props.inputPlaceholder === void 0 && $$bindings.inputPlaceholder && inputPlaceholder !== void 0) $$bindings.inputPlaceholder(inputPlaceholder);
  if ($$props.noResultsLabel === void 0 && $$bindings.noResultsLabel && noResultsLabel !== void 0) $$bindings.noResultsLabel(noResultsLabel);
  $$result.css.add(css$3);
  let $$settled;
  let $$rendered;
  let previous_head = $$result.head;
  do {
    $$settled = true;
    $$result.head = previous_head;
    $$rendered = `${validate_component(CommandDialog, "Command.Dialog").$$render(
      $$result,
      { label: "Command Menu", value, open },
      {
        value: ($$value) => {
          value = $$value;
          $$settled = false;
        },
        open: ($$value) => {
          open = $$value;
          $$settled = false;
        }
      },
      {
        default: () => {
          return `${validate_component(CommandInput, "Command.Input").$$render($$result, { placeholder: inputPlaceholder }, {}, {})} ${validate_component(CommandList, "Command.List").$$render($$result, {}, {}, {
            default: () => {
              return `${validate_component(CommandEmpty, "Command.Empty").$$render($$result, {}, {}, {
                default: () => {
                  return `${escape(noResultsLabel)}`;
                }
              })} ${validate_component(CommandGroup, "Command.Group").$$render($$result, { heading: "Letters" }, {}, {
                default: () => {
                  return `${validate_component(CommandItem, "Command.Item").$$render($$result, {}, {}, {
                    default: () => {
                      return `a`;
                    }
                  })} ${validate_component(CommandItem, "Command.Item").$$render($$result, {}, {}, {
                    default: () => {
                      return `b`;
                    }
                  })} ${validate_component(CommandSeparator, "Command.Separator").$$render($$result, {}, {}, {})} ${validate_component(CommandItem, "Command.Item").$$render($$result, {}, {}, {
                    default: () => {
                      return `c`;
                    }
                  })}`;
                }
              })} ${validate_component(CommandItem, "Command.Item").$$render($$result, {}, {}, {
                default: () => {
                  return `Apple`;
                }
              })}`;
            }
          })} <div data-cmdk-footer data-svelte-h="svelte-1sncpe4"></div>`;
        }
      }
    )}`;
  } while (!$$settled);
  return $$rendered;
});
const css$2 = {
  code: '.drawer-overlay.svelte-189iqgo{position:fixed;inset:0px;z-index:var(--dialog-z-index);background-color:rgb(0 0 0 / 0.5)\n}.drawer-container.svelte-189iqgo{position:fixed;top:0px;z-index:var(--dialog-z-index);display:flex;height:100vh;width:100%;max-width:32rem;flex-direction:column;--tw-bg-opacity:1;background-color:rgb(var(--base1) / var(--tw-bg-opacity));--tw-shadow:0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);--tw-shadow-colored:0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)\n}.drawer-container.svelte-189iqgo:focus{outline:2px solid transparent;outline-offset:2px\n}.drawer-container.left.svelte-189iqgo{left:0px\n}.drawer-container.right.svelte-189iqgo{right:0px\n}.drawer-container__title.svelte-189iqgo{margin-bottom:0px;--tw-bg-opacity:1;background-color:rgb(var(--neutral) / var(--tw-bg-opacity));padding:1.5rem;font-size:1.125rem;line-height:1.75rem;font-weight:600;--tw-text-opacity:1;color:rgb(var(--neutral-fg) / var(--tw-text-opacity))\n}.drawer-container__close.svelte-189iqgo{position:absolute;right:10px;top:26px;display:inline-flex;height:1.5rem;width:1.5rem;-webkit-appearance:none;-moz-appearance:none;appearance:none;align-items:center;justify-content:center;border-radius:9999px;--tw-text-opacity:1;color:rgb(var(--neutral-fg) / var(--tw-text-opacity));transition-property:color, background-color, border-color, text-decoration-color, fill, stroke;transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1);transition-duration:150ms\n}.drawer-container__close.svelte-189iqgo:hover{background-color:rgb(var(--error) / 0.2);--tw-text-opacity:1;color:rgb(var(--error) / var(--tw-text-opacity))\n}.drawer-container__close.svelte-189iqgo:focus{--tw-text-opacity:1;color:rgb(var(--error) / var(--tw-text-opacity));outline:2px solid transparent;outline-offset:2px;--tw-ring-offset-shadow:var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color);--tw-ring-shadow:var(--tw-ring-inset) 0 0 0 calc(2px + var(--tw-ring-offset-width)) var(--tw-ring-color);box-shadow:var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000);--tw-ring-opacity:1;--tw-ring-color:rgb(var(--error) / var(--tw-ring-opacity))\n}.drawer-container__close.svelte-189iqgo:where([dir="rtl"], [dir="rtl"] *){left:10px;right:auto\n}',
  map: '{"version":3,"file":"Drawer.svelte","sources":["Drawer.svelte"],"sourcesContent":["<script lang=\\"ts\\">import { fade, fly } from \\"svelte/transition\\";\\nimport { createDialog, melt } from \\"@melt-ui/svelte\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nexport const Open = () => {\\n  $open = true;\\n};\\nexport const Close = () => {\\n  $open = false;\\n};\\nexport const Toggle = () => {\\n  $open != $open;\\n};\\nconst {\\n  elements: {\\n    overlay,\\n    content,\\n    title,\\n    close,\\n    portalled\\n  },\\n  states: { open }\\n} = createDialog({\\n  forceVisible: true\\n});\\nexport let direction = \\"right\\";\\n<\/script>\\r\\n\\r\\n{#if $open}\\r\\n    <div class=\\"\\" {...$portalled} use:$portalled.action>\\r\\n        <div\\r\\n            {...$overlay} use:$overlay.action\\r\\n            class=\\"drawer-overlay\\"\\r\\n            transition:fade={{ duration: 150 }}\\r\\n        />\\r\\n        <div\\r\\n            {...$content} use:$content.action\\r\\n            class={\\"drawer-container\\" + cn(direction)}\\r\\n            transition:fly={{\\r\\n                x: 512 * (direction == \\"left\\" ? -1 : 1),\\r\\n                duration: 300,\\r\\n                opacity: 1,\\r\\n            }}\\r\\n        >\\r\\n            <button\\r\\n                {...$close} use:$close.action\\r\\n                aria-label=\\"Close\\"\\r\\n                class=\\"drawer-container__close\\"\\r\\n            >\\r\\n                <i class=\\"fa-solid fa-xmark\\"></i>\\r\\n            </button>\\r\\n            <div {...$title} use:$title.action class=\\"drawer-container__title\\">\\r\\n                <slot name=\\"title\\" />\\r\\n            </div>\\r\\n            <section class=\\"p-6 max-h-full overflow-y-auto\\">\\r\\n                <slot />\\r\\n            </section>\\r\\n        </div>\\r\\n    </div>\\r\\n{/if}\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .drawer-overlay {\\r\\n    position: fixed;\\r\\n    inset: 0px;\\r\\n    z-index: var(--dialog-z-index);\\r\\n    background-color: rgb(0 0 0 / 0.5)\\n}\\r\\n    .drawer-container {\\r\\n    position: fixed;\\r\\n    top: 0px;\\r\\n    z-index: var(--dialog-z-index);\\r\\n    display: flex;\\r\\n    height: 100vh;\\r\\n    width: 100%;\\r\\n    max-width: 32rem;\\r\\n    flex-direction: column;\\r\\n    --tw-bg-opacity: 1;\\r\\n    background-color: rgb(var(--base1) / var(--tw-bg-opacity));\\r\\n    --tw-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);\\r\\n    --tw-shadow-colored: 0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);\\r\\n    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)\\n}\\r\\n    .drawer-container:focus {\\r\\n    outline: 2px solid transparent;\\r\\n    outline-offset: 2px\\n}\\r\\n    .drawer-container.left {\\r\\n    left: 0px\\n}\\r\\n    .drawer-container.right {\\r\\n    right: 0px\\n}\\r\\n    .drawer-container__title {\\r\\n    margin-bottom: 0px;\\r\\n    --tw-bg-opacity: 1;\\r\\n    background-color: rgb(var(--neutral) / var(--tw-bg-opacity));\\r\\n    padding: 1.5rem;\\r\\n    font-size: 1.125rem;\\r\\n    line-height: 1.75rem;\\r\\n    font-weight: 600;\\r\\n    --tw-text-opacity: 1;\\r\\n    color: rgb(var(--neutral-fg) / var(--tw-text-opacity))\\n}\\r\\n    .drawer-container__close {\\r\\n    position: absolute;\\r\\n    right: 10px;\\r\\n    top: 26px;\\r\\n    display: inline-flex;\\r\\n    height: 1.5rem;\\r\\n    width: 1.5rem;\\r\\n    -webkit-appearance: none;\\r\\n       -moz-appearance: none;\\r\\n            appearance: none;\\r\\n    align-items: center;\\r\\n    justify-content: center;\\r\\n    border-radius: 9999px;\\r\\n    --tw-text-opacity: 1;\\r\\n    color: rgb(var(--neutral-fg) / var(--tw-text-opacity));\\r\\n    transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;\\r\\n    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\\r\\n    transition-duration: 150ms\\n}\\r\\n    .drawer-container__close:hover {\\r\\n    background-color: rgb(var(--error) / 0.2);\\r\\n    --tw-text-opacity: 1;\\r\\n    color: rgb(var(--error) / var(--tw-text-opacity))\\n}\\r\\n    .drawer-container__close:focus {\\r\\n    --tw-text-opacity: 1;\\r\\n    color: rgb(var(--error) / var(--tw-text-opacity));\\r\\n    outline: 2px solid transparent;\\r\\n    outline-offset: 2px;\\r\\n    --tw-ring-offset-shadow: var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color);\\r\\n    --tw-ring-shadow: var(--tw-ring-inset) 0 0 0 calc(2px + var(--tw-ring-offset-width)) var(--tw-ring-color);\\r\\n    box-shadow: var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000);\\r\\n    --tw-ring-opacity: 1;\\r\\n    --tw-ring-color: rgb(var(--error) / var(--tw-ring-opacity))\\n}\\r\\n    .drawer-container__close:where([dir=\\"rtl\\"], [dir=\\"rtl\\"] *) {\\r\\n    left: 10px;\\r\\n    right: auto\\n}\\r\\n</style>"],"names":[],"mappings":"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"}'
};
const Drawer = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $open, $$unsubscribe_open;
  let $portalled, $$unsubscribe_portalled;
  let $overlay, $$unsubscribe_overlay;
  let $content, $$unsubscribe_content;
  let $close, $$unsubscribe_close;
  let $title, $$unsubscribe_title;
  const Open = () => {
    set_store_value(open, $open = true, $open);
  };
  const Close = () => {
    set_store_value(open, $open = false, $open);
  };
  const Toggle = () => {
  };
  const { elements: { overlay, content, title, close, portalled }, states: { open } } = createDialog$1({ forceVisible: true });
  $$unsubscribe_overlay = subscribe(overlay, (value) => $overlay = value);
  $$unsubscribe_content = subscribe(content, (value) => $content = value);
  $$unsubscribe_title = subscribe(title, (value) => $title = value);
  $$unsubscribe_close = subscribe(close, (value) => $close = value);
  $$unsubscribe_portalled = subscribe(portalled, (value) => $portalled = value);
  $$unsubscribe_open = subscribe(open, (value) => $open = value);
  let { direction = "right" } = $$props;
  if ($$props.Open === void 0 && $$bindings.Open && Open !== void 0) $$bindings.Open(Open);
  if ($$props.Close === void 0 && $$bindings.Close && Close !== void 0) $$bindings.Close(Close);
  if ($$props.Toggle === void 0 && $$bindings.Toggle && Toggle !== void 0) $$bindings.Toggle(Toggle);
  if ($$props.direction === void 0 && $$bindings.direction && direction !== void 0) $$bindings.direction(direction);
  $$result.css.add(css$2);
  $$unsubscribe_open();
  $$unsubscribe_portalled();
  $$unsubscribe_overlay();
  $$unsubscribe_content();
  $$unsubscribe_close();
  $$unsubscribe_title();
  return `${$open ? `<div${spread([{ class: "" }, escape_object($portalled)], { classes: "svelte-189iqgo" })}><div${spread([escape_object($overlay), { class: "drawer-overlay" }], { classes: "svelte-189iqgo" })}></div> <div${spread(
    [
      escape_object($content),
      {
        class: escape_attribute_value("drawer-container" + cn(direction))
      }
    ],
    { classes: "svelte-189iqgo" }
  )}><button${spread(
    [
      escape_object($close),
      { "aria-label": "Close" },
      { class: "drawer-container__close" }
    ],
    { classes: "svelte-189iqgo" }
  )}><i class="fa-solid fa-xmark"></i></button> <div${spread([escape_object($title), { class: "drawer-container__title" }], { classes: "svelte-189iqgo" })}>${slots.title ? slots.title({}) : ``}</div> <section class="p-6 max-h-full overflow-y-auto">${slots.default ? slots.default({}) : ``}</section></div></div>` : ``}`;
});
const css$1 = {
  code: ".kbd.svelte-1g9o8cd{--kdb-main-color:var(--base3);--kbd-text-color:var(--basec);--kbd-border-color:var(--neutral) / .2;display:inline-flex;min-height:2.25rem;min-width:2.25rem;align-items:center;justify-content:center;border-width:1px;border-bottom-width:2px;border-color:rgb(var(--kbd-border-color));background-color:rgb(var(--kdb-main-color));padding-left:0.5rem;padding-right:0.5rem;color:rgb(var(--kbd-text-color));border-radius:var(--rounded-btn)}.kbd.primary.svelte-1g9o8cd{--kdb-main-color:var(--primary);--kbd-text-color:var(--primary-fg)}.kbd.secondary.svelte-1g9o8cd{--kdb-main-color:var(--secondary);--kbd-text-color:var(--secondary-fg)}.kbd.neutral.svelte-1g9o8cd{--kdb-main-color:var(--neutral);--kbd-text-color:var(--neutral-fg)}.kbd.success.svelte-1g9o8cd{--kdb-main-color:var(--success);--kbd-text-color:var(--success-fg)}.kbd.info.svelte-1g9o8cd{--kdb-main-color:var(--info);--kbd-text-color:var(--info-fg)}.kbd.warning.svelte-1g9o8cd{--kdb-main-color:var(--warning);--kbd-text-color:var(--warning-fg)}.kbd.error.svelte-1g9o8cd{--kdb-main-color:var(--error);--kbd-text-color:var(--error-fg)}.kbd.smoke.svelte-1g9o8cd{--kdb-main-color:var(--smoke);--kbd-text-color:var(--smoke-fg)}",
  map: '{"version":3,"file":"Keyboard.svelte","sources":["Keyboard.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">export const DEFAULT = \\"\\";\\nexport const PRIMARY = \\"primary\\";\\nexport const SECONDARY = \\"secondary\\";\\nexport const NEUTRAL = \\"neutral\\";\\nexport const SUCCESS = \\"success\\";\\nexport const INFO = \\"info\\";\\nexport const WARNING = \\"warning\\";\\nexport const ERROR = \\"error\\";\\nexport const SMOKE = \\"smoke\\";\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { cn } from \\"@/lib/helpers\\";\\nlet className = \\"\\";\\nexport { className as class };\\nexport let char = \\"\\";\\nexport let palette = DEFAULT;\\nexport let style = \\"\\";\\n<\/script>\\r\\n\\r\\n<kbd class={\\"kbd\\" + cn(palette) + cn(className)} style={style}>{ char }</kbd>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .kbd {\\r\\n        --kdb-main-color: var(--base3);\\r\\n        --kbd-text-color: var(--basec);\\r\\n        --kbd-border-color: var(--neutral) / .2;\\r\\n        display: inline-flex;\\r\\n        min-height: 2.25rem;\\r\\n        min-width: 2.25rem;\\r\\n        align-items: center;\\r\\n        justify-content: center;\\r\\n        border-width: 1px;\\r\\n        border-bottom-width: 2px;\\r\\n        border-color: rgb(var(--kbd-border-color));\\r\\n        background-color: rgb(var(--kdb-main-color));\\r\\n        padding-left: 0.5rem;\\r\\n        padding-right: 0.5rem;\\r\\n        color: rgb(var(--kbd-text-color));\\r\\n        border-radius: var(--rounded-btn);\\r\\n    }\\r\\n\\r\\n        .kbd.primary {\\r\\n            --kdb-main-color: var(--primary);\\r\\n            --kbd-text-color: var(--primary-fg);\\r\\n        }\\r\\n\\r\\n        .kbd.secondary {\\r\\n            --kdb-main-color: var(--secondary);\\r\\n            --kbd-text-color: var(--secondary-fg);\\r\\n        }\\r\\n\\r\\n        .kbd.neutral {\\r\\n            --kdb-main-color: var(--neutral);\\r\\n            --kbd-text-color: var(--neutral-fg);\\r\\n        }\\r\\n\\r\\n        .kbd.success {\\r\\n            --kdb-main-color: var(--success);\\r\\n            --kbd-text-color: var(--success-fg);\\r\\n        }\\r\\n\\r\\n        .kbd.info {\\r\\n            --kdb-main-color: var(--info);\\r\\n            --kbd-text-color: var(--info-fg);\\r\\n        }\\r\\n\\r\\n        .kbd.warning {\\r\\n            --kdb-main-color: var(--warning);\\r\\n            --kbd-text-color: var(--warning-fg);\\r\\n        }\\r\\n\\r\\n        .kbd.error {\\r\\n            --kdb-main-color: var(--error);\\r\\n            --kbd-text-color: var(--error-fg);\\r\\n        }\\r\\n\\r\\n        .kbd.smoke {\\r\\n            --kdb-main-color: var(--smoke);\\r\\n            --kbd-text-color: var(--smoke-fg);\\r\\n        }\\r\\n</style>"],"names":[],"mappings":"AAsBI,mBAAK,CACD,gBAAgB,CAAE,YAAY,CAC9B,gBAAgB,CAAE,YAAY,CAC9B,kBAAkB,CAAE,mBAAmB,CACvC,OAAO,CAAE,WAAW,CACpB,UAAU,CAAE,OAAO,CACnB,SAAS,CAAE,OAAO,CAClB,WAAW,CAAE,MAAM,CACnB,eAAe,CAAE,MAAM,CACvB,YAAY,CAAE,GAAG,CACjB,mBAAmB,CAAE,GAAG,CACxB,YAAY,CAAE,IAAI,IAAI,kBAAkB,CAAC,CAAC,CAC1C,gBAAgB,CAAE,IAAI,IAAI,gBAAgB,CAAC,CAAC,CAC5C,YAAY,CAAE,MAAM,CACpB,aAAa,CAAE,MAAM,CACrB,KAAK,CAAE,IAAI,IAAI,gBAAgB,CAAC,CAAC,CACjC,aAAa,CAAE,IAAI,aAAa,CACpC,CAEI,IAAI,uBAAS,CACT,gBAAgB,CAAE,cAAc,CAChC,gBAAgB,CAAE,iBACtB,CAEA,IAAI,yBAAW,CACX,gBAAgB,CAAE,gBAAgB,CAClC,gBAAgB,CAAE,mBACtB,CAEA,IAAI,uBAAS,CACT,gBAAgB,CAAE,cAAc,CAChC,gBAAgB,CAAE,iBACtB,CAEA,IAAI,uBAAS,CACT,gBAAgB,CAAE,cAAc,CAChC,gBAAgB,CAAE,iBACtB,CAEA,IAAI,oBAAM,CACN,gBAAgB,CAAE,WAAW,CAC7B,gBAAgB,CAAE,cACtB,CAEA,IAAI,uBAAS,CACT,gBAAgB,CAAE,cAAc,CAChC,gBAAgB,CAAE,iBACtB,CAEA,IAAI,qBAAO,CACP,gBAAgB,CAAE,YAAY,CAC9B,gBAAgB,CAAE,eACtB,CAEA,IAAI,qBAAO,CACP,gBAAgB,CAAE,YAAY,CAC9B,gBAAgB,CAAE,eACtB"}'
};
const DEFAULT = "";
const Keyboard = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { class: className = "" } = $$props;
  let { char = "" } = $$props;
  let { palette = DEFAULT } = $$props;
  let { style = "" } = $$props;
  if ($$props.class === void 0 && $$bindings.class && className !== void 0) $$bindings.class(className);
  if ($$props.char === void 0 && $$bindings.char && char !== void 0) $$bindings.char(char);
  if ($$props.palette === void 0 && $$bindings.palette && palette !== void 0) $$bindings.palette(palette);
  if ($$props.style === void 0 && $$bindings.style && style !== void 0) $$bindings.style(style);
  $$result.css.add(css$1);
  return `<kbd class="${escape(null_to_empty("kbd" + cn(palette) + cn(className)), true) + " svelte-1g9o8cd"}"${add_attribute("style", style, 0)}>${escape(char)}</kbd>`;
});
const defaultComponentStyle = "data:image/svg+xml,%3csvg%20width='106'%20height='69'%20viewBox='0%200%20106%2069'%20fill='none'%20xmlns='http://www.w3.org/2000/svg'%3e%3crect%20y='0.217621'%20width='106'%20height='68'%20rx='4'%20fill='%234B465C'%20fill-opacity='0.02'/%3e%3cpath%20d='M0%204.21762C0%202.00848%201.79086%200.217621%204%200.217621H28V68.2176H4C1.79086%2068.2176%200%2066.4268%200%2064.2176V4.21762Z'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='5'%20y='24.8253'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='24.8253'%20width='18'%20height='2.87399'%20rx='1.43699'%20stroke='%23DBDADE'/%3e%3crect%20x='9'%20y='6.27713'%20width='10'%20height='10'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='9'%20y='6.27713'%20width='10'%20height='10'%20rx='2'%20stroke='%23DBDADE'/%3e%3crect%20x='5'%20y='35.6993'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='35.6993'%20width='18'%20height='2.87399'%20rx='1.43699'%20stroke='%23DBDADE'/%3e%3crect%20x='5'%20y='46.5733'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='46.5733'%20width='18'%20height='2.87399'%20rx='1.43699'%20stroke='%23DBDADE'/%3e%3crect%20x='5'%20y='57.4472'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='57.4472'%20width='18'%20height='2.87399'%20rx='1.43699'%20stroke='%23DBDADE'/%3e%3crect%20x='34.7715'%20y='5.03091'%20width='66'%20height='9.06667'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='37.752'%20y='7.29752'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='37.752'%20y='7.29752'%20width='4'%20height='4.53333'%20rx='1'%20stroke='%23DBDADE'/%3e%3crect%20x='81.752'%20y='7.29752'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='81.752'%20y='7.29752'%20width='4'%20height='4.53333'%20rx='1'%20stroke='%23DBDADE'/%3e%3crect%20x='87.752'%20y='7.29752'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='87.752'%20y='7.29752'%20width='4'%20height='4.53333'%20rx='1'%20stroke='%23DBDADE'/%3e%3crect%20x='93.752'%20y='7.29752'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='93.752'%20y='7.29752'%20width='4'%20height='4.53333'%20rx='1'%20stroke='%23DBDADE'/%3e%3crect%20x='59.6094'%20y='20.4254'%20width='41'%20height='18.1333'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='34.7715'%20y='20.4254'%20width='19.4118'%20height='18.1333'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='34.7715'%20y='43.9586'%20width='66'%20height='18.1333'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3c/svg%3e";
const borderedComponentStyle = "data:image/svg+xml,%3csvg%20width='107'%20height='69'%20viewBox='0%200%20107%2069'%20fill='none'%20xmlns='http://www.w3.org/2000/svg'%3e%3crect%20x='0.390625'%20y='0.284363'%20width='106'%20height='68'%20rx='4'%20fill='%234B465C'%20fill-opacity='0.02'/%3e%3crect%20x='5.39062'%20y='24.8921'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='9.39062'%20y='6.34387'%20width='10'%20height='10'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5.39062'%20y='35.7661'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5.39062'%20y='46.64'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5.39062'%20y='57.514'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='35.6621'%20y='5.59766'%20width='65'%20height='8.06667'%20rx='1.5'%20stroke='%23DBDADE'/%3e%3crect%20x='38.1426'%20y='7.36426'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='82.1426'%20y='7.36426'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='88.1426'%20y='7.36426'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='93.1426'%20y='7.36426'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='60.5'%20y='20.9922'%20width='40'%20height='17.1333'%20rx='1.5'%20stroke='%23DBDADE'/%3e%3crect%20x='35.5'%20y='20.5'%20width='18.4118'%20height='17.1333'%20rx='1.5'%20stroke='%23DBDADE'/%3e%3crect%20x='35.6621'%20y='44.5254'%20width='65'%20height='17.1333'%20rx='1.5'%20stroke='%23DBDADE'/%3e%3c/svg%3e";
const stickyNav = "data:image/svg+xml,%3csvg%20width='106'%20height='69'%20viewBox='0%200%20106%2069'%20fill='none'%20xmlns='http://www.w3.org/2000/svg'%3e%3crect%20y='0.568726'%20width='106'%20height='68'%20rx='4'%20fill='%234B465C'%20fill-opacity='0.02'/%3e%3cpath%20d='M0%204.56873C0%202.35959%201.79086%200.568726%204%200.568726H28V68.5687H4C1.79086%2068.5687%200%2066.7779%200%2064.5687V4.56873Z'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='5'%20y='25.1765'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='9'%20y='6.62823'%20width='10'%20height='10'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='36.0504'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='46.9244'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='57.7983'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3cpath%20d='M32.7715%200.284363H98.7715V7.35103C98.7715%208.4556%2097.8761%209.35103%2096.7715%209.35103H34.7715C33.6669%209.35103%2032.7715%208.4556%2032.7715%207.35103V0.284363Z'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='35.752'%20y='2.55096'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='79.752'%20y='2.55096'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='85.752'%20y='2.55096'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='91.752'%20y='2.55096'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='58.1836'%20y='13'%20width='41'%20height='18.1333'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='32.7715'%20y='13'%20width='19.4118'%20height='18.1333'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='33'%20y='36'%20width='66'%20height='28'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3c/svg%3e";
const staticNav = "data:image/svg+xml,%3csvg%20width='106'%20height='69'%20viewBox='0%200%20106%2069'%20fill='none'%20xmlns='http://www.w3.org/2000/svg'%3e%3crect%20y='0.568726'%20width='106'%20height='68'%20rx='4'%20fill='%234B465C'%20fill-opacity='0.02'/%3e%3cpath%20d='M0%204.56873C0%202.35959%201.79086%200.568726%204%200.568726H28V68.5687H4C1.79086%2068.5687%200%2066.7779%200%2064.5687V4.56873Z'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='5'%20y='25.1765'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='9'%20y='6.62823'%20width='10'%20height='10'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='36.0504'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='46.9244'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='57.7983'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='32.7715'%20y='5.38202'%20width='66'%20height='9.06667'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='35.752'%20y='7.64862'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='79.752'%20y='7.64862'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='85.752'%20y='7.64862'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='91.752'%20y='7.64862'%20width='4'%20height='4.53333'%20rx='1'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='58.1836'%20y='20.7766'%20width='41'%20height='18.1333'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='32.7715'%20y='20.7766'%20width='19.4118'%20height='18.1333'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='32.7715'%20y='44.3098'%20width='66.4121'%20height='18.1333'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3c/svg%3e";
const hiddenNav = "data:image/svg+xml,%3csvg%20width='106'%20height='69'%20viewBox='0%200%20106%2069'%20fill='none'%20xmlns='http://www.w3.org/2000/svg'%3e%3crect%20y='0.568726'%20width='106'%20height='68'%20rx='4'%20fill='%234B465C'%20fill-opacity='0.02'/%3e%3cpath%20d='M0%204.56873C0%202.35959%201.79086%200.568726%204%200.568726H28V68.5687H4C1.79086%2068.5687%200%2066.7779%200%2064.5687V4.56873Z'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='5'%20y='25.1765'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='9'%20y='6.62823'%20width='10'%20height='10'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='36.0504'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='46.9244'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5'%20y='57.7983'%20width='18'%20height='2.87399'%20rx='1.43699'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='58.1836'%20y='6.28436'%20width='41'%20height='18.1333'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='32.7715'%20y='6.28436'%20width='19.4118'%20height='18.1333'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='33'%20y='30.2844'%20width='66'%20height='32'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3c/svg%3e";
const ltrImage = "data:image/svg+xml,%3csvg%20width='106'%20height='69'%20viewBox='0%200%20106%2069'%20fill='none'%20xmlns='http://www.w3.org/2000/svg'%3e%3crect%20y='0.284424'%20width='106'%20height='68'%20rx='4'%20fill='%234B465C'%20fill-opacity='0.02'/%3e%3crect%20x='5.30176'%20y='4.53076'%20width='24.561'%20height='59.3336'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='10.5225'%20y='17.6654'%20width='14.1209'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='10.5225'%20y='17.6654'%20width='14.1209'%20height='2.15549'%20rx='1.07775'%20stroke='%23DBDADE'/%3e%3crect%20x='10.5225'%20y='26.6208'%20width='10.0694'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='10.5225'%20y='26.6208'%20width='10.0694'%20height='2.15549'%20rx='1.07775'%20stroke='%23DBDADE'/%3e%3crect%20x='10.5225'%20y='35.5763'%20width='12.6207'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='10.5225'%20y='35.5763'%20width='12.6207'%20height='2.15549'%20rx='1.07775'%20stroke='%23DBDADE'/%3e%3crect%20x='10.5225'%20y='44.5319'%20width='6.20526'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='10.5225'%20y='44.5319'%20width='6.20526'%20height='2.15549'%20rx='1.07775'%20stroke='%23DBDADE'/%3e%3crect%20x='10.5225'%20y='53.4875'%20width='8.24653'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='10.5225'%20y='53.4875'%20width='8.24653'%20height='2.15549'%20rx='1.07775'%20stroke='%23DBDADE'/%3e%3crect%20x='36.1953'%20y='4.53076'%20width='63.5883'%20height='59.3336'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='44.5977'%20y='14.8976'%20width='14.1209'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='44.5977'%20y='14.8976'%20width='14.1209'%20height='2.15549'%20rx='1.07775'%20stroke='%23DBDADE'/%3e%3crect%20x='44.5977'%20y='23.853'%20width='33.432'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='44.5977'%20y='23.853'%20width='33.432'%20height='2.15549'%20rx='1.07775'%20stroke='%23DBDADE'/%3e%3crect%20x='44.5977'%20y='32.8085'%20width='42'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='44.5977'%20y='32.8085'%20width='42'%20height='2.15549'%20rx='1.07775'%20stroke='%23DBDADE'/%3e%3crect%20x='44.5977'%20y='41.7638'%20width='33.432'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='44.5977'%20y='41.7638'%20width='33.432'%20height='2.15549'%20rx='1.07775'%20stroke='%23DBDADE'/%3e%3crect%20x='44.5977'%20y='50.7196'%20width='5.88587'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='44.5977'%20y='50.7196'%20width='5.88587'%20height='2.15549'%20rx='1.07775'%20stroke='%23DBDADE'/%3e%3c/svg%3e";
const rtlImage = "data:image/svg+xml,%3csvg%20width='106'%20height='69'%20viewBox='0%200%20106%2069'%20fill='none'%20xmlns='http://www.w3.org/2000/svg'%3e%3crect%20y='0.284424'%20width='106'%20height='68'%20rx='4'%20fill='%234B465C'%20fill-opacity='0.02'/%3e%3crect%20x='74.8896'%20y='4.53076'%20width='24.561'%20height='59.3336'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='80.1104'%20y='17.6654'%20width='14.1209'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='84.1611'%20y='26.6208'%20width='10.0694'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='81.6104'%20y='35.5763'%20width='12.6207'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='88.0254'%20y='44.5319'%20width='6.20526'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='85.9844'%20y='53.4875'%20width='8.24653'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='5.30176'%20y='4.53076'%20width='63.5883'%20height='59.3336'%20rx='2'%20fill='%234B465C'%20fill-opacity='0.08'/%3e%3crect%20x='46.5879'%20y='14.8976'%20width='14.1209'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='27.2773'%20y='23.853'%20width='33.432'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='18.709'%20y='32.8085'%20width='42'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='27.2773'%20y='41.7638'%20width='33.432'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3crect%20x='54.8232'%20y='50.7196'%20width='5.88587'%20height='2.15549'%20rx='1.07775'%20fill='%234B465C'%20fill-opacity='0.16'/%3e%3c/svg%3e";
const css = {
  code: '.page.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{display:flex;height:100vh;width:100vw;flex-direction:row;overflow:hidden}.page__navigation.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{--navigation-width:200px;position:absolute;top:0px;bottom:0px;z-index:20;display:flex;height:100%;width:5rem;max-width:var(--navigation-width);flex-direction:column;--tw-bg-opacity:1;background-color:rgb(var(--neutral) / var(--tw-bg-opacity));--tw-text-opacity:1;color:rgb(var(--neutral-fg) / var(--tw-text-opacity))}@media(min-width: 768px){.page__navigation.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{position:relative;z-index:10}}.page__navigation.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{transition-property:width;transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1);transition-duration:300ms}.page__navigation__header.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{--logo-size:40px;display:flex;height:3.5rem;-webkit-user-select:none;-moz-user-select:none;user-select:none;align-items:center;justify-content:center;gap:0.5rem;overflow:hidden;padding-left:0.5rem;padding-right:0.5rem;padding-top:0.5rem;padding-bottom:0.5rem}.page__navigation__header.svelte-yqfnma>.logo.svelte-yqfnma.svelte-yqfnma{height:100%;max-height:var(--logo-size);width:100%;max-width:var(--logo-size)}.page__navigation__header.svelte-yqfnma>.logo.svelte-yqfnma>.svelte-yqfnma{height:var(--logo-size);width:var(--logo-size)}.page__navigation__list.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{margin-top:1rem;max-height:100%;overflow-y:auto;overflow-x:hidden}.page__navigation__list.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma::-webkit-scrollbar{width:0.5rem}.page__navigation__list.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma::-webkit-scrollbar-track{background-color:transparent}.page__navigation__list.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma::-webkit-scrollbar-thumb{background-color:rgb(var(--neutral-fg) / 0.2)}.page__navigation__list.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma::-webkit-scrollbar-thumb:hover{background-color:rgb(var(--neutral-fg) / 0.4);background-color:rgb(var(--neutral-fg) / 0.8)}.page__navigation__list__link.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{--icon-size:25px;display:flex;width:100%;align-items:center;justify-content:center;gap:0.5rem;overflow:hidden;padding-left:0.5rem;padding-right:0.5rem;padding-top:0.75rem;padding-bottom:0.75rem}.page__navigation__list__link.svelte-yqfnma>.icon.svelte-yqfnma.svelte-yqfnma{display:flex;height:100%;max-height:var(--icon-size);width:100%;max-width:var(--icon-size);align-items:center;justify-content:center}.page__navigation__list__link.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma:hover{background-color:rgb(0 0 0 / 0.1)}.page__navigation.toggle.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{width:var(--navigation-width)}.page__navigation.toggle.svelte-yqfnma .page__navigation__list__link.svelte-yqfnma.svelte-yqfnma{justify-content:flex-start}.page__container.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{z-index:10;display:flex;height:100%;width:100%;flex-direction:column;overflow:hidden;--tw-bg-opacity:1;background-color:rgb(var(--base2) / var(--tw-bg-opacity));padding-left:5rem}@media(min-width: 768px){.page__container.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{padding-left:0px}}.page__container.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma:where([dir="rtl"], [dir="rtl"] *){padding-right:5rem;padding-left:0px}@media(min-width: 768px){.page__container.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma:where([dir="rtl"], [dir="rtl"] *){padding-right:0px}}.page__header.separated.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{padding:1rem}.page__header.separated.svelte-yqfnma>.svelte-yqfnma.svelte-yqfnma:first-child{border-radius:var(--rounded-box)}.page__header__content.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{display:flex;width:100%;flex-direction:row;align-items:center;justify-content:space-between;--tw-bg-opacity:1;background-color:rgb(var(--base1) / var(--tw-bg-opacity));padding-left:1rem;padding-right:1rem;padding-top:0.5rem;padding-bottom:0.5rem;--tw-shadow:0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);--tw-shadow-colored:0 4px 6px -1px var(--tw-shadow-color), 0 2px 4px -2px var(--tw-shadow-color);box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)}.page__content.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{display:flex;height:100%;width:100%;flex-direction:column;overflow:auto;--tw-bg-opacity:1;background-color:rgb(var(--base2) / var(--tw-bg-opacity));padding:1rem}.shortcut.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma{--icon-color:var(--base3);display:flex;min-width:100px;-webkit-user-select:none;-moz-user-select:none;user-select:none;flex-direction:column;align-items:center;justify-content:space-between;gap:0.5rem;padding-top:1.25rem;padding-bottom:1.25rem;transition-property:all;transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1);transition-duration:150ms;border-radius:var(--rounded-box)}.shortcut.svelte-yqfnma>.svelte-yqfnma.svelte-yqfnma:first-child{color:rgb(var(--icon-color))}.shortcut.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma:hover{background-color:rgb(0 0 0 / 0.05)}.shortcut.svelte-yqfnma.svelte-yqfnma.svelte-yqfnma:active:active{--tw-scale-x:.9;--tw-scale-y:.9;transform:translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))}',
  map: '{"version":3,"file":"Navigation.svelte","sources":["Navigation.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">export const SHOW_LOGO_WHEN_EXPANDED = 1;\\nexport const SEPARATED_PAGE_HEADER = 2;\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { getContext, onMount, setContext, tick } from \\"svelte\\";\\nimport { readable, writable } from \\"svelte/store\\";\\nimport { fade, scale } from \\"svelte/transition\\";\\nimport NavigationLinkComponent from \\"./NavigationLink.svelte\\";\\nimport NotificationsBlock from \\"./NotificationsBlock.svelte\\";\\nimport CustomizerOption from \\"./CustomizerOption.svelte\\";\\nimport { Search } from \\"@/components/Search\\";\\nimport { Drawer } from \\"@/components/Drawer\\";\\nimport { Button, CIRCLE, GHOST, NO_ANIMATION, NO_RIPPLE, PLAIN } from \\"@/widgets/actions/button\\";\\nimport { LanguageSelect } from \\"@/widgets/actions/language-select\\";\\nimport { Keyboard } from \\"@/widgets/data-displays/keyboard\\";\\nimport { DISABLE_CLOSE_BTN, Popover } from \\"@/components/Popover\\";\\nimport { GetText, Text } from \\"@/components/i18n/Text\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nimport { SettingsEnity } from \\"@/lib/LocalStorageEntity\\";\\nimport defaultComponentStyle from \\"@/assets/images/customizer/default-style.svg\\";\\nimport borderedComponentStyle from \\"@/assets/images/customizer/bordered-style.svg\\";\\nimport stickyNav from \\"@/assets/images/customizer/sticky-navbar.svg\\";\\nimport staticNav from \\"@/assets/images/customizer/static-navbar.svg\\";\\nimport hiddenNav from \\"@/assets/images/customizer/hidden-navbar.svg\\";\\nimport ltrImage from \\"@/assets/images/customizer/ltr.svg\\";\\nimport rtlImage from \\"@/assets/images/customizer/rtl.svg\\";\\nimport { navigating } from \\"$app/stores\\";\\nimport { browser } from \\"$app/environment\\";\\nconst i18nLoadDictionary = getContext(\\"i18nLoadDictionary\\");\\nconst isApple = browser && (navigator.platform.indexOf(\\"Mac\\") === 0 || navigator.platform === \\"iPhone\\" || navigator.platform === \\"iPad\\");\\nconst findPage = function(pages, url) {\\n  const currentPage = pages.find((lnk) => lnk.href == url);\\n  if (currentPage) return currentPage;\\n  for (var i = 0; i < pages.length; i++) {\\n    const lnk = pages[i];\\n    if (!lnk.children) continue;\\n    const currentPage2 = findPage(lnk.children, url.replace(lnk.href, \\"\\"));\\n    if (currentPage2) return currentPage2;\\n  }\\n};\\nlet routeMeta = writable({\\n  url: \\"\\",\\n  title: \\"\\"\\n});\\nconst UpdateRouteMeta = function() {\\n  $routeMeta.url = $navigating?.to?.url.pathname ?? \\"\\";\\n  const currentPage = findPage(links, $routeMeta.url);\\n  $routeMeta.title = currentPage?.title ?? \\"\\";\\n};\\n$: if ($navigating) UpdateRouteMeta();\\nlet loaded = false;\\nlet dictionary;\\nlet dictionaryStore;\\nlet title = \\"\\";\\nconst SetPageTitle = function() {\\n  if (!dictionaryStore) return;\\n  title = GetText(\\"dashboard.sidebar.title\\", $dictionaryStore);\\n};\\nlet settingsEntity;\\nlet settings = {\\n  theme: \\"\\",\\n  sidebarImage: \\"\\",\\n  direction: \\"ltr\\",\\n  navbarApperance: \\"sticky\\",\\n  navbarPosition: \\"attached\\",\\n  componentStyle: \\"default\\"\\n};\\nlet settingsErrors = {\\n  theme: \\"\\",\\n  sidebarImage: \\"\\",\\n  direction: \\"\\",\\n  navbarApperance: \\"\\",\\n  navbarPosition: \\"\\",\\n  componentStyle: \\"\\"\\n};\\nconst UpdateSettingsEntity = () => {\\n  if (!settingsEntity || !loaded) return;\\n  settingsEntity.fromObject(settings);\\n  settingsErrors = Object.fromEntries(Object.entries(settingsEntity.ValidateEntity()).map((v) => v[1].status == \\"fail\\" ? [v[0], v[1].description] : void 0).filter((v) => v != void 0));\\n  settingsEntity.Save();\\n  window.location.reload();\\n};\\n$: settings, UpdateSettingsEntity();\\nlet drawer;\\nlet drawerDirection = \\"right\\";\\nlet search;\\nconst OpenSearch = function(e) {\\n  if (e) {\\n    if (e.key == \\"f\\" && e.ctrlKey) {\\n      e.preventDefault();\\n      search.Open();\\n    }\\n    return;\\n  }\\n  search.Open();\\n};\\nlet search_i18n = {\\n  inputPlaceholder: \\"\\",\\n  noResultsLabel: \\"\\"\\n};\\nconst SetSearchText = function() {\\n  search_i18n.inputPlaceholder = GetText(\\"dashboard.navbar.search.input.placeholder\\", $dictionaryStore) || \\"Type something to search...\\";\\n  search_i18n.noResultsLabel = GetText(\\"dashboard.navbar.search.no_results_found\\", $dictionaryStore) || \\"No results found\\";\\n};\\nlet toggleMenu = writable(false);\\nsetContext(\\"toggleMenu\\", readable(false, (set) => toggleMenu.subscribe((val) => set(val))));\\nlet fullscreen = false;\\nconst ToggleFullscreen = function() {\\n  if (!browser) return;\\n  if (fullscreen) document.documentElement.requestFullscreen();\\n  else if (document.fullscreenElement) document.exitFullscreen();\\n};\\n$: fullscreen, ToggleFullscreen();\\nexport let links = [];\\nexport let flags = 0;\\nexport let shortcuts = [];\\nonMount(() => {\\n  $routeMeta.url = window.location.pathname;\\n  const currentPage = findPage(links, $routeMeta.url);\\n  $routeMeta.title = currentPage?.title ?? \\"\\";\\n});\\nonMount(async () => {\\n  dictionary = i18nLoadDictionary([\\"dashboard.sidebar\\", \\"dashboard.navbar\\", \\"dashboard.theme_customizer\\"]);\\n  dictionary.then((res) => {\\n    dictionaryStore = res;\\n    SetPageTitle();\\n    SetSearchText();\\n  });\\n});\\nonMount(async () => {\\n  settingsEntity = SettingsEnity.GetInstance().Load();\\n  settings = settingsEntity.toObject();\\n  {\\n    if (settings.direction == \\"rtl\\") drawerDirection = \\"left\\";\\n  }\\n  {\\n    if (settingsEntity.GetProperty(\\"expandNavbar\\") == true) $toggleMenu = true;\\n  }\\n  await tick();\\n  loaded = true;\\n});\\nonMount(() => {\\n  if (browser) {\\n    document.documentElement.removeEventListener(\\"keydown\\", OpenSearch);\\n    document.documentElement.addEventListener(\\"keydown\\", OpenSearch);\\n  }\\n});\\n<\/script>\\r\\n\\r\\n<svelte:head>\\r\\n    <title>{title}</title>\\r\\n</svelte:head>\\r\\n\\r\\n<Search inputPlaceholder={search_i18n.inputPlaceholder} noResultsLabel={search_i18n.noResultsLabel} bind:this={search} />\\r\\n\\r\\n<Drawer direction={drawerDirection} bind:this={drawer}>\\r\\n    <svelte:fragment slot=\\"title\\">\\r\\n        <span>\\r\\n            <Text key=\\"dashboard.theme_customizer.title\\" source={dictionary} />\\r\\n        </span>\\r\\n    </svelte:fragment>\\r\\n    <div class=\\"flex flex-col divide-y-2 divide-base-200\\">\\r\\n        <div class=\\"flex flex-col py-2\\">\\r\\n            <span class=\\"font-medium text-lg\\">\\r\\n                <Text key=\\"dashboard.theme_customizer.settings.theme.title\\" source={dictionary} />\\r\\n            </span>\\r\\n            <span class=\\"text-neutral-400 text-sm\\">\\r\\n                <Text key=\\"dashboard.theme_customizer.settings.theme.description\\" source={dictionary} />\\r\\n            </span>\\r\\n        </div>\\r\\n        <div class=\\"flex flex-col py-2\\">\\r\\n            <span class=\\"font-medium text-lg\\">\\r\\n                <Text key=\\"dashboard.theme_customizer.settings.sidebar_images.title\\" source={dictionary} />\\r\\n            </span>\\r\\n            <span class=\\"text-neutral-400 text-sm\\">\\r\\n                <Text key=\\"dashboard.theme_customizer.settings.sidebar_images.description\\" source={dictionary} />\\r\\n            </span>\\r\\n        </div>\\r\\n        <div class=\\"flex flex-col py-2\\">\\r\\n            <span class=\\"font-medium text-lg\\">\\r\\n                <Text key=\\"dashboard.theme_customizer.settings.components.title\\" source={dictionary} />\\r\\n            </span>\\r\\n            <span class=\\"text-neutral-400 text-sm\\">\\r\\n                <Text key=\\"dashboard.theme_customizer.settings.components.description\\" source={dictionary} />\\r\\n            </span>\\r\\n            <div class=\\"flex flex-row flex-wrap gap-3 mt-5\\" role=\\"radiogroup\\">\\r\\n                <CustomizerOption value=\\"default\\" bind:userSelect={settings.componentStyle}>\\r\\n                    <svelte:fragment slot=\\"body\\">\\r\\n                        <img src={defaultComponentStyle} alt=\\"\\">\\r\\n                    </svelte:fragment>\\r\\n                    <svelte:fragment slot=\\"label\\">\\r\\n                        <Text key=\\"dashboard.theme_customizer.settings.components.select.default.title\\" source={dictionary} />\\r\\n                    </svelte:fragment>\\r\\n                </CustomizerOption>\\r\\n                <CustomizerOption value=\\"bordered\\" bind:userSelect={settings.componentStyle}>\\r\\n                    <svelte:fragment slot=\\"body\\">\\r\\n                        <img src={borderedComponentStyle} alt=\\"\\">\\r\\n                    </svelte:fragment>\\r\\n                    <svelte:fragment slot=\\"label\\">\\r\\n                        <Text key=\\"dashboard.theme_customizer.settings.components.select.bordered.title\\" source={dictionary} />\\r\\n                    </svelte:fragment>\\r\\n                </CustomizerOption>\\r\\n            </div>\\r\\n        </div>\\r\\n        <div class=\\"flex flex-col py-2\\">\\r\\n            <span class=\\"font-medium text-lg\\">\\r\\n                <Text key=\\"dashboard.theme_customizer.settings.navbar_type.title\\" source={dictionary} />\\r\\n            </span>\\r\\n            <span class=\\"text-neutral-400 text-sm\\">\\r\\n                <Text key=\\"dashboard.theme_customizer.settings.navbar_type.description\\" source={dictionary} />\\r\\n            </span>\\r\\n            <div class=\\"flex flex-row flex-wrap gap-3 mt-5\\" role=\\"radiogroup\\">\\r\\n                <CustomizerOption value=\\"sticky\\" bind:userSelect={settings.navbarApperance}>\\r\\n                    <svelte:fragment slot=\\"body\\">\\r\\n                        <img src={stickyNav} alt=\\"\\">\\r\\n                    </svelte:fragment>\\r\\n                    <svelte:fragment slot=\\"label\\">\\r\\n                        <Text key=\\"dashboard.theme_customizer.settings.navbar_type.select.sticky.title\\" source={dictionary} />\\r\\n                    </svelte:fragment>\\r\\n                </CustomizerOption>\\r\\n                <CustomizerOption value=\\"static\\" bind:userSelect={settings.navbarApperance}>\\r\\n                    <svelte:fragment slot=\\"body\\">\\r\\n                        <img src={staticNav} alt=\\"\\">\\r\\n                    </svelte:fragment>\\r\\n                    <svelte:fragment slot=\\"label\\">\\r\\n                        <Text key=\\"dashboard.theme_customizer.settings.navbar_type.select.static.title\\" source={dictionary} />\\r\\n                    </svelte:fragment>\\r\\n                </CustomizerOption>\\r\\n                <CustomizerOption value=\\"hidden\\" bind:userSelect={settings.navbarApperance}>\\r\\n                    <svelte:fragment slot=\\"body\\">\\r\\n                        <img src={hiddenNav} alt=\\"\\">\\r\\n                    </svelte:fragment>\\r\\n                    <svelte:fragment slot=\\"label\\">\\r\\n                        <Text key=\\"dashboard.theme_customizer.settings.navbar_type.select.hidden.title\\" source={dictionary} />\\r\\n                    </svelte:fragment>\\r\\n                </CustomizerOption>\\r\\n            </div>\\r\\n            <div class=\\"flex flex-row flex-wrap gap-3 mt-5\\" role=\\"radiogroup\\">\\r\\n                <CustomizerOption value=\\"attached\\" bind:userSelect={settings.navbarPosition}>\\r\\n                    <svelte:fragment slot=\\"body\\">\\r\\n                        <img src={stickyNav} alt=\\"\\">\\r\\n                    </svelte:fragment>\\r\\n                    <svelte:fragment slot=\\"label\\">\\r\\n                        <Text key=\\"dashboard.theme_customizer.settings.navbar_type.select.attached.title\\" source={dictionary} />\\r\\n                    </svelte:fragment>\\r\\n                </CustomizerOption>\\r\\n                <CustomizerOption value=\\"separated\\" bind:userSelect={settings.navbarPosition}>\\r\\n                    <svelte:fragment slot=\\"body\\">\\r\\n                        <img src={stickyNav} alt=\\"\\">\\r\\n                    </svelte:fragment>\\r\\n                    <svelte:fragment slot=\\"label\\">\\r\\n                        <Text key=\\"dashboard.theme_customizer.settings.navbar_type.select.separated.title\\" source={dictionary} />\\r\\n                    </svelte:fragment>\\r\\n                </CustomizerOption>\\r\\n            </div>\\r\\n        </div>\\r\\n        <div class=\\"flex flex-col py-2\\">\\r\\n            <span class=\\"font-medium text-lg\\">\\r\\n                <Text key=\\"dashboard.theme_customizer.settings.direction.title\\" source={dictionary} />\\r\\n            </span>\\r\\n            <span class=\\"text-neutral-400 text-sm\\">\\r\\n                <Text key=\\"dashboard.theme_customizer.settings.direction.description\\" source={dictionary} />\\r\\n            </span>\\r\\n            <div class=\\"flex flex-row flex-wrap gap-3 mt-5\\" role=\\"radiogroup\\">\\r\\n                <CustomizerOption value=\\"ltr\\" bind:userSelect={settings.direction}>\\r\\n                    <svelte:fragment slot=\\"body\\">\\r\\n                        <img src={ltrImage} alt=\\"\\">\\r\\n                    </svelte:fragment>\\r\\n                    <svelte:fragment slot=\\"label\\">\\r\\n                        <Text key=\\"dashboard.theme_customizer.settings.direction.select.ltr.title\\" source={dictionary} />\\r\\n                    </svelte:fragment>\\r\\n                </CustomizerOption>\\r\\n                <CustomizerOption value=\\"rtl\\" bind:userSelect={settings.direction}>\\r\\n                    <svelte:fragment slot=\\"body\\">\\r\\n                        <img src={rtlImage} alt=\\"\\">\\r\\n                    </svelte:fragment>\\r\\n                    <svelte:fragment slot=\\"label\\">\\r\\n                        <Text key=\\"dashboard.theme_customizer.settings.direction.select.rtl.title\\" source={dictionary} />\\r\\n                    </svelte:fragment>\\r\\n                </CustomizerOption>\\r\\n            </div>\\r\\n        </div>\\r\\n    </div>\\r\\n</Drawer>\\r\\n\\r\\n<div class=\\"page\\">\\r\\n    <div class=\\"page__navigation\\" class:toggle={$toggleMenu}>\\r\\n        <div class=\\"page__navigation__header\\">\\r\\n            {#if $toggleMenu && (flags & SHOW_LOGO_WHEN_EXPANDED) || !$toggleMenu}\\r\\n                <div class=\\"logo\\" in:fade>\\r\\n                    <svg viewBox=\\"0 0 128 128\\" xmlns=\\"http://www.w3.org/2000/svg\\" xmlns:xlink=\\"http://www.w3.org/1999/xlink\\" aria-hidden=\\"true\\" role=\\"img\\" preserveAspectRatio=\\"xMidYMid meet\\" fill=\\"#000000\\">\\r\\n                        <g stroke-width=\\"0\\"></g>\\r\\n                        <g stroke-linecap=\\"round\\" stroke-linejoin=\\"round\\"></g>\\r\\n                        <g>\\r\\n                            <path d=\\"M118.89 75.13a15.693 15.693 0 0 0-7-7.33a22.627 22.627 0 0 0-6-2.63c1.53-5.6-.64-10.06-3.69-13.39c-4.51-4.88-9.2-5.59-9.2-5.59c1.62-3.07 2.11-6.61 1.36-10c-.77-3.69-3.08-6.86-6.36-8.72c-3.1-1.83-6.92-2.73-10.84-3.47c-1.88-.34-9.81-1.45-13.1-6c-2.65-3.69-2.73-10.33-3.45-12.32s-3.38-1.15-6.23.76C51.05 8.7 44.15 15.83 41.49 23a24.6 24.6 0 0 0-1.28 13.89c-2.14.35-4.23.97-6.21 1.85c-.16 0-.32.1-.49.17c-3 1.24-9.43 7-10 15.85c-.21 3.13.19 6.26 1.17 9.24c-2.19.57-4.3 1.43-6.26 2.57c-2.29.98-4.38 2.38-6.15 4.13c-1.95 2.41-3.37 5.2-4.15 8.2a27.594 27.594 0 0 0 2 19.77c1.8 3.47 4.06 6.67 6.74 9.52c8.55 8.79 23.31 12.11 35 14c14.19 2.34 29.05 1.52 42.33-4c19.92-8.22 25.22-21.44 26-25.17c1.73-8.25-.39-16.02-1.3-17.89z\\" fill=\\"#885742\\"></path>\\r\\n                            <path d=\\"M87.45 92.89c-1.57.8-3.17 1.52-4.78 2.16c-1.08.43-2.17.82-3.27 1.17c-1.1.36-2.21.67-3.33 1c-2.24.56-4.52.97-6.82 1.21c-1.74.19-3.5.28-5.25.28c-4.62 0-9.22-.65-13.67-1.91l-1.46-.44a55.12 55.12 0 0 1-7.15-2.84l-1.39-.69a22.722 22.722 0 0 0 12.72 15.31c3.43 1.59 7.17 2.4 10.95 2.38c3.82.03 7.6-.75 11.09-2.31a21.868 21.868 0 0 0 12.58-15.44l-.22.12z\\" fill=\\"#35220b\\"></path>\\r\\n                            <path d=\\"M85.19 90c-7 1.23-14.09 1.82-21.19 1.77c-7.1.04-14.19-.55-21.19-1.77a2.16 2.16 0 0 0-2.53 2.54v.25A51.578 51.578 0 0 0 64 98.66c1.75 0 3.51-.09 5.25-.28c2.3-.24 4.58-.65 6.82-1.21c1.12-.28 2.23-.59 3.33-1s2.19-.74 3.27-1.17c1.62-.67 3.21-1.39 4.78-2.16l.22-.12l.06-.27c.17-1.19-.66-2.29-1.86-2.46a2.22 2.22 0 0 0-.68.01z\\" fill=\\"#ffffff\\"></path>\\r\\n                            <g>\\r\\n                                <circle cx=\\"80.13\\" cy=\\"69.49\\" r=\\"12.4\\" fill=\\"#ffffff\\"></circle>\\r\\n                                <ellipse cx=\\"80.13\\" cy=\\"69.49\\" rx=\\"5.73\\" ry=\\"5.82\\" fill=\\"#35220b\\"></ellipse>\\r\\n                                <circle cx=\\"47.87\\" cy=\\"69.49\\" r=\\"12.4\\" fill=\\"#ffffff\\"></circle>\\r\\n                                <ellipse cx=\\"47.87\\" cy=\\"69.49\\" rx=\\"5.73\\" ry=\\"5.82\\" fill=\\"#35220b\\"></ellipse>\\r\\n                            </g>\\r\\n                        </g>\\r\\n                    </svg>\\r\\n                </div>\\r\\n            {/if}\\r\\n            {#if $toggleMenu}\\r\\n                <p class=\\"text-center mx-auto md:text-lg\\" in:fade>SM Box</p>\\r\\n\\r\\n                <Button \\r\\n                    class=\\"!p-1 ml-auto rtl:mr-auto rtl:ml-0 w-10 h-10 md:!hidden\\"\\r\\n                    palette={GHOST} \\r\\n                    flags={NO_RIPPLE | NO_ANIMATION}\\r\\n                    OnClick={() => {\\r\\n                        $toggleMenu = false\\r\\n                        settingsEntity.SetProperty(\\"expandNavbar\\", false).Save()\\r\\n                    }}\\r\\n                >\\r\\n                    <i class=\\"fa-solid fa-xmark\\"></i>\\r\\n                </Button>\\r\\n            {/if}\\r\\n        </div>\\r\\n        <div class=\\"page__navigation__list\\">\\r\\n            {#each links as link, idx (idx)}\\r\\n                <NavigationLinkComponent data={link}/>\\r\\n            {/each}\\r\\n        </div>\\r\\n    </div>\\r\\n    <div class=\\"page__container\\">\\r\\n        <div class={\\"page__header\\" + cn(flags & SEPARATED_PAGE_HEADER ? \\"separated\\" : \\"\\")}>\\r\\n            <div class=\\"page__header__content\\">\\r\\n                <Button \\r\\n                    class=\\"!p-1 w-10 h-10\\"\\r\\n                    palette={GHOST} \\r\\n                    flags={NO_RIPPLE | NO_ANIMATION}\\r\\n                    OnClick={() => {\\r\\n                        $toggleMenu = !$toggleMenu\\r\\n                        settingsEntity.SetProperty(\\"expandNavbar\\", $toggleMenu).Save()\\r\\n                    }}\\r\\n                >\\r\\n                    {#if $toggleMenu}\\r\\n                        <i class=\\"absolute fa-solid fa-xmark fa-xl\\" transition:scale></i>\\r\\n                    {:else}\\r\\n                        <i class=\\"absolute fa-solid fa-bars fa-xl\\" transition:scale></i>\\r\\n                    {/if}\\r\\n                </Button>\\r\\n                <Button\\r\\n                    class=\\"!bg-base-200 ml-5 !hidden md:!flex\\"\\r\\n                    palette={GHOST}\\r\\n                    flags={NO_RIPPLE | NO_ANIMATION}\\r\\n                    OnClick={() => OpenSearch()}\\r\\n                >\\r\\n                    <span class=\\"fa-solid fa-magnifying-glass mr-2\\"></span>\\r\\n                    <div class=\\"flex flex-row items-center gap-1\\">\\r\\n                        {#if isApple}\\r\\n                            <Keyboard char=\\"⌘\\" />\\r\\n                        {:else}\\r\\n                            <Keyboard char=\\"ctrl\\" />\\r\\n                        {/if}\\r\\n                        +\\r\\n                        <Keyboard char=\\"f\\" />\\r\\n                    </div>\\r\\n                </Button>\\r\\n                <div class=\\"ml-auto rtl:mr-auto rtl:ml-0 flex items-center flex-wrap\\">\\r\\n                    <LanguageSelect />\\r\\n                    <Button\\r\\n                        class=\\"!p-1 w-10 h-10 !text-neutral-400 active:!text-primary md:!hidden\\"\\r\\n                        flags={NO_RIPPLE | PLAIN | CIRCLE}\\r\\n                        OnClick={() => search.Open()}\\r\\n                    >\\r\\n                        <span class=\\"fa-solid fa-magnifying-glass\\"></span>\\r\\n                    </Button>\\r\\n                    <Popover class=\\"!p-0\\" flags={DISABLE_CLOSE_BTN}>\\r\\n                        <div slot=\\"trigger\\" let:trigger let:melt {...trigger} use:trigger.action>\\r\\n                            <Button\\r\\n                                class=\\"!p-1 w-10 h-10 !text-neutral-400 active:!text-primary\\"\\r\\n                                flags={NO_RIPPLE | PLAIN | CIRCLE}\\r\\n                            >\\r\\n                                <svg xmlns=\\"http://www.w3.org/2000/svg\\" width=\\"24\\" height=\\"24\\" viewBox=\\"0 0 24 24\\" style=\\"fill: currentColor\\">\\r\\n                                    <path d=\\"M10 3H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zM9 9H5V5h4v4zm11 4h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1zm-1 6h-4v-4h4v4zM17 3c-2.206 0-4 1.794-4 4s1.794 4 4 4 4-1.794 4-4-1.794-4-4-4zm0 6c-1.103 0-2-.897-2-2s.897-2 2-2 2 .897 2 2-.897 2-2 2zM7 13c-2.206 0-4 1.794-4 4s1.794 4 4 4 4-1.794 4-4-1.794-4-4-4zm0 6c-1.103 0-2-.897-2-2s.897-2 2-2 2 .897 2 2-.897 2-2 2z\\"></path>\\r\\n                                </svg>\\r\\n                            </Button>\\r\\n                        </div>\\r\\n                        <div class=\\"flex flex-col\\">\\r\\n                            <div class=\\"flex flex-row items-center justify-between col-span-3 border-b border-base-300 border-dashed py-4 px-5\\">\\r\\n                                <span class=\\"text-xl font-semibold\\">\\r\\n                                    <Text key=\\"dashboard.navbar.shortcuts.title\\" source={dictionary} />\\r\\n                                </span>\\r\\n                            </div>\\r\\n                            <div class=\\"grid grid-cols-3 p-2\\">\\r\\n                                {#each shortcuts as shortcut, idx (idx)}\\r\\n                                    <a target={shortcut.target ?? \\"_self\\"} rel=\\"noreferrer\\" href={shortcut.link} class={\\"shortcut\\" + cn(shortcut.iconPalette ?? \\"\\")}>\\r\\n                                        {#if shortcut.icon}\\r\\n                                            {#if shortcut.iconType == \\"custom\\"}\\r\\n                                                <img width=\\"24\\" height=\\"24\\" src={shortcut.icon} alt=\\"\\" style={shortcut.style ?? \\"\\"}>\\r\\n                                            {:else}\\r\\n                                                <span class={\\"text-2xl\\" + cn(shortcut.icon)} style={shortcut.style ?? \\"\\"}></span>\\r\\n                                            {/if}\\r\\n                                        {:else}\\r\\n                                            <span class={\\"fa-regular fa-circle\\"}></span>\\r\\n                                        {/if}\\r\\n                                        <span class=\\"font-semibold truncate\\">{shortcut.title}</span>\\r\\n                                    </a>\\r\\n                                {/each}\\r\\n                            </div>\\r\\n                        </div>\\r\\n                    </Popover>\\r\\n                    <Button\\r\\n                        class=\\"!p-1 w-10 h-10 !text-neutral-400 active:!text-primary\\"\\r\\n                        flags={NO_RIPPLE | PLAIN | CIRCLE}\\r\\n                        OnClick={() => fullscreen = !fullscreen}\\r\\n                    >\\r\\n                        {#if fullscreen}\\r\\n                            <svg xmlns=\\"http://www.w3.org/2000/svg\\" width=\\"24\\" height=\\"24\\" viewBox=\\"0 0 24 24\\" style=\\"fill: currentColor\\">\\r\\n                                <path d=\\"M10 4H8v4H4v2h6zM8 20h2v-6H4v2h4zm12-6h-6v6h2v-4h4zm0-6h-4V4h-2v6h6z\\"></path>\\r\\n                            </svg>\\r\\n                        {:else}\\r\\n                            <svg xmlns=\\"http://www.w3.org/2000/svg\\" width=\\"24\\" height=\\"24\\" viewBox=\\"0 0 24 24\\" style=\\"fill: currentColor\\">\\r\\n                                <path d=\\"M5 5h5V3H3v7h2zm5 14H5v-5H3v7h7zm11-5h-2v5h-5v2h7zm-2-4h2V3h-7v2h5z\\"></path>\\r\\n                            </svg>\\r\\n                        {/if}\\r\\n                    </Button>\\r\\n                    <Popover class=\\"!p-0\\" flags={DISABLE_CLOSE_BTN}>\\r\\n                        <div slot=\\"trigger\\" let:trigger let:melt {...trigger} use:trigger.action>\\r\\n                            <Button\\r\\n                                class=\\"!p-1 w-10 h-10 !text-neutral-400 active:!text-primary\\"\\r\\n                                flags={NO_RIPPLE | PLAIN | CIRCLE}\\r\\n                            >\\r\\n                                <svg xmlns=\\"http://www.w3.org/2000/svg\\" width=\\"24\\" height=\\"24\\" viewBox=\\"0 0 24 24\\" style=\\"fill: currentColor\\">\\r\\n                                    <path d=\\"M19 13.586V10c0-3.217-2.185-5.927-5.145-6.742C13.562 2.52 12.846 2 12 2s-1.562.52-1.855 1.258C7.185 4.074 5 6.783 5 10v3.586l-1.707 1.707A.996.996 0 0 0 3 16v2a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-2a.996.996 0 0 0-.293-.707L19 13.586zM19 17H5v-.586l1.707-1.707A.996.996 0 0 0 7 14v-4c0-2.757 2.243-5 5-5s5 2.243 5 5v4c0 .266.105.52.293.707L19 16.414V17zm-7 5a2.98 2.98 0 0 0 2.818-2H9.182A2.98 2.98 0 0 0 12 22z\\"></path>\\r\\n                                </svg>\\r\\n                            </Button>\\r\\n                        </div>\\r\\n                        <NotificationsBlock dictionary={dictionary} />\\r\\n                    </Popover>\\r\\n                </div>\\r\\n            </div>\\r\\n        </div>\\r\\n        <div class=\\"page__content\\">\\r\\n            <div class=\\"z-[100] absolute top-1/2 -translate-y-1/2 right-0 rtl:left-0 rtl:right-auto\\">\\r\\n                <Button \\r\\n                    flags={NO_ANIMATION} \\r\\n                    class=\\"!w-10 !h-10 !rounded-r-none rtl:!rounded-l-none rtl:!rounded-r-btn hover:*:animate-spin\\"\\r\\n                    OnClick={() => drawer.Open()}\\r\\n                >\\r\\n                    <i class=\\"fa-solid fa-gear\\"></i>\\r\\n                </Button>\\r\\n            </div>\\r\\n            <slot meta={$routeMeta} />\\r\\n        </div>\\r\\n    </div>\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .page {\\r\\n            display: flex;\\r\\n            height: 100vh;\\r\\n            width: 100vw;\\r\\n            flex-direction: row;\\r\\n            overflow: hidden;\\r\\n}\\r\\n\\r\\n        .page__navigation {\\r\\n            --navigation-width: 200px;\\r\\n            position: absolute;\\r\\n            top: 0px;\\r\\n            bottom: 0px;\\r\\n            z-index: 20;\\r\\n            display: flex;\\r\\n            height: 100%;\\r\\n            width: 5rem;\\r\\n            max-width: var(--navigation-width);\\r\\n            flex-direction: column;\\r\\n            --tw-bg-opacity: 1;\\r\\n            background-color: rgb(var(--neutral) / var(--tw-bg-opacity));\\r\\n            --tw-text-opacity: 1;\\r\\n            color: rgb(var(--neutral-fg) / var(--tw-text-opacity));\\r\\n        }\\r\\n\\r\\n        @media (min-width: 768px) {\\r\\n\\r\\n            .page__navigation {\\r\\n                        position: relative;\\r\\n                        z-index: 10;\\r\\n            }\\r\\n}\\r\\n\\r\\n        .page__navigation {\\r\\n\\r\\n            transition-property: width;\\r\\n            transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\\r\\n            transition-duration: 300ms;\\r\\n        }\\r\\n\\r\\n        .page__navigation__header {\\r\\n                --logo-size: 40px;\\r\\n                display: flex;\\r\\n                height: 3.5rem;\\r\\n                -webkit-user-select: none;\\r\\n                   -moz-user-select: none;\\r\\n                        user-select: none;\\r\\n                align-items: center;\\r\\n                justify-content: center;\\r\\n                gap: 0.5rem;\\r\\n                overflow: hidden;\\r\\n                padding-left: 0.5rem;\\r\\n                padding-right: 0.5rem;\\r\\n                padding-top: 0.5rem;\\r\\n                padding-bottom: 0.5rem;\\r\\n            }\\r\\n\\r\\n        .page__navigation__header > .logo {\\r\\n            height: 100%;\\r\\n            max-height: var(--logo-size);\\r\\n            width: 100%;\\r\\n            max-width: var(--logo-size);\\r\\n}\\r\\n\\r\\n        .page__navigation__header > .logo > * {\\r\\n            height: var(--logo-size);\\r\\n            width: var(--logo-size);\\r\\n}\\r\\n\\r\\n        .page__navigation__list {\\r\\n            margin-top: 1rem;\\r\\n            max-height: 100%;\\r\\n            overflow-y: auto;\\r\\n            overflow-x: hidden;\\r\\n}\\r\\n\\r\\n        .page__navigation__list::-webkit-scrollbar {\\r\\n            width: 0.5rem;\\r\\n}\\r\\n\\r\\n        .page__navigation__list::-webkit-scrollbar-track {\\r\\n            background-color: transparent;\\r\\n}\\r\\n\\r\\n        .page__navigation__list::-webkit-scrollbar-thumb {\\r\\n            background-color: rgb(var(--neutral-fg) / 0.2);\\r\\n}\\r\\n\\r\\n        .page__navigation__list::-webkit-scrollbar-thumb:hover {\\r\\n            background-color: rgb(var(--neutral-fg) / 0.4);\\r\\n            background-color: rgb(var(--neutral-fg) / 0.8);\\r\\n}\\r\\n\\r\\n        .page__navigation__list__link {\\r\\n                    --icon-size: 25px;\\r\\n                    display: flex;\\r\\n                    width: 100%;\\r\\n                    align-items: center;\\r\\n                    justify-content: center;\\r\\n                    gap: 0.5rem;\\r\\n                    overflow: hidden;\\r\\n                    padding-left: 0.5rem;\\r\\n                    padding-right: 0.5rem;\\r\\n                    padding-top: 0.75rem;\\r\\n                    padding-bottom: 0.75rem;\\r\\n                }\\r\\n\\r\\n        .page__navigation__list__link > .icon {\\r\\n            display: flex;\\r\\n            height: 100%;\\r\\n            max-height: var(--icon-size);\\r\\n            width: 100%;\\r\\n            max-width: var(--icon-size);\\r\\n            align-items: center;\\r\\n            justify-content: center;\\r\\n}\\r\\n\\r\\n        .page__navigation__list__link:hover {\\r\\n            background-color: rgb(0 0 0 / 0.1);\\r\\n}\\r\\n\\r\\n        .page__navigation.toggle {\\r\\n            width: var(--navigation-width);\\r\\n}\\r\\n\\r\\n        .page__navigation.toggle .page__navigation__list__link {\\r\\n            justify-content: flex-start;\\r\\n}\\r\\n\\r\\n        .page__container {\\r\\n            z-index: 10;\\r\\n            display: flex;\\r\\n            height: 100%;\\r\\n            width: 100%;\\r\\n            flex-direction: column;\\r\\n            overflow: hidden;\\r\\n            --tw-bg-opacity: 1;\\r\\n            background-color: rgb(var(--base2) / var(--tw-bg-opacity));\\r\\n            padding-left: 5rem;\\r\\n}\\r\\n\\r\\n        @media (min-width: 768px) {\\r\\n\\r\\n            .page__container {\\r\\n                        padding-left: 0px;\\r\\n            }\\r\\n}\\r\\n\\r\\n        .page__container:where([dir=\\"rtl\\"], [dir=\\"rtl\\"] *) {\\r\\n            padding-right: 5rem;\\r\\n            padding-left: 0px;\\r\\n}\\r\\n\\r\\n        @media (min-width: 768px) {\\r\\n\\r\\n            .page__container:where([dir=\\"rtl\\"], [dir=\\"rtl\\"] *) {\\r\\n                        padding-right: 0px;\\r\\n            }\\r\\n}\\r\\n\\r\\n        .page__header.separated {\\r\\n            padding: 1rem;\\r\\n}\\r\\n\\r\\n        .page__header.separated > :first-child {\\r\\n            border-radius: var(--rounded-box);\\r\\n}\\r\\n\\r\\n        .page__header__content {\\r\\n            display: flex;\\r\\n            width: 100%;\\r\\n            flex-direction: row;\\r\\n            align-items: center;\\r\\n            justify-content: space-between;\\r\\n            --tw-bg-opacity: 1;\\r\\n            background-color: rgb(var(--base1) / var(--tw-bg-opacity));\\r\\n            padding-left: 1rem;\\r\\n            padding-right: 1rem;\\r\\n            padding-top: 0.5rem;\\r\\n            padding-bottom: 0.5rem;\\r\\n            --tw-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);\\r\\n            --tw-shadow-colored: 0 4px 6px -1px var(--tw-shadow-color), 0 2px 4px -2px var(--tw-shadow-color);\\r\\n            box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);\\r\\n}\\r\\n\\r\\n        .page__content {\\r\\n            display: flex;\\r\\n            height: 100%;\\r\\n            width: 100%;\\r\\n            flex-direction: column;\\r\\n            overflow: auto;\\r\\n            --tw-bg-opacity: 1;\\r\\n            background-color: rgb(var(--base2) / var(--tw-bg-opacity));\\r\\n            padding: 1rem;\\r\\n}\\r\\n\\r\\n    .shortcut {\\r\\n        --icon-color: var(--base3);\\r\\n        display: flex;\\r\\n        min-width: 100px;\\r\\n        -webkit-user-select: none;\\r\\n           -moz-user-select: none;\\r\\n                user-select: none;\\r\\n        flex-direction: column;\\r\\n        align-items: center;\\r\\n        justify-content: space-between;\\r\\n        gap: 0.5rem;\\r\\n        padding-top: 1.25rem;\\r\\n        padding-bottom: 1.25rem;\\r\\n        transition-property: all;\\r\\n        transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\\r\\n        transition-duration: 150ms;\\r\\n        border-radius: var(--rounded-box);\\r\\n    }\\r\\n\\r\\n    .shortcut > :first-child {\\r\\n            color: rgb(var(--icon-color));\\r\\n}\\r\\n\\r\\n    .shortcut:hover {\\r\\n            background-color: rgb(0 0 0 / 0.05);\\r\\n}\\r\\n\\r\\n    .shortcut:active:active {\\r\\n            --tw-scale-x: .9;\\r\\n            --tw-scale-y: .9;\\r\\n            transform: translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y));\\r\\n}\\r\\n</style>"],"names":[],"mappings":"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"}'
};
const SHOW_LOGO_WHEN_EXPANDED = 1;
const SEPARATED_PAGE_HEADER = 2;
const Navigation = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $toggleMenu, $$unsubscribe_toggleMenu;
  let $routeMeta, $$unsubscribe_routeMeta;
  let $navigating, $$unsubscribe_navigating;
  $$unsubscribe_navigating = subscribe(navigating, (value) => $navigating = value);
  getContext("i18nLoadDictionary");
  const findPage = function(pages, url) {
    const currentPage = pages.find((lnk) => lnk.href == url);
    if (currentPage) return currentPage;
    for (var i = 0; i < pages.length; i++) {
      const lnk = pages[i];
      if (!lnk.children) continue;
      const currentPage2 = findPage(lnk.children, url.replace(lnk.href, ""));
      if (currentPage2) return currentPage2;
    }
  };
  let routeMeta = writable({ url: "", title: "" });
  $$unsubscribe_routeMeta = subscribe(routeMeta, (value) => $routeMeta = value);
  const UpdateRouteMeta = function() {
    set_store_value(routeMeta, $routeMeta.url = $navigating?.to?.url.pathname ?? "", $routeMeta);
    const currentPage = findPage(links, $routeMeta.url);
    set_store_value(routeMeta, $routeMeta.title = currentPage?.title ?? "", $routeMeta);
  };
  let dictionary;
  let title = "";
  let settingsEntity;
  let settings = {
    theme: "",
    sidebarImage: "",
    direction: "ltr",
    navbarApperance: "sticky",
    navbarPosition: "attached",
    componentStyle: "default"
  };
  let drawer;
  let drawerDirection = "right";
  let search;
  const OpenSearch = function(e) {
    search.Open();
  };
  let search_i18n = { inputPlaceholder: "", noResultsLabel: "" };
  let toggleMenu = writable(false);
  $$unsubscribe_toggleMenu = subscribe(toggleMenu, (value) => $toggleMenu = value);
  setContext("toggleMenu", readable(false, (set) => toggleMenu.subscribe((val) => set(val))));
  let fullscreen = false;
  let { links = [] } = $$props;
  let { flags = 0 } = $$props;
  let { shortcuts = [] } = $$props;
  if ($$props.links === void 0 && $$bindings.links && links !== void 0) $$bindings.links(links);
  if ($$props.flags === void 0 && $$bindings.flags && flags !== void 0) $$bindings.flags(flags);
  if ($$props.shortcuts === void 0 && $$bindings.shortcuts && shortcuts !== void 0) $$bindings.shortcuts(shortcuts);
  $$result.css.add(css);
  let $$settled;
  let $$rendered;
  let previous_head = $$result.head;
  do {
    $$settled = true;
    $$result.head = previous_head;
    {
      if ($navigating) UpdateRouteMeta();
    }
    $$rendered = `${$$result.head += `<!-- HEAD_svelte-1uo06u1_START -->${$$result.title = `<title>${escape(title)}</title>`, ""}<!-- HEAD_svelte-1uo06u1_END -->`, ""} ${validate_component(Search, "Search").$$render(
      $$result,
      {
        inputPlaceholder: search_i18n.inputPlaceholder,
        noResultsLabel: search_i18n.noResultsLabel,
        this: search
      },
      {
        this: ($$value) => {
          search = $$value;
          $$settled = false;
        }
      },
      {}
    )} ${validate_component(Drawer, "Drawer").$$render(
      $$result,
      { direction: drawerDirection, this: drawer },
      {
        this: ($$value) => {
          drawer = $$value;
          $$settled = false;
        }
      },
      {
        title: () => {
          return `<span>${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.title",
              source: dictionary
            },
            {},
            {}
          )}</span> `;
        },
        default: () => {
          return `<div class="flex flex-col divide-y-2 divide-base-200"><div class="flex flex-col py-2"><span class="font-medium text-lg">${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.settings.theme.title",
              source: dictionary
            },
            {},
            {}
          )}</span> <span class="text-neutral-400 text-sm">${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.settings.theme.description",
              source: dictionary
            },
            {},
            {}
          )}</span></div> <div class="flex flex-col py-2"><span class="font-medium text-lg">${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.settings.sidebar_images.title",
              source: dictionary
            },
            {},
            {}
          )}</span> <span class="text-neutral-400 text-sm">${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.settings.sidebar_images.description",
              source: dictionary
            },
            {},
            {}
          )}</span></div> <div class="flex flex-col py-2"><span class="font-medium text-lg">${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.settings.components.title",
              source: dictionary
            },
            {},
            {}
          )}</span> <span class="text-neutral-400 text-sm">${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.settings.components.description",
              source: dictionary
            },
            {},
            {}
          )}</span> <div class="flex flex-row flex-wrap gap-3 mt-5" role="radiogroup">${validate_component(CustomizerOption, "CustomizerOption").$$render(
            $$result,
            {
              value: "default",
              userSelect: settings.componentStyle
            },
            {
              userSelect: ($$value) => {
                settings.componentStyle = $$value;
                $$settled = false;
              }
            },
            {
              label: () => {
                return `${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.theme_customizer.settings.components.select.default.title",
                    source: dictionary
                  },
                  {},
                  {}
                )} `;
              },
              body: () => {
                return `<img${add_attribute("src", defaultComponentStyle, 0)} alt="">`;
              }
            }
          )} ${validate_component(CustomizerOption, "CustomizerOption").$$render(
            $$result,
            {
              value: "bordered",
              userSelect: settings.componentStyle
            },
            {
              userSelect: ($$value) => {
                settings.componentStyle = $$value;
                $$settled = false;
              }
            },
            {
              label: () => {
                return `${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.theme_customizer.settings.components.select.bordered.title",
                    source: dictionary
                  },
                  {},
                  {}
                )} `;
              },
              body: () => {
                return `<img${add_attribute("src", borderedComponentStyle, 0)} alt="">`;
              }
            }
          )}</div></div> <div class="flex flex-col py-2"><span class="font-medium text-lg">${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.settings.navbar_type.title",
              source: dictionary
            },
            {},
            {}
          )}</span> <span class="text-neutral-400 text-sm">${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.settings.navbar_type.description",
              source: dictionary
            },
            {},
            {}
          )}</span> <div class="flex flex-row flex-wrap gap-3 mt-5" role="radiogroup">${validate_component(CustomizerOption, "CustomizerOption").$$render(
            $$result,
            {
              value: "sticky",
              userSelect: settings.navbarApperance
            },
            {
              userSelect: ($$value) => {
                settings.navbarApperance = $$value;
                $$settled = false;
              }
            },
            {
              label: () => {
                return `${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.theme_customizer.settings.navbar_type.select.sticky.title",
                    source: dictionary
                  },
                  {},
                  {}
                )} `;
              },
              body: () => {
                return `<img${add_attribute("src", stickyNav, 0)} alt="">`;
              }
            }
          )} ${validate_component(CustomizerOption, "CustomizerOption").$$render(
            $$result,
            {
              value: "static",
              userSelect: settings.navbarApperance
            },
            {
              userSelect: ($$value) => {
                settings.navbarApperance = $$value;
                $$settled = false;
              }
            },
            {
              label: () => {
                return `${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.theme_customizer.settings.navbar_type.select.static.title",
                    source: dictionary
                  },
                  {},
                  {}
                )} `;
              },
              body: () => {
                return `<img${add_attribute("src", staticNav, 0)} alt="">`;
              }
            }
          )} ${validate_component(CustomizerOption, "CustomizerOption").$$render(
            $$result,
            {
              value: "hidden",
              userSelect: settings.navbarApperance
            },
            {
              userSelect: ($$value) => {
                settings.navbarApperance = $$value;
                $$settled = false;
              }
            },
            {
              label: () => {
                return `${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.theme_customizer.settings.navbar_type.select.hidden.title",
                    source: dictionary
                  },
                  {},
                  {}
                )} `;
              },
              body: () => {
                return `<img${add_attribute("src", hiddenNav, 0)} alt="">`;
              }
            }
          )}</div> <div class="flex flex-row flex-wrap gap-3 mt-5" role="radiogroup">${validate_component(CustomizerOption, "CustomizerOption").$$render(
            $$result,
            {
              value: "attached",
              userSelect: settings.navbarPosition
            },
            {
              userSelect: ($$value) => {
                settings.navbarPosition = $$value;
                $$settled = false;
              }
            },
            {
              label: () => {
                return `${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.theme_customizer.settings.navbar_type.select.attached.title",
                    source: dictionary
                  },
                  {},
                  {}
                )} `;
              },
              body: () => {
                return `<img${add_attribute("src", stickyNav, 0)} alt="">`;
              }
            }
          )} ${validate_component(CustomizerOption, "CustomizerOption").$$render(
            $$result,
            {
              value: "separated",
              userSelect: settings.navbarPosition
            },
            {
              userSelect: ($$value) => {
                settings.navbarPosition = $$value;
                $$settled = false;
              }
            },
            {
              label: () => {
                return `${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.theme_customizer.settings.navbar_type.select.separated.title",
                    source: dictionary
                  },
                  {},
                  {}
                )} `;
              },
              body: () => {
                return `<img${add_attribute("src", stickyNav, 0)} alt="">`;
              }
            }
          )}</div></div> <div class="flex flex-col py-2"><span class="font-medium text-lg">${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.settings.direction.title",
              source: dictionary
            },
            {},
            {}
          )}</span> <span class="text-neutral-400 text-sm">${validate_component(Text, "Text").$$render(
            $$result,
            {
              key: "dashboard.theme_customizer.settings.direction.description",
              source: dictionary
            },
            {},
            {}
          )}</span> <div class="flex flex-row flex-wrap gap-3 mt-5" role="radiogroup">${validate_component(CustomizerOption, "CustomizerOption").$$render(
            $$result,
            {
              value: "ltr",
              userSelect: settings.direction
            },
            {
              userSelect: ($$value) => {
                settings.direction = $$value;
                $$settled = false;
              }
            },
            {
              label: () => {
                return `${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.theme_customizer.settings.direction.select.ltr.title",
                    source: dictionary
                  },
                  {},
                  {}
                )} `;
              },
              body: () => {
                return `<img${add_attribute("src", ltrImage, 0)} alt="">`;
              }
            }
          )} ${validate_component(CustomizerOption, "CustomizerOption").$$render(
            $$result,
            {
              value: "rtl",
              userSelect: settings.direction
            },
            {
              userSelect: ($$value) => {
                settings.direction = $$value;
                $$settled = false;
              }
            },
            {
              label: () => {
                return `${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.theme_customizer.settings.direction.select.rtl.title",
                    source: dictionary
                  },
                  {},
                  {}
                )} `;
              },
              body: () => {
                return `<img${add_attribute("src", rtlImage, 0)} alt="">`;
              }
            }
          )}</div></div></div>`;
        }
      }
    )} <div class="page svelte-yqfnma"><div class="${["page__navigation svelte-yqfnma", $toggleMenu ? "toggle" : ""].join(" ").trim()}"><div class="page__navigation__header svelte-yqfnma">${$toggleMenu && flags & SHOW_LOGO_WHEN_EXPANDED || !$toggleMenu ? `<div class="logo svelte-yqfnma" data-svelte-h="svelte-6pd7vu"><svg viewBox="0 0 128 128" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" preserveAspectRatio="xMidYMid meet" fill="#000000" class="svelte-yqfnma"><g stroke-width="0"></g><g stroke-linecap="round" stroke-linejoin="round"></g><g><path d="M118.89 75.13a15.693 15.693 0 0 0-7-7.33a22.627 22.627 0 0 0-6-2.63c1.53-5.6-.64-10.06-3.69-13.39c-4.51-4.88-9.2-5.59-9.2-5.59c1.62-3.07 2.11-6.61 1.36-10c-.77-3.69-3.08-6.86-6.36-8.72c-3.1-1.83-6.92-2.73-10.84-3.47c-1.88-.34-9.81-1.45-13.1-6c-2.65-3.69-2.73-10.33-3.45-12.32s-3.38-1.15-6.23.76C51.05 8.7 44.15 15.83 41.49 23a24.6 24.6 0 0 0-1.28 13.89c-2.14.35-4.23.97-6.21 1.85c-.16 0-.32.1-.49.17c-3 1.24-9.43 7-10 15.85c-.21 3.13.19 6.26 1.17 9.24c-2.19.57-4.3 1.43-6.26 2.57c-2.29.98-4.38 2.38-6.15 4.13c-1.95 2.41-3.37 5.2-4.15 8.2a27.594 27.594 0 0 0 2 19.77c1.8 3.47 4.06 6.67 6.74 9.52c8.55 8.79 23.31 12.11 35 14c14.19 2.34 29.05 1.52 42.33-4c19.92-8.22 25.22-21.44 26-25.17c1.73-8.25-.39-16.02-1.3-17.89z" fill="#885742"></path><path d="M87.45 92.89c-1.57.8-3.17 1.52-4.78 2.16c-1.08.43-2.17.82-3.27 1.17c-1.1.36-2.21.67-3.33 1c-2.24.56-4.52.97-6.82 1.21c-1.74.19-3.5.28-5.25.28c-4.62 0-9.22-.65-13.67-1.91l-1.46-.44a55.12 55.12 0 0 1-7.15-2.84l-1.39-.69a22.722 22.722 0 0 0 12.72 15.31c3.43 1.59 7.17 2.4 10.95 2.38c3.82.03 7.6-.75 11.09-2.31a21.868 21.868 0 0 0 12.58-15.44l-.22.12z" fill="#35220b"></path><path d="M85.19 90c-7 1.23-14.09 1.82-21.19 1.77c-7.1.04-14.19-.55-21.19-1.77a2.16 2.16 0 0 0-2.53 2.54v.25A51.578 51.578 0 0 0 64 98.66c1.75 0 3.51-.09 5.25-.28c2.3-.24 4.58-.65 6.82-1.21c1.12-.28 2.23-.59 3.33-1s2.19-.74 3.27-1.17c1.62-.67 3.21-1.39 4.78-2.16l.22-.12l.06-.27c.17-1.19-.66-2.29-1.86-2.46a2.22 2.22 0 0 0-.68.01z" fill="#ffffff"></path><g><circle cx="80.13" cy="69.49" r="12.4" fill="#ffffff"></circle><ellipse cx="80.13" cy="69.49" rx="5.73" ry="5.82" fill="#35220b"></ellipse><circle cx="47.87" cy="69.49" r="12.4" fill="#ffffff"></circle><ellipse cx="47.87" cy="69.49" rx="5.73" ry="5.82" fill="#35220b"></ellipse></g></g></svg></div>` : ``} ${$toggleMenu ? `<p class="text-center mx-auto md:text-lg" data-svelte-h="svelte-nq936n">SM Box</p> ${validate_component(Button, "Button").$$render(
      $$result,
      {
        class: "!p-1 ml-auto rtl:mr-auto rtl:ml-0 w-10 h-10 md:!hidden",
        palette: GHOST,
        flags: NO_RIPPLE | NO_ANIMATION,
        OnClick: () => {
          $toggleMenu = false;
          settingsEntity.SetProperty("expandNavbar", false).Save();
        }
      },
      {},
      {
        default: () => {
          return `<i class="fa-solid fa-xmark"></i>`;
        }
      }
    )}` : ``}</div> <div class="page__navigation__list svelte-yqfnma">${each(links, (link, idx) => {
      return `${validate_component(NavigationLink, "NavigationLinkComponent").$$render($$result, { data: link }, {}, {})}`;
    })}</div></div> <div class="page__container svelte-yqfnma"><div class="${escape(null_to_empty("page__header" + cn(flags & SEPARATED_PAGE_HEADER ? "separated" : "")), true) + " svelte-yqfnma"}"><div class="page__header__content svelte-yqfnma">${validate_component(Button, "Button").$$render(
      $$result,
      {
        class: "!p-1 w-10 h-10",
        palette: GHOST,
        flags: NO_RIPPLE | NO_ANIMATION,
        OnClick: () => {
          $toggleMenu = !$toggleMenu;
          settingsEntity.SetProperty("expandNavbar", $toggleMenu).Save();
        }
      },
      {},
      {
        default: () => {
          return `${$toggleMenu ? `<i class="absolute fa-solid fa-xmark fa-xl"></i>` : `<i class="absolute fa-solid fa-bars fa-xl"></i>`}`;
        }
      }
    )} ${validate_component(Button, "Button").$$render(
      $$result,
      {
        class: "!bg-base-200 ml-5 !hidden md:!flex",
        palette: GHOST,
        flags: NO_RIPPLE | NO_ANIMATION,
        OnClick: () => OpenSearch()
      },
      {},
      {
        default: () => {
          return `<span class="fa-solid fa-magnifying-glass mr-2"></span> <div class="flex flex-row items-center gap-1">${`${validate_component(Keyboard, "Keyboard").$$render($$result, { char: "ctrl" }, {}, {})}`}
                        +
                        ${validate_component(Keyboard, "Keyboard").$$render($$result, { char: "f" }, {}, {})}</div>`;
        }
      }
    )} <div class="ml-auto rtl:mr-auto rtl:ml-0 flex items-center flex-wrap">${validate_component(LanguageSelect, "LanguageSelect").$$render($$result, {}, {}, {})} ${validate_component(Button, "Button").$$render(
      $$result,
      {
        class: "!p-1 w-10 h-10 !text-neutral-400 active:!text-primary md:!hidden",
        flags: NO_RIPPLE | PLAIN | CIRCLE,
        OnClick: () => search.Open()
      },
      {},
      {
        default: () => {
          return `<span class="fa-solid fa-magnifying-glass"></span>`;
        }
      }
    )} ${validate_component(Popover, "Popover").$$render($$result, { class: "!p-0", flags: DISABLE_CLOSE_BTN }, {}, {
      trigger: ({ melt, trigger }) => {
        return `<div${spread([{ slot: "trigger" }, escape_object(trigger)], { classes: "svelte-yqfnma" })}>${validate_component(Button, "Button").$$render(
          $$result,
          {
            class: "!p-1 w-10 h-10 !text-neutral-400 active:!text-primary",
            flags: NO_RIPPLE | PLAIN | CIRCLE
          },
          {},
          {
            default: () => {
              return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" style="fill: currentColor" class="svelte-yqfnma"><path d="M10 3H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zM9 9H5V5h4v4zm11 4h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1zm-1 6h-4v-4h4v4zM17 3c-2.206 0-4 1.794-4 4s1.794 4 4 4 4-1.794 4-4-1.794-4-4-4zm0 6c-1.103 0-2-.897-2-2s.897-2 2-2 2 .897 2 2-.897 2-2 2zM7 13c-2.206 0-4 1.794-4 4s1.794 4 4 4 4-1.794 4-4-1.794-4-4-4zm0 6c-1.103 0-2-.897-2-2s.897-2 2-2 2 .897 2 2-.897 2-2 2z"></path></svg>`;
            }
          }
        )}</div>`;
      },
      default: () => {
        return `<div class="flex flex-col"><div class="flex flex-row items-center justify-between col-span-3 border-b border-base-300 border-dashed py-4 px-5"><span class="text-xl font-semibold">${validate_component(Text, "Text").$$render(
          $$result,
          {
            key: "dashboard.navbar.shortcuts.title",
            source: dictionary
          },
          {},
          {}
        )}</span></div> <div class="grid grid-cols-3 p-2">${each(shortcuts, (shortcut, idx) => {
          return `<a${add_attribute("target", shortcut.target ?? "_self", 0)} rel="noreferrer"${add_attribute("href", shortcut.link, 0)} class="${escape(null_to_empty("shortcut" + cn(shortcut.iconPalette ?? "")), true) + " svelte-yqfnma"}">${shortcut.icon ? `${shortcut.iconType == "custom" ? `<img width="24" height="24"${add_attribute("src", shortcut.icon, 0)} alt=""${add_attribute("style", shortcut.style ?? "", 0)} class="svelte-yqfnma">` : `<span class="${escape(null_to_empty("text-2xl" + cn(shortcut.icon)), true) + " svelte-yqfnma"}"${add_attribute("style", shortcut.style ?? "", 0)}></span>`}` : `<span class="${escape(null_to_empty("fa-regular fa-circle"), true) + " svelte-yqfnma"}"></span>`} <span class="font-semibold truncate svelte-yqfnma">${escape(shortcut.title)}</span> </a>`;
        })}</div></div>`;
      }
    })} ${validate_component(Button, "Button").$$render(
      $$result,
      {
        class: "!p-1 w-10 h-10 !text-neutral-400 active:!text-primary",
        flags: NO_RIPPLE | PLAIN | CIRCLE,
        OnClick: () => fullscreen = !fullscreen
      },
      {},
      {
        default: () => {
          return `${fullscreen ? `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" style="fill: currentColor"><path d="M10 4H8v4H4v2h6zM8 20h2v-6H4v2h4zm12-6h-6v6h2v-4h4zm0-6h-4V4h-2v6h6z"></path></svg>` : `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" style="fill: currentColor"><path d="M5 5h5V3H3v7h2zm5 14H5v-5H3v7h7zm11-5h-2v5h-5v2h7zm-2-4h2V3h-7v2h5z"></path></svg>`}`;
        }
      }
    )} ${validate_component(Popover, "Popover").$$render($$result, { class: "!p-0", flags: DISABLE_CLOSE_BTN }, {}, {
      trigger: ({ melt, trigger }) => {
        return `<div${spread([{ slot: "trigger" }, escape_object(trigger)], { classes: "svelte-yqfnma" })}>${validate_component(Button, "Button").$$render(
          $$result,
          {
            class: "!p-1 w-10 h-10 !text-neutral-400 active:!text-primary",
            flags: NO_RIPPLE | PLAIN | CIRCLE
          },
          {},
          {
            default: () => {
              return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" style="fill: currentColor" class="svelte-yqfnma"><path d="M19 13.586V10c0-3.217-2.185-5.927-5.145-6.742C13.562 2.52 12.846 2 12 2s-1.562.52-1.855 1.258C7.185 4.074 5 6.783 5 10v3.586l-1.707 1.707A.996.996 0 0 0 3 16v2a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-2a.996.996 0 0 0-.293-.707L19 13.586zM19 17H5v-.586l1.707-1.707A.996.996 0 0 0 7 14v-4c0-2.757 2.243-5 5-5s5 2.243 5 5v4c0 .266.105.52.293.707L19 16.414V17zm-7 5a2.98 2.98 0 0 0 2.818-2H9.182A2.98 2.98 0 0 0 12 22z"></path></svg>`;
            }
          }
        )}</div>`;
      },
      default: () => {
        return `${validate_component(NotificationsBlock, "NotificationsBlock").$$render($$result, { dictionary }, {}, {})}`;
      }
    })}</div></div></div> <div class="page__content svelte-yqfnma"><div class="z-[100] absolute top-1/2 -translate-y-1/2 right-0 rtl:left-0 rtl:right-auto">${validate_component(Button, "Button").$$render(
      $$result,
      {
        flags: NO_ANIMATION,
        class: "!w-10 !h-10 !rounded-r-none rtl:!rounded-l-none rtl:!rounded-r-btn hover:*:animate-spin",
        OnClick: () => drawer.Open()
      },
      {},
      {
        default: () => {
          return `<i class="fa-solid fa-gear"></i>`;
        }
      }
    )}</div> ${slots.default ? slots.default({ meta: $routeMeta }) : ``}</div></div> </div>`;
  } while (!$$settled);
  $$unsubscribe_toggleMenu();
  $$unsubscribe_routeMeta();
  $$unsubscribe_navigating();
  return $$rendered;
});
const Layout = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `${validate_component(Navigation, "Navigation").$$render(
    $$result,
    {
      links: [
        {
          title: "Home page",
          href: "/",
          icon: "fa-solid fa-house"
        },
        {
          title: "U.R.L.S",
          href: "/system/urls",
          icon: "fa-solid fa-link"
        }
      ],
      shortcuts: [
        {
          title: "test1",
          iconType: "built-in",
          link: "/"
        },
        {
          title: "test2",
          iconType: "custom",
          icon: "https://img.icons8.com/?size=200&id=44442&format=png&color=000000",
          link: "https://go.dev/",
          target: "_blank"
        }
      ]
    },
    {},
    {
      default: ({ meta }) => {
        return `<h2>${escape(meta.title)}</h2> ${slots.default ? slots.default({}) : ``}`;
      }
    }
  )}`;
});

export { Layout as default };
//# sourceMappingURL=_layout.svelte-CuhXgmuQ.js.map
