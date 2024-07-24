import { c as create_ssr_component, s as subscribe, d as each, v as validate_component, h as spread, i as escape_object, j as escape_attribute_value, b as add_attribute, e as escape, n as noop$1 } from './ssr-0mLmEaQb.js';
import { t as toWritableStores, m as makeElement, u as usePortal, j as createElHelpers, e as executeCallbacks, b as addMeltEventListener, D as generateId, n as noop, I as isTouch, k as kbd, o as omit, a as overridable } from './create-DiqL_npW.js';
import { d as derived, a as readonly, w as writable } from './index2-Cn3cpumX.js';
import { c as cn } from './index-VQC3TRid.js';

const defaults$1 = {
  defaultValue: 0,
  max: 100
};
const { name: name$1 } = createElHelpers("progress");
const createProgress = (props) => {
  const withDefaults = { ...defaults$1, ...props };
  const options = toWritableStores(omit(withDefaults, "value"));
  const { max } = options;
  const valueWritable = withDefaults.value ?? writable(withDefaults.defaultValue);
  const value = overridable(valueWritable, withDefaults?.onValueChange);
  const root = makeElement(name$1(), {
    stores: [value, max],
    returned: ([$value, $max]) => {
      return {
        value: $value,
        max: $max,
        role: "meter",
        "aria-valuemin": 0,
        "aria-valuemax": $max,
        "aria-valuenow": $value,
        "data-value": $value,
        "data-state": $value === null ? "indeterminate" : $value === $max ? "complete" : "loading",
        "data-max": $max
      };
    }
  });
  return {
    elements: {
      root
    },
    states: {
      value
    },
    options
  };
};
const { name } = createElHelpers("toast");
const defaults = {
  closeDelay: 5e3,
  type: "foreground"
};
function createToaster(props) {
  const withDefaults = { ...defaults, ...props };
  const options = toWritableStores(withDefaults);
  const { closeDelay, type } = options;
  const toastsMap = writable(/* @__PURE__ */ new Map());
  const addToast2 = (props2) => {
    const propsWithDefaults = {
      closeDelay: closeDelay.get(),
      type: type.get(),
      ...props2
    };
    const ids = {
      content: generateId(),
      title: generateId(),
      description: generateId()
    };
    const timeout = propsWithDefaults.closeDelay === 0 ? null : window.setTimeout(() => {
      removeToast(ids.content);
    }, propsWithDefaults.closeDelay);
    const getPercentage = () => {
      const { createdAt, pauseDuration, closeDelay: closeDelay2, pausedAt } = toast;
      if (closeDelay2 === 0)
        return 0;
      if (pausedAt) {
        return 100 * (pausedAt - createdAt - pauseDuration) / closeDelay2;
      } else {
        const now = performance.now();
        return 100 * (now - createdAt - pauseDuration) / closeDelay2;
      }
    };
    const toast = {
      id: ids.content,
      ids,
      ...propsWithDefaults,
      timeout,
      createdAt: performance.now(),
      pauseDuration: 0,
      getPercentage
    };
    toastsMap.update((currentMap) => {
      currentMap.set(ids.content, toast);
      return new Map(currentMap);
    });
    return toast;
  };
  const removeToast = (id) => {
    toastsMap.update((currentMap) => {
      currentMap.delete(id);
      return new Map(currentMap);
    });
  };
  const updateToast = (id, data) => {
    toastsMap.update((currentMap) => {
      const toast = currentMap.get(id);
      if (!toast)
        return currentMap;
      currentMap.set(id, { ...toast, data });
      return new Map(currentMap);
    });
  };
  const content = makeElement(name("content"), {
    stores: toastsMap,
    returned: ($toasts) => {
      return (id) => {
        const t = $toasts.get(id);
        if (!t)
          return null;
        const { ...toast } = t;
        return {
          id,
          role: "alert",
          "aria-describedby": toast.ids.description,
          "aria-labelledby": toast.ids.title,
          "aria-live": toast.type === "foreground" ? "assertive" : "polite",
          tabindex: -1
        };
      };
    },
    action: (node) => {
      let destroy = noop;
      destroy = executeCallbacks(addMeltEventListener(node, "pointerenter", (e) => {
        if (isTouch(e))
          return;
        toastsMap.update((currentMap) => {
          const currentToast = currentMap.get(node.id);
          if (!currentToast || currentToast.closeDelay === 0)
            return currentMap;
          if (currentToast.timeout !== null) {
            window.clearTimeout(currentToast.timeout);
          }
          currentToast.pausedAt = performance.now();
          return new Map(currentMap);
        });
      }), addMeltEventListener(node, "pointerleave", (e) => {
        if (isTouch(e))
          return;
        toastsMap.update((currentMap) => {
          const currentToast = currentMap.get(node.id);
          if (!currentToast || currentToast.closeDelay === 0)
            return currentMap;
          const pausedAt = currentToast.pausedAt ?? currentToast.createdAt;
          const elapsed = pausedAt - currentToast.createdAt - currentToast.pauseDuration;
          const remaining = currentToast.closeDelay - elapsed;
          currentToast.timeout = window.setTimeout(() => {
            removeToast(node.id);
          }, remaining);
          currentToast.pauseDuration += performance.now() - pausedAt;
          currentToast.pausedAt = void 0;
          return new Map(currentMap);
        });
      }), () => {
        removeToast(node.id);
      });
      return {
        destroy
      };
    }
  });
  const title = makeElement(name("title"), {
    stores: toastsMap,
    returned: ($toasts) => {
      return (id) => {
        const toast = $toasts.get(id);
        if (!toast)
          return null;
        return {
          id: toast.ids.title
        };
      };
    }
  });
  const description = makeElement(name("description"), {
    stores: toastsMap,
    returned: ($toasts) => {
      return (id) => {
        const toast = $toasts.get(id);
        if (!toast)
          return null;
        return {
          id: toast.ids.description
        };
      };
    }
  });
  const close = makeElement(name("close"), {
    returned: () => {
      return (id) => ({
        type: "button",
        "data-id": id
      });
    },
    action: (node) => {
      function handleClose() {
        if (!node.dataset.id)
          return;
        removeToast(node.dataset.id);
      }
      const unsub = executeCallbacks(addMeltEventListener(node, "click", () => {
        handleClose();
      }), addMeltEventListener(node, "keydown", (e) => {
        if (e.key !== kbd.ENTER && e.key !== kbd.SPACE)
          return;
        e.preventDefault();
        handleClose();
      }));
      return {
        destroy: unsub
      };
    }
  });
  const toasts2 = derived(toastsMap, ($toastsMap) => {
    return Array.from($toastsMap.values());
  });
  return {
    elements: {
      content,
      title,
      description,
      close
    },
    states: {
      toasts: readonly(toasts2)
    },
    helpers: {
      addToast: addToast2,
      removeToast,
      updateToast
    },
    actions: {
      portal: usePortal
    },
    options
  };
}
const css$1 = {
  code: ".toast.svelte-14k9omq.svelte-14k9omq{--toast-main-color:var(--primary);position:relative;overflow:hidden;--tw-bg-opacity:1;background-color:rgb(var(--base1) / var(--tw-bg-opacity));--tw-text-opacity:1;color:rgb(var(--basec) / var(--tw-text-opacity));--tw-shadow:0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);--tw-shadow-colored:0 4px 6px -1px var(--tw-shadow-color), 0 2px 4px -2px var(--tw-shadow-color);box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);border-radius:var(--rounded-box)}.toast.with-main-color.primary.svelte-14k9omq.svelte-14k9omq{--toast-main-color:var(--primary)}.toast.with-main-color.secondary.svelte-14k9omq.svelte-14k9omq{--toast-main-color:var(--secondary)}.toast.with-main-color.neutral.svelte-14k9omq.svelte-14k9omq{--toast-main-color:var(--neutral)}.toast.with-main-color.success.svelte-14k9omq.svelte-14k9omq{--toast-main-color:var(--success)}.toast.with-main-color.info.svelte-14k9omq.svelte-14k9omq{--toast-main-color:var(--info)}.toast.with-main-color.warning.svelte-14k9omq.svelte-14k9omq{--toast-main-color:var(--warning)}.toast.with-main-color.error.svelte-14k9omq.svelte-14k9omq{--toast-main-color:var(--error)}.toast__indicator.svelte-14k9omq.svelte-14k9omq{position:absolute;left:52px;top:0.75rem;height:0.25rem;width:20%;overflow:hidden;border-radius:9999px;background-color:rgb(0 0 0 / 0.2)}.toast__indicator.svelte-14k9omq>div.svelte-14k9omq{height:100%;width:100%;background-color:rgb(var(--toast-main-color))}.toast__content.svelte-14k9omq.svelte-14k9omq{position:relative;display:flex;width:24rem;max-width:calc(100vw - 2rem);align-items:center;justify-content:space-between;gap:1rem;padding:1.25rem;padding-top:1.5rem}.toast__close.svelte-14k9omq.svelte-14k9omq{position:absolute;right:1rem;top:1rem;display:grid;width:1.5rem;height:1.5rem;place-items:center;border-radius:9999px;color:rgb(var(--toast-main-color))}.toast__close.svelte-14k9omq.svelte-14k9omq:hover{background-color:rgb(var(--toast-main-color)/.2)}",
  map: '{"version":3,"file":"Toast.svelte","sources":["Toast.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">export const CHANGE_MAIN_COLOR = 1;\\nexport const RENDER_HTML = 2;\\nexport const INFO = \\"info\\";\\nexport const SUCCESS = \\"success\\";\\nexport const WARNING = \\"warning\\";\\nexport const ERROR = \\"error\\";\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { onMount } from \\"svelte\\";\\nimport { writable } from \\"svelte/store\\";\\nimport { fly } from \\"svelte/transition\\";\\nimport {\\n  createProgress,\\n  melt\\n} from \\"@melt-ui/svelte\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nexport let elements;\\n$: ({ content, title, description, close } = elements);\\nexport let toast;\\n$: ({ data, id, getPercentage } = toast);\\nconst percentage = writable(0);\\nconst {\\n  elements: { root: progress },\\n  options: { max }\\n} = createProgress({\\n  max: 100,\\n  value: percentage\\n});\\nconst GetIcon = function(type) {\\n  switch (type) {\\n    case INFO:\\n      return `<i class=\\"fa-solid fa-circle-info text-info text-xl\\"></i>`;\\n    case SUCCESS:\\n      return `<i class=\\"fa-solid fa-circle-check text-success text-xl\\"></i>`;\\n    case WARNING:\\n      return `<i class=\\"fa-solid fa-triangle-exclamation text-warning text-xl\\"></i>`;\\n    case ERROR:\\n      return `<i class=\\"fa-solid fa-circle-exclamation text-error text-xl\\"></i>`;\\n  }\\n  return \\"\\";\\n};\\nonMount(() => {\\n  let frame;\\n  const updatePercentage = () => {\\n    percentage.set(getPercentage());\\n    frame = requestAnimationFrame(updatePercentage);\\n  };\\n  frame = requestAnimationFrame(updatePercentage);\\n  return () => cancelAnimationFrame(frame);\\n});\\n\\n\\t$: __MELTUI_BUILDER_0__ = $content(id);\\n\\t$: __MELTUI_BUILDER_1__ = $title(id);\\n\\t$: __MELTUI_BUILDER_2__ = $description(id);\\n\\t$: __MELTUI_BUILDER_3__ = $close(id);\\n<\/script>\\r\\n  \\r\\n<div\\r\\n    {...__MELTUI_BUILDER_0__} use:__MELTUI_BUILDER_0__.action\\r\\n    in:fly={{ duration: 150, x: \'100%\' }}\\r\\n    out:fly={{ duration: 150, x: \'100%\' }}\\r\\n    class={\\"toast\\" + cn(data.type) + cn((data.flags & CHANGE_MAIN_COLOR) ? \\"with-main-color\\" : \\"\\")}\\r\\n>\\r\\n    <div\\r\\n        {...$progress} use:$progress.action\\r\\n        class=\\"toast__indicator\\"\\r\\n    >\\r\\n        <div style={`transform: translateX(-${(100 * ($percentage ?? 0)) / ($max ?? 1)}%)`} />\\r\\n    </div>\\r\\n  \\r\\n    <div\\r\\n        class=\\"toast__content\\"\\r\\n    >\\r\\n        <div class=\\"flex flex-row\\">\\r\\n            <div class=\\"flex flex-col items-center justify-center mr-3\\">\\r\\n                {@html GetIcon(data.type)}\\r\\n            </div>\\r\\n            <div class=\\"flex flex-col\\">\\r\\n                <span {...__MELTUI_BUILDER_1__} use:__MELTUI_BUILDER_1__.action class=\\"flex items-center gap-2 text-lg font-semibold\\">\\r\\n                    {data.title}\\r\\n                </span>\\r\\n                <span {...__MELTUI_BUILDER_2__} use:__MELTUI_BUILDER_2__.action class=\\"text-sm\\">\\r\\n                    {#if data.flags & RENDER_HTML}\\r\\n                        {@html data.description}\\r\\n                    {:else}\\r\\n                        {data.description}\\r\\n                    {/if}\\r\\n                </span>\\r\\n            </div>\\r\\n        </div>\\r\\n        <button\\r\\n            {...__MELTUI_BUILDER_3__} use:__MELTUI_BUILDER_3__.action\\r\\n            class=\\"toast__close\\"\\r\\n        >\\r\\n            <i class=\\"fa-solid fa-xmark\\"></i>\\r\\n        </button>\\r\\n    </div>\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .toast {\\r\\n        --toast-main-color: var(--primary);\\r\\n        position: relative;\\r\\n        overflow: hidden;\\r\\n        --tw-bg-opacity: 1;\\r\\n        background-color: rgb(var(--base1) / var(--tw-bg-opacity));\\r\\n        --tw-text-opacity: 1;\\r\\n        color: rgb(var(--basec) / var(--tw-text-opacity));\\r\\n        --tw-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);\\r\\n        --tw-shadow-colored: 0 4px 6px -1px var(--tw-shadow-color), 0 2px 4px -2px var(--tw-shadow-color);\\r\\n        box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);\\r\\n        border-radius: var(--rounded-box);\\r\\n    }\\r\\n\\r\\n        .toast.with-main-color.primary {\\r\\n                --toast-main-color: var(--primary);\\r\\n            }\\r\\n\\r\\n        .toast.with-main-color.secondary {\\r\\n                --toast-main-color: var(--secondary);\\r\\n            }\\r\\n\\r\\n        .toast.with-main-color.neutral {\\r\\n                --toast-main-color: var(--neutral);\\r\\n            }\\r\\n\\r\\n        .toast.with-main-color.success {\\r\\n                --toast-main-color: var(--success);\\r\\n            }\\r\\n\\r\\n        .toast.with-main-color.info {\\r\\n                --toast-main-color: var(--info);\\r\\n            }\\r\\n\\r\\n        .toast.with-main-color.warning {\\r\\n                --toast-main-color: var(--warning);\\r\\n            }\\r\\n\\r\\n        .toast.with-main-color.error {\\r\\n                --toast-main-color: var(--error);\\r\\n            }\\r\\n\\r\\n        .toast__indicator {\\r\\n        position: absolute;\\r\\n        left: 52px;\\r\\n        top: 0.75rem;\\r\\n        height: 0.25rem;\\r\\n        width: 20%;\\r\\n        overflow: hidden;\\r\\n        border-radius: 9999px;\\r\\n        background-color: rgb(0 0 0 / 0.2);\\r\\n}\\r\\n\\r\\n        .toast__indicator > div {\\r\\n        height: 100%;\\r\\n        width: 100%;\\r\\n        background-color: rgb(var(--toast-main-color));\\r\\n}\\r\\n\\r\\n        .toast__content {\\r\\n        position: relative;\\r\\n        display: flex;\\r\\n        width: 24rem;\\r\\n        max-width: calc(100vw - 2rem);\\r\\n        align-items: center;\\r\\n        justify-content: space-between;\\r\\n        gap: 1rem;\\r\\n        padding: 1.25rem;\\r\\n        padding-top: 1.5rem;\\r\\n}\\r\\n\\r\\n        .toast__close {\\r\\n        position: absolute;\\r\\n        right: 1rem;\\r\\n        top: 1rem;\\r\\n        display: grid;\\r\\n        width: 1.5rem;\\r\\n        height: 1.5rem;\\r\\n        place-items: center;\\r\\n        border-radius: 9999px;\\r\\n        color: rgb(var(--toast-main-color));\\r\\n}\\r\\n\\r\\n        .toast__close:hover {\\r\\n        background-color: rgb(var(--toast-main-color)/.2);\\r\\n}\\r\\n</style>"],"names":[],"mappings":"AAoGI,oCAAO,CACH,kBAAkB,CAAE,cAAc,CAClC,QAAQ,CAAE,QAAQ,CAClB,QAAQ,CAAE,MAAM,CAChB,eAAe,CAAE,CAAC,CAClB,gBAAgB,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,eAAe,CAAC,CAAC,CAC1D,iBAAiB,CAAE,CAAC,CACpB,KAAK,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,iBAAiB,CAAC,CAAC,CACjD,WAAW,CAAE,gEAAgE,CAC7E,mBAAmB,CAAE,4EAA4E,CACjG,UAAU,CAAE,IAAI,uBAAuB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,gBAAgB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,WAAW,CAAC,CACvG,aAAa,CAAE,IAAI,aAAa,CACpC,CAEI,MAAM,gBAAgB,sCAAS,CACvB,kBAAkB,CAAE,cACxB,CAEJ,MAAM,gBAAgB,wCAAW,CACzB,kBAAkB,CAAE,gBACxB,CAEJ,MAAM,gBAAgB,sCAAS,CACvB,kBAAkB,CAAE,cACxB,CAEJ,MAAM,gBAAgB,sCAAS,CACvB,kBAAkB,CAAE,cACxB,CAEJ,MAAM,gBAAgB,mCAAM,CACpB,kBAAkB,CAAE,WACxB,CAEJ,MAAM,gBAAgB,sCAAS,CACvB,kBAAkB,CAAE,cACxB,CAEJ,MAAM,gBAAgB,oCAAO,CACrB,kBAAkB,CAAE,YACxB,CAEJ,+CAAkB,CAClB,QAAQ,CAAE,QAAQ,CAClB,IAAI,CAAE,IAAI,CACV,GAAG,CAAE,OAAO,CACZ,MAAM,CAAE,OAAO,CACf,KAAK,CAAE,GAAG,CACV,QAAQ,CAAE,MAAM,CAChB,aAAa,CAAE,MAAM,CACrB,gBAAgB,CAAE,IAAI,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,GAAG,CACzC,CAEQ,gCAAiB,CAAG,kBAAI,CACxB,MAAM,CAAE,IAAI,CACZ,KAAK,CAAE,IAAI,CACX,gBAAgB,CAAE,IAAI,IAAI,kBAAkB,CAAC,CACrD,CAEQ,6CAAgB,CAChB,QAAQ,CAAE,QAAQ,CAClB,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,KAAK,CACZ,SAAS,CAAE,KAAK,KAAK,CAAC,CAAC,CAAC,IAAI,CAAC,CAC7B,WAAW,CAAE,MAAM,CACnB,eAAe,CAAE,aAAa,CAC9B,GAAG,CAAE,IAAI,CACT,OAAO,CAAE,OAAO,CAChB,WAAW,CAAE,MACrB,CAEQ,2CAAc,CACd,QAAQ,CAAE,QAAQ,CAClB,KAAK,CAAE,IAAI,CACX,GAAG,CAAE,IAAI,CACT,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,MAAM,CACb,MAAM,CAAE,MAAM,CACd,WAAW,CAAE,MAAM,CACnB,aAAa,CAAE,MAAM,CACrB,KAAK,CAAE,IAAI,IAAI,kBAAkB,CAAC,CAC1C,CAEQ,2CAAa,MAAO,CACpB,gBAAgB,CAAE,IAAI,IAAI,kBAAkB,CAAC,CAAC,EAAE,CACxD"}'
};
const CHANGE_MAIN_COLOR = 1;
const RENDER_HTML = 2;
const INFO = "info";
const SUCCESS = "success";
const WARNING = "warning";
const ERROR = "error";
const Toast = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let content;
  let title;
  let description;
  let close;
  let data;
  let id;
  let getPercentage;
  let __MELTUI_BUILDER_0__;
  let __MELTUI_BUILDER_1__;
  let __MELTUI_BUILDER_2__;
  let __MELTUI_BUILDER_3__;
  let $close, $$unsubscribe_close = noop$1, $$subscribe_close = () => ($$unsubscribe_close(), $$unsubscribe_close = subscribe(close, ($$value) => $close = $$value), close);
  let $description, $$unsubscribe_description = noop$1, $$subscribe_description = () => ($$unsubscribe_description(), $$unsubscribe_description = subscribe(description, ($$value) => $description = $$value), description);
  let $title, $$unsubscribe_title = noop$1, $$subscribe_title = () => ($$unsubscribe_title(), $$unsubscribe_title = subscribe(title, ($$value) => $title = $$value), title);
  let $content, $$unsubscribe_content = noop$1, $$subscribe_content = () => ($$unsubscribe_content(), $$unsubscribe_content = subscribe(content, ($$value) => $content = $$value), content);
  let $progress, $$unsubscribe_progress;
  let $percentage, $$unsubscribe_percentage;
  let $max, $$unsubscribe_max;
  let { elements: elements2 } = $$props;
  let { toast } = $$props;
  const percentage = writable(0);
  $$unsubscribe_percentage = subscribe(percentage, (value) => $percentage = value);
  const { elements: { root: progress }, options: { max } } = createProgress({ max: 100, value: percentage });
  $$unsubscribe_progress = subscribe(progress, (value) => $progress = value);
  $$unsubscribe_max = subscribe(max, (value) => $max = value);
  const GetIcon = function(type) {
    switch (type) {
      case INFO:
        return `<i class="fa-solid fa-circle-info text-info text-xl"></i>`;
      case SUCCESS:
        return `<i class="fa-solid fa-circle-check text-success text-xl"></i>`;
      case WARNING:
        return `<i class="fa-solid fa-triangle-exclamation text-warning text-xl"></i>`;
      case ERROR:
        return `<i class="fa-solid fa-circle-exclamation text-error text-xl"></i>`;
    }
    return "";
  };
  if ($$props.elements === void 0 && $$bindings.elements && elements2 !== void 0) $$bindings.elements(elements2);
  if ($$props.toast === void 0 && $$bindings.toast && toast !== void 0) $$bindings.toast(toast);
  $$result.css.add(css$1);
  $$subscribe_content({ content, title, description, close } = elements2, $$subscribe_title(), $$subscribe_description(), $$subscribe_close());
  ({ data, id, getPercentage } = toast);
  __MELTUI_BUILDER_0__ = $content(id);
  __MELTUI_BUILDER_1__ = $title(id);
  __MELTUI_BUILDER_2__ = $description(id);
  __MELTUI_BUILDER_3__ = $close(id);
  $$unsubscribe_close();
  $$unsubscribe_description();
  $$unsubscribe_title();
  $$unsubscribe_content();
  $$unsubscribe_progress();
  $$unsubscribe_percentage();
  $$unsubscribe_max();
  return `<div${spread(
    [
      escape_object(__MELTUI_BUILDER_0__),
      {
        class: escape_attribute_value("toast" + cn(data.type) + cn(data.flags & CHANGE_MAIN_COLOR ? "with-main-color" : ""))
      }
    ],
    { classes: "svelte-14k9omq" }
  )}><div${spread([escape_object($progress), { class: "toast__indicator" }], { classes: "svelte-14k9omq" })}><div${add_attribute("style", `transform: translateX(-${100 * ($percentage ?? 0) / ($max ?? 1)}%)`, 0)} class="svelte-14k9omq"></div></div> <div class="toast__content svelte-14k9omq"><div class="flex flex-row"><div class="flex flex-col items-center justify-center mr-3"><!-- HTML_TAG_START -->${GetIcon(data.type)}<!-- HTML_TAG_END --></div> <div class="flex flex-col"><span${spread(
    [
      escape_object(__MELTUI_BUILDER_1__),
      {
        class: "flex items-center gap-2 text-lg font-semibold"
      }
    ],
    { classes: "svelte-14k9omq" }
  )}>${escape(data.title)}</span> <span${spread([escape_object(__MELTUI_BUILDER_2__), { class: "text-sm" }], { classes: "svelte-14k9omq" })}>${data.flags & RENDER_HTML ? `<!-- HTML_TAG_START -->${data.description}<!-- HTML_TAG_END -->` : `${escape(data.description)}`}</span></div></div> <button${spread([escape_object(__MELTUI_BUILDER_3__), { class: "toast__close" }], { classes: "svelte-14k9omq" })}><i class="fa-solid fa-xmark"></i></button></div> </div>`;
});
const css = {
  code: ".toast-container.svelte-zln6l1{position:fixed;right:0px;top:0px;z-index:var(--toast-z-index);margin:1rem;display:flex;flex-direction:column;align-items:flex-end;gap:0.5rem\n}@media(min-width: 768px){.toast-container.svelte-zln6l1{bottom:0px;top:auto\n    }}",
  map: '{"version":3,"file":"Root.svelte","sources":["Root.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">const {\\n  elements,\\n  helpers: { addToast },\\n  states: { toasts },\\n  actions: { portal }\\n} = createToaster();\\nexport const showToast = addToast;\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { flip } from \\"svelte/animate\\";\\nimport { createToaster } from \\"@melt-ui/svelte\\";\\nimport Toast, {} from \\"./Toast.svelte\\";\\n<\/script>\\r\\n\\r\\n  \\r\\n<div\\r\\n    class=\\"toast-container\\"\\r\\n    use:portal\\r\\n>\\r\\n    {#each $toasts as toast (toast.id)}\\r\\n        <div animate:flip={{ duration: 500 }}>\\r\\n            <Toast {elements} {toast} />\\r\\n        </div>\\r\\n    {/each}\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .toast-container {\\n    position: fixed;\\n    right: 0px;\\n    top: 0px;\\n    z-index: var(--toast-z-index);\\n    margin: 1rem;\\n    display: flex;\\n    flex-direction: column;\\n    align-items: flex-end;\\n    gap: 0.5rem\\n}\\n@media (min-width: 768px) {\\n    .toast-container {\\n        bottom: 0px;\\n        top: auto\\n    }\\n}\\r\\n</style>"],"names":[],"mappings":"AA2BI,8BAAiB,CACjB,QAAQ,CAAE,KAAK,CACf,KAAK,CAAE,GAAG,CACV,GAAG,CAAE,GAAG,CACR,OAAO,CAAE,IAAI,eAAe,CAAC,CAC7B,MAAM,CAAE,IAAI,CACZ,OAAO,CAAE,IAAI,CACb,cAAc,CAAE,MAAM,CACtB,WAAW,CAAE,QAAQ,CACrB,GAAG,CAAE,MAAM;AACf,CACA,MAAO,YAAY,KAAK,CAAE,CACtB,8BAAiB,CACb,MAAM,CAAE,GAAG,CACX,GAAG,CAAE,IAAI;AACjB,IAAI,CACJ"}'
};
const { elements, helpers: { addToast }, states: { toasts }, actions: { portal } } = createToaster();
const showToast = addToast;
const Root = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $toasts, $$unsubscribe_toasts;
  $$unsubscribe_toasts = subscribe(toasts, (value) => $toasts = value);
  $$result.css.add(css);
  $$unsubscribe_toasts();
  return `<div class="toast-container svelte-zln6l1">${each($toasts, (toast) => {
    return `<div>${validate_component(Toast, "Toast").$$render($$result, { elements, toast }, {}, {})} </div>`;
  })} </div>`;
});

export { CHANGE_MAIN_COLOR as C, ERROR as E, Root as R, showToast as s };
//# sourceMappingURL=Root-CbJvCgYd.js.map
