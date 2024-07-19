import { c as create_ssr_component, g as getContext, p as onDestroy, e as escape, b as add_attribute, v as validate_component, n as noop, a as subscribe, f as null_to_empty, h as spread, i as escape_object, o as set_store_value, j as escape_attribute_value, d as each, l as get_store_value } from './ssr-BISMo5iU.js';
import { w as writable, r as readable, d as derived, a as readonly } from './index-DyJcXWsV.js';
import { c as cn, m as makeElement, o as omit, e as executeCallbacks, y as addEventListener, b as addMeltEventListener, z as isHTMLInputElement, A as isContentEditable, w as withGet, a as overridable, t as toWritableStores, x as createElHelpers, g as generateIds, q as dequal, B as isObject, C as stripValues, h as disabledAttr, k as kbd, D as isHTMLButtonElement, d as tick, j as isHTMLElement, F as FIRST_LAST_KEYS, E as isElementDisabled, s as styleToString, f as isElement, G as generateId, H as createHiddenInput, i as isBrowser, I as getElementByMeltId, J as isHTMLLabelElement, n as noop$1 } from './index2-C4fvnV3D.js';
import { s as showToast, E as ERROR, C as CHANGE_MAIN_COLOR } from './Root-D8IkUrtT.js';
import { G as GetText, T as Text } from './Text-BcVg2DSs.js';
import { B as Button, S as SUBMIT, i as SMOKE, e as effect, d as derivedVisible, l as last, j as back, k as forward, p as prev, n as next, f as usePopper, g as getPortalDestination, r as removeScroll, m as debounce, w as wrapArray, t as toggle, s as sleep } from './Button-BQNu2r8c.js';

function addHighlight(element) {
  element.setAttribute("data-highlighted", "");
}
function removeHighlight(element) {
  element.removeAttribute("data-highlighted");
}
function getOptions(el) {
  return Array.from(el.querySelectorAll('[role="option"]:not([data-disabled])')).filter((el2) => isHTMLElement(el2));
}
function handleRovingFocus(nextElement) {
  if (!isBrowser)
    return;
  sleep(1).then(() => {
    const currentFocusedElement = document.activeElement;
    if (!isHTMLElement(currentFocusedElement) || currentFocusedElement === nextElement)
      return;
    currentFocusedElement.tabIndex = -1;
    if (nextElement) {
      nextElement.tabIndex = 0;
      nextElement.focus();
    }
  });
}
const ignoredKeys = /* @__PURE__ */ new Set(["Shift", "Control", "Alt", "Meta", "CapsLock", "NumLock"]);
const defaults$1 = {
  onMatch: handleRovingFocus,
  getCurrentItem: () => document.activeElement
};
function createTypeaheadSearch(args = {}) {
  const withDefaults = { ...defaults$1, ...args };
  const typed = withGet(writable([]));
  const resetTyped = debounce(() => {
    typed.update(() => []);
  });
  const handleTypeaheadSearch = (key, items) => {
    if (ignoredKeys.has(key))
      return;
    const currentItem = withDefaults.getCurrentItem();
    const $typed = get_store_value(typed);
    if (!Array.isArray($typed)) {
      return;
    }
    $typed.push(key.toLowerCase());
    typed.set($typed);
    const candidateItems = items.filter((item) => {
      if (item.getAttribute("disabled") === "true" || item.getAttribute("aria-disabled") === "true" || item.hasAttribute("data-disabled")) {
        return false;
      }
      return true;
    });
    const isRepeated = $typed.length > 1 && $typed.every((char) => char === $typed[0]);
    const normalizeSearch = isRepeated ? $typed[0] : $typed.join("");
    const currentItemIndex = isHTMLElement(currentItem) ? candidateItems.indexOf(currentItem) : -1;
    let wrappedItems = wrapArray(candidateItems, Math.max(currentItemIndex, 0));
    const excludeCurrentItem = normalizeSearch.length === 1;
    if (excludeCurrentItem) {
      wrappedItems = wrappedItems.filter((v) => v !== currentItem);
    }
    const nextItem = wrappedItems.find((item) => item?.innerText && item.innerText.toLowerCase().startsWith(normalizeSearch.toLowerCase()));
    if (isHTMLElement(nextItem) && nextItem !== currentItem) {
      withDefaults.onMatch(nextItem);
    }
    resetTyped();
  };
  return {
    typed,
    resetTyped,
    handleTypeaheadSearch
  };
}
function createClickOutsideIgnore(meltId) {
  return (e) => {
    const target = e.target;
    const triggerEl = getElementByMeltId(meltId);
    if (!triggerEl || !isElement(target))
      return false;
    const id = triggerEl.id;
    if (isHTMLLabelElement(target) && id === target.htmlFor) {
      return true;
    }
    if (target.closest(`label[for="${id}"]`)) {
      return true;
    }
    return false;
  };
}
function createLabel() {
  const root = makeElement("label", {
    action: (node) => {
      const mouseDown = addMeltEventListener(node, "mousedown", (e) => {
        if (!e.defaultPrevented && e.detail > 1) {
          e.preventDefault();
        }
      });
      return {
        destroy: mouseDown
      };
    }
  });
  return {
    elements: {
      root
    }
  };
}
const INTERACTION_KEYS = [kbd.ARROW_LEFT, kbd.ESCAPE, kbd.ARROW_RIGHT, kbd.SHIFT, kbd.CAPS_LOCK, kbd.CONTROL, kbd.ALT, kbd.META, kbd.ENTER, kbd.F1, kbd.F2, kbd.F3, kbd.F4, kbd.F5, kbd.F6, kbd.F7, kbd.F8, kbd.F9, kbd.F10, kbd.F11, kbd.F12];
const defaults = {
  positioning: {
    placement: "bottom",
    sameWidth: true
  },
  scrollAlignment: "nearest",
  loop: true,
  defaultOpen: false,
  closeOnOutsideClick: true,
  preventScroll: true,
  escapeBehavior: "close",
  forceVisible: false,
  portal: "body",
  builder: "listbox",
  disabled: false,
  required: false,
  name: void 0,
  typeahead: true,
  highlightOnHover: true,
  onOutsideClick: void 0,
  preventTextSelectionOverflow: true
};
const listboxIdParts = ["trigger", "menu", "label"];
function createListbox(props) {
  const withDefaults = { ...defaults, ...props };
  const activeTrigger = withGet(writable(null));
  const highlightedItem = withGet(writable(null));
  const selectedWritable = withDefaults.selected ?? writable(withDefaults.defaultSelected);
  const selected = overridable(selectedWritable, withDefaults?.onSelectedChange);
  const highlighted = derived(highlightedItem, ($highlightedItem) => $highlightedItem ? getOptionProps($highlightedItem) : void 0);
  const openWritable = withDefaults.open ?? writable(withDefaults.defaultOpen);
  const open = overridable(openWritable, withDefaults?.onOpenChange);
  const options = toWritableStores({
    ...omit(withDefaults, "open", "defaultOpen", "builder", "ids"),
    multiple: withDefaults.multiple ?? false
  });
  const { scrollAlignment, loop, closeOnOutsideClick, escapeBehavior, preventScroll, portal, forceVisible, positioning, multiple, arrowSize, disabled, required, typeahead, name: nameProp, highlightOnHover, onOutsideClick, preventTextSelectionOverflow } = options;
  const { name: name2, selector } = createElHelpers(withDefaults.builder);
  const ids = toWritableStores({ ...generateIds(listboxIdParts), ...withDefaults.ids });
  const { handleTypeaheadSearch } = createTypeaheadSearch({
    onMatch: (element) => {
      highlightedItem.set(element);
      element.scrollIntoView({ block: scrollAlignment.get() });
    },
    getCurrentItem() {
      return highlightedItem.get();
    }
  });
  function getOptionProps(el) {
    const value = el.getAttribute("data-value");
    const label2 = el.getAttribute("data-label");
    const disabled2 = el.hasAttribute("data-disabled");
    return {
      value: value ? JSON.parse(value) : value,
      label: label2 ?? el.textContent ?? void 0,
      disabled: disabled2 ? true : false
    };
  }
  const setOption = (newOption) => {
    selected.update(($option) => {
      const $multiple = multiple.get();
      if ($multiple) {
        const optionArr = Array.isArray($option) ? [...$option] : [];
        return toggle(newOption, optionArr, (itemA, itemB) => dequal(itemA.value, itemB.value));
      }
      return newOption;
    });
  };
  function selectItem(item) {
    const props2 = getOptionProps(item);
    setOption(props2);
  }
  async function openMenu() {
    open.set(true);
    await tick();
    const menuElement = document.getElementById(ids.menu.get());
    if (!isHTMLElement(menuElement))
      return;
    const selectedItem = menuElement.querySelector("[aria-selected=true]");
    if (!isHTMLElement(selectedItem))
      return;
    highlightedItem.set(selectedItem);
  }
  function closeMenu() {
    open.set(false);
    highlightedItem.set(null);
  }
  const isVisible = derivedVisible({ open, forceVisible, activeTrigger });
  const isSelected = derived([selected], ([$selected]) => {
    return (value) => {
      if (Array.isArray($selected)) {
        return $selected.some((o) => dequal(o.value, value));
      }
      if (isObject(value)) {
        return dequal($selected?.value, stripValues(value, void 0));
      }
      return dequal($selected?.value, value);
    };
  });
  const isHighlighted = derived([highlighted], ([$value]) => {
    return (item) => {
      return dequal($value?.value, item);
    };
  });
  const trigger = makeElement(name2("trigger"), {
    stores: [open, highlightedItem, disabled, ids.menu, ids.trigger, ids.label],
    returned: ([$open, $highlightedItem, $disabled, $menuId, $triggerId, $labelId]) => {
      return {
        "aria-activedescendant": $highlightedItem?.id,
        "aria-autocomplete": "list",
        "aria-controls": $menuId,
        "aria-expanded": $open,
        "aria-labelledby": $labelId,
        "data-state": $open ? "open" : "closed",
        // autocomplete: 'off',
        id: $triggerId,
        role: "combobox",
        disabled: disabledAttr($disabled),
        type: withDefaults.builder === "select" ? "button" : void 0
      };
    },
    action: (node) => {
      activeTrigger.set(node);
      const isInput = isHTMLInputElement(node);
      const unsubscribe = executeCallbacks(
        addMeltEventListener(node, "click", () => {
          node.focus();
          const $open = open.get();
          if ($open) {
            closeMenu();
          } else {
            openMenu();
          }
        }),
        // Handle all input key events including typing, meta, and navigation.
        addMeltEventListener(node, "keydown", (e) => {
          const $open = open.get();
          if (!$open) {
            if (INTERACTION_KEYS.includes(e.key)) {
              return;
            }
            if (e.key === kbd.TAB) {
              return;
            }
            if (e.key === kbd.BACKSPACE && isInput && node.value === "") {
              return;
            }
            if (e.key === kbd.SPACE && isHTMLButtonElement(node)) {
              return;
            }
            openMenu();
            tick().then(() => {
              const $selectedItem = selected.get();
              if ($selectedItem)
                return;
              const menuEl = document.getElementById(ids.menu.get());
              if (!isHTMLElement(menuEl))
                return;
              const enabledItems = Array.from(menuEl.querySelectorAll(`${selector("item")}:not([data-disabled]):not([data-hidden])`)).filter((item) => isHTMLElement(item));
              if (!enabledItems.length)
                return;
              if (e.key === kbd.ARROW_DOWN) {
                highlightedItem.set(enabledItems[0]);
                enabledItems[0].scrollIntoView({ block: scrollAlignment.get() });
              } else if (e.key === kbd.ARROW_UP) {
                highlightedItem.set(last(enabledItems));
                last(enabledItems).scrollIntoView({ block: scrollAlignment.get() });
              }
            });
          }
          if (e.key === kbd.TAB) {
            closeMenu();
            return;
          }
          if (e.key === kbd.ENTER && !e.isComposing || e.key === kbd.SPACE && isHTMLButtonElement(node)) {
            e.preventDefault();
            const $highlightedItem = highlightedItem.get();
            if ($highlightedItem) {
              selectItem($highlightedItem);
            }
            if (!multiple.get()) {
              closeMenu();
            }
          }
          if (e.key === kbd.ARROW_UP && e.altKey) {
            closeMenu();
          }
          if (FIRST_LAST_KEYS.includes(e.key)) {
            e.preventDefault();
            const menuElement = document.getElementById(ids.menu.get());
            if (!isHTMLElement(menuElement))
              return;
            const itemElements = getOptions(menuElement);
            if (!itemElements.length)
              return;
            const candidateNodes = itemElements.filter((opt) => !isElementDisabled(opt) && opt.dataset.hidden === void 0);
            const $currentItem = highlightedItem.get();
            const currentIndex = $currentItem ? candidateNodes.indexOf($currentItem) : -1;
            const $loop = loop.get();
            const $scrollAlignment = scrollAlignment.get();
            let nextItem;
            switch (e.key) {
              case kbd.ARROW_DOWN:
                nextItem = next(candidateNodes, currentIndex, $loop);
                break;
              case kbd.ARROW_UP:
                nextItem = prev(candidateNodes, currentIndex, $loop);
                break;
              case kbd.PAGE_DOWN:
                nextItem = forward(candidateNodes, currentIndex, 10, $loop);
                break;
              case kbd.PAGE_UP:
                nextItem = back(candidateNodes, currentIndex, 10, $loop);
                break;
              case kbd.HOME:
                nextItem = candidateNodes[0];
                break;
              case kbd.END:
                nextItem = last(candidateNodes);
                break;
              default:
                return;
            }
            highlightedItem.set(nextItem);
            nextItem?.scrollIntoView({ block: $scrollAlignment });
          } else if (typeahead.get()) {
            const menuEl = document.getElementById(ids.menu.get());
            if (!isHTMLElement(menuEl))
              return;
            handleTypeaheadSearch(e.key, getOptions(menuEl));
          }
        })
      );
      return {
        destroy() {
          activeTrigger.set(null);
          unsubscribe();
        }
      };
    }
  });
  const menu = makeElement(name2("menu"), {
    stores: [isVisible, ids.menu],
    returned: ([$isVisible, $menuId]) => {
      return {
        hidden: $isVisible ? void 0 : true,
        id: $menuId,
        role: "listbox",
        style: $isVisible ? void 0 : styleToString({ display: "none" })
      };
    },
    action: (node) => {
      let unsubPopper = noop$1;
      const unsubscribe = executeCallbacks(
        // Bind the popper portal to the input element.
        effect([isVisible, portal, closeOnOutsideClick, positioning, activeTrigger], ([$isVisible, $portal, $closeOnOutsideClick, $positioning, $activeTrigger]) => {
          unsubPopper();
          if (!$isVisible || !$activeTrigger)
            return;
          tick().then(() => {
            unsubPopper();
            const ignoreHandler = createClickOutsideIgnore(ids.trigger.get());
            unsubPopper = usePopper(node, {
              anchorElement: $activeTrigger,
              open,
              options: {
                floating: $positioning,
                focusTrap: null,
                modal: {
                  closeOnInteractOutside: $closeOnOutsideClick,
                  onClose: closeMenu,
                  shouldCloseOnInteractOutside: (e) => {
                    onOutsideClick.get()?.(e);
                    if (e.defaultPrevented)
                      return false;
                    const target = e.target;
                    if (!isElement(target))
                      return false;
                    if (target === $activeTrigger || $activeTrigger.contains(target)) {
                      return false;
                    }
                    if (ignoreHandler(e))
                      return false;
                    return true;
                  }
                },
                escapeKeydown: { handler: closeMenu, behaviorType: escapeBehavior },
                portal: getPortalDestination(node, $portal),
                preventTextSelectionOverflow: { enabled: preventTextSelectionOverflow }
              }
            }).destroy;
          });
        })
      );
      return {
        destroy: () => {
          unsubscribe();
          unsubPopper();
        }
      };
    }
  });
  const { elements: { root: labelBuilder } } = createLabel();
  const { action: labelAction } = get_store_value(labelBuilder);
  const label = makeElement(name2("label"), {
    stores: [ids.label, ids.trigger],
    returned: ([$labelId, $triggerId]) => {
      return {
        id: $labelId,
        for: $triggerId
      };
    },
    action: labelAction
  });
  const option = makeElement(name2("option"), {
    stores: [isSelected],
    returned: ([$isSelected]) => (props2) => {
      const selected2 = $isSelected(props2.value);
      return {
        "data-value": JSON.stringify(props2.value),
        "data-label": props2.label,
        "data-disabled": disabledAttr(props2.disabled),
        "aria-disabled": props2.disabled ? true : void 0,
        "aria-selected": selected2,
        "data-selected": selected2 ? "" : void 0,
        id: generateId(),
        role: "option"
      };
    },
    action: (node) => {
      const unsubscribe = executeCallbacks(addMeltEventListener(node, "click", (e) => {
        if (isElementDisabled(node)) {
          e.preventDefault();
          return;
        }
        selectItem(node);
        if (!multiple.get()) {
          closeMenu();
        }
      }), effect(highlightOnHover, ($highlightOnHover) => {
        if (!$highlightOnHover)
          return;
        const unsub = executeCallbacks(addMeltEventListener(node, "mouseover", () => {
          highlightedItem.set(node);
        }), addMeltEventListener(node, "mouseleave", () => {
          highlightedItem.set(null);
        }));
        return unsub;
      }));
      return { destroy: unsubscribe };
    }
  });
  const group = makeElement(name2("group"), {
    returned: () => {
      return (groupId) => ({
        role: "group",
        "aria-labelledby": groupId
      });
    }
  });
  const groupLabel = makeElement(name2("group-label"), {
    returned: () => {
      return (groupId) => ({
        id: groupId
      });
    }
  });
  const hiddenInput = createHiddenInput({
    value: derived([selected], ([$selected]) => {
      const value = Array.isArray($selected) ? $selected.map((o) => o.value) : $selected?.value;
      return typeof value === "string" ? value : JSON.stringify(value);
    }),
    name: readonly(nameProp),
    required,
    prefix: withDefaults.builder
  });
  const arrow = makeElement(name2("arrow"), {
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
  effect([highlightedItem], ([$highlightedItem]) => {
    if (!isBrowser)
      return;
    const menuElement = document.getElementById(ids.menu.get());
    if (!isHTMLElement(menuElement))
      return;
    getOptions(menuElement).forEach((node) => {
      if (node === $highlightedItem) {
        addHighlight(node);
      } else {
        removeHighlight(node);
      }
    });
  });
  effect([open, preventScroll], ([$open, $preventScroll]) => {
    if (!isBrowser || !$open || !$preventScroll)
      return;
    return removeScroll();
  });
  return {
    ids,
    elements: {
      trigger,
      group,
      option,
      menu,
      groupLabel,
      label,
      hiddenInput,
      arrow
    },
    states: {
      open,
      selected,
      highlighted,
      highlightedItem
    },
    helpers: {
      isSelected,
      isHighlighted,
      closeMenu
    },
    options
  };
}
const { name } = createElHelpers("combobox");
function createCombobox(props) {
  const listbox = createListbox({ ...props, builder: "combobox", typeahead: false });
  const inputValue = writable("");
  const touchedInput = writable(false);
  const input = makeElement(name("input"), {
    stores: [listbox.elements.trigger, inputValue],
    returned: ([$trigger, $inputValue]) => {
      return {
        ...omit($trigger, "action"),
        role: "combobox",
        value: $inputValue,
        autocomplete: "off"
      };
    },
    action: (node) => {
      const unsubscribe = executeCallbacks(
        addMeltEventListener(node, "input", (e) => {
          if (!isHTMLInputElement(e.target) && !isContentEditable(e.target))
            return;
          touchedInput.set(true);
        }),
        // This shouldn't be cancelled ever, so we don't use addMeltEventListener.
        addEventListener(node, "input", (e) => {
          if (isHTMLInputElement(e.target)) {
            inputValue.set(e.target.value);
          }
          if (isContentEditable(e.target)) {
            inputValue.set(e.target.innerText);
          }
        })
      );
      const { destroy } = listbox.elements.trigger(node);
      return {
        destroy() {
          destroy?.();
          unsubscribe();
        }
      };
    }
  });
  effect(listbox.states.open, ($open) => {
    if (!$open) {
      touchedInput.set(false);
    }
  });
  return {
    ...listbox,
    elements: {
      ...omit(listbox.elements, "trigger"),
      input
    },
    states: {
      ...listbox.states,
      touchedInput,
      inputValue
    }
  };
}
const defaultVal = function(type) {
  switch (type) {
    case "string":
      return "";
    case "number":
      return 0;
    case "boolean":
      return false;
    case "date":
      return /* @__PURE__ */ new Date(0);
  }
};
class FormField {
  name;
  title;
  type;
  required;
  constructor(data) {
    this.name = data.name;
    this.title = data.title ?? data.name;
    this.type = data.type ?? "string";
    this.required = data.required ?? false;
  }
  validate(data) {
    if (this.required) {
      switch (this.type) {
        case "string":
          if (typeof data.val != "string") return {
            result: false,
            msg: data.requiredValueText || "Invalid value"
          };
          if (data.val == "") return {
            result: false,
            msg: data.requiredValueText || "Value is required"
          };
          break;
        case "number":
          if (typeof data.val != "number") return {
            result: false,
            msg: data.requiredValueText || "Invalid value"
          };
          if (data.val == 0) return {
            result: false,
            msg: data.requiredValueText || "Value is required"
          };
          break;
        case "boolean":
          if (typeof data.val != "boolean") return {
            result: false,
            msg: data.requiredValueText || "Invalid value"
          };
          if (data.val == false) return {
            result: false,
            msg: data.requiredValueText || "Value is required"
          };
          break;
        case "date":
          if (!(data.val instanceof Date) || data.val.toString() == "Invalid Date") return {
            result: false,
            msg: data.requiredValueText || "Invalid value"
          };
          break;
      }
    }
    return { result: true };
  }
  parse(val) {
    switch (this.type) {
      case "string":
        return `${val}`;
      case "number":
        const num = Number(val);
        return Number.isNaN(num) ? 0 : num;
      case "boolean":
        switch (val) {
          case "t":
          case "T":
          case "true":
          case "True":
          case "1":
          case 1:
            return true;
          default:
            return false;
        }
      case "date":
        switch (typeof val) {
          case "string":
          case "number":
            const date = new Date(val);
            if (date.toString() == "Invalid Date") return defaultVal(this.type);
            return date;
          default:
            return defaultVal(this.type);
        }
      default:
        return defaultVal(this.type);
    }
  }
}
const Form = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $field_errors, $$unsubscribe_field_errors;
  let $values, $$unsubscribe_values;
  let { class: className = "" } = $$props;
  let { values = writable({}) } = $$props;
  $$unsubscribe_values = subscribe(values, (value) => $values = value);
  const InitValues = () => fields.forEach((field) => set_store_value(values, $values[field.name] = defaultVal(field.type), $values));
  let field_errors = writable({});
  $$unsubscribe_field_errors = subscribe(field_errors, (value) => $field_errors = value);
  const errors = readable({}, (set) => {
    const unsubscribe = field_errors.subscribe((val) => set(val));
    return () => unsubscribe();
  });
  const SetFieldError = function(field, msg) {
    if (!fields.find((f) => f.name == field)) return console.warn(`[Form] Failed to update error of selected field: field "${field}" is not defined.`);
    set_store_value(field_errors, $field_errors[field] = msg, $field_errors);
  };
  const submit = function() {
    if (!validate()) return;
    OnSubmit({
      values: Object.fromEntries(fields.map((field) => [field.name, field.parse($values[field.name])])),
      error
    });
  };
  const validate = function() {
    let validationFailed = false;
    for (var i = 0; i < fields.length; i++) {
      const field = fields[i];
      const { result, msg } = field.validate({
        val: $values[field.name],
        invalidValueText,
        requiredValueText
      });
      if (!result) {
        set_store_value(field_errors, $field_errors[field.name] = msg ?? "", $field_errors);
        validationFailed = true;
      } else set_store_value(field_errors, $field_errors[field.name] = "", $field_errors);
    }
    return !validationFailed;
  };
  const clear = function() {
    fields.forEach((field) => {
      set_store_value(field_errors, $field_errors[field.name] = "", $field_errors);
      set_store_value(values, $values[field.name] = defaultVal(field.type), $values);
    });
  };
  let { fields = [] } = $$props;
  let { flags = 0 } = $$props;
  let { error = (data) => {
    if (data.field) {
      set_store_value(field_errors, $field_errors[data.field] = data.msg, $field_errors);
      return;
    }
    alert(data.msg);
  } } = $$props;
  let { OnSubmit = () => {
  } } = $$props;
  let { invalidValueText = "Invalid value" } = $$props;
  let { requiredValueText = "Value is required" } = $$props;
  if ($$props.class === void 0 && $$bindings.class && className !== void 0) $$bindings.class(className);
  if ($$props.values === void 0 && $$bindings.values && values !== void 0) $$bindings.values(values);
  if ($$props.errors === void 0 && $$bindings.errors && errors !== void 0) $$bindings.errors(errors);
  if ($$props.SetFieldError === void 0 && $$bindings.SetFieldError && SetFieldError !== void 0) $$bindings.SetFieldError(SetFieldError);
  if ($$props.submit === void 0 && $$bindings.submit && submit !== void 0) $$bindings.submit(submit);
  if ($$props.validate === void 0 && $$bindings.validate && validate !== void 0) $$bindings.validate(validate);
  if ($$props.clear === void 0 && $$bindings.clear && clear !== void 0) $$bindings.clear(clear);
  if ($$props.fields === void 0 && $$bindings.fields && fields !== void 0) $$bindings.fields(fields);
  if ($$props.flags === void 0 && $$bindings.flags && flags !== void 0) $$bindings.flags(flags);
  if ($$props.error === void 0 && $$bindings.error && error !== void 0) $$bindings.error(error);
  if ($$props.OnSubmit === void 0 && $$bindings.OnSubmit && OnSubmit !== void 0) $$bindings.OnSubmit(OnSubmit);
  if ($$props.invalidValueText === void 0 && $$bindings.invalidValueText && invalidValueText !== void 0) $$bindings.invalidValueText(invalidValueText);
  if ($$props.requiredValueText === void 0 && $$bindings.requiredValueText && requiredValueText !== void 0) $$bindings.requiredValueText(requiredValueText);
  {
    InitValues();
  }
  $$unsubscribe_field_errors();
  $$unsubscribe_values();
  return `   <form${add_attribute("class", cn(className), 0)}>${slots.default ? slots.default({}) : ``}</form>`;
});
const css$2 = {
  code: ".ds-card.svelte-1byia0x.svelte-1byia0x{position:relative;border-width:1px;--tw-border-opacity:1;border-color:rgb(var(--base3) / var(--tw-border-opacity));--tw-bg-opacity:1;background-color:rgb(var(--base1) / var(--tw-bg-opacity));--tw-shadow:0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);--tw-shadow-colored:0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);border-radius:var(--rounded-box)\n}.ds-card.no-border.svelte-1byia0x.svelte-1byia0x{border-style:none\n}.ds-card.no-shadow.svelte-1byia0x.svelte-1byia0x{--tw-shadow:0 0 #0000;--tw-shadow-colored:0 0 #0000;box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)\n}.ds-card.svelte-1byia0x>.header.svelte-1byia0x{display:flex;flex-direction:column;gap:0.5rem;padding-left:1rem;padding-right:1rem;padding-top:1.5rem\n}.ds-card.svelte-1byia0x>.header.svelte-1byia0x:empty{padding:0px;display:none\n}.ds-card.svelte-1byia0x>.body.svelte-1byia0x{display:flex;flex-direction:column;overflow-y:auto;padding:1rem\n}.ds-card.svelte-1byia0x>.body.svelte-1byia0x:empty{padding:0px;display:none\n}.ds-card.svelte-1byia0x>.footer.svelte-1byia0x{display:flex;flex-direction:row;justify-content:flex-end;gap:0.5rem;padding-left:1rem;padding-right:1rem;padding-top:1.5rem;padding-bottom:0.75rem\n}.ds-card.svelte-1byia0x>.footer.svelte-1byia0x:empty{padding:0px;display:none\n}.ds-card.svelte-1byia0x>.fallback.svelte-1byia0x{position:relative;display:flex;height:100%;width:100%;flex-direction:column\n}.ds-card.svelte-1byia0x>.fallback.svelte-1byia0x:empty{display:none\n}",
  map: '{"version":3,"file":"Card.svelte","sources":["Card.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">export const NO_BORDER = 1;\\nexport const NO_SHADOW = 2;\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { cn } from \\"@/lib/helpers\\";\\nlet className = \\"\\";\\nexport { className as class };\\nexport let flags = 0;\\nexport let style = \\"\\";\\n<\/script>\\r\\n\\r\\n<div class={\\"ds-card\\" + cn(className) + cn(flags & NO_BORDER ? \\"no-border\\" : \\"\\") + cn(flags & NO_SHADOW ? \\"no-shadow\\" : \\"\\")} style={style}>\\r\\n    <div class=\\"header\\">\\r\\n        <slot name=\\"header\\" />\\r\\n    </div>\\r\\n    <div class=\\"body\\">\\r\\n        <slot />\\r\\n    </div>\\r\\n    <div class=\\"footer\\">\\r\\n        <slot name=\\"footer\\" />\\r\\n    </div>\\r\\n    <div class=\\"fallback\\">\\r\\n        <slot name=\\"fallback\\" />\\r\\n    </div>\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .ds-card {\\r\\n\\r\\n    position: relative;\\r\\n\\r\\n    border-width: 1px;\\r\\n\\r\\n    --tw-border-opacity: 1;\\r\\n\\r\\n    border-color: rgb(var(--base3) / var(--tw-border-opacity));\\r\\n\\r\\n    --tw-bg-opacity: 1;\\r\\n\\r\\n    background-color: rgb(var(--base1) / var(--tw-bg-opacity));\\r\\n\\r\\n    --tw-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);\\r\\n\\r\\n    --tw-shadow-colored: 0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);\\r\\n\\r\\n    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);\\r\\n\\r\\n    border-radius: var(--rounded-box)\\n}\\r\\n\\r\\n        .ds-card.no-border {\\r\\n\\r\\n    border-style: none\\n}\\r\\n\\r\\n        .ds-card.no-shadow {\\r\\n\\r\\n    --tw-shadow: 0 0 #0000;\\r\\n\\r\\n    --tw-shadow-colored: 0 0 #0000;\\r\\n\\r\\n    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)\\n}\\r\\n\\r\\n        .ds-card > .header {\\r\\n\\r\\n    display: flex;\\r\\n\\r\\n    flex-direction: column;\\r\\n\\r\\n    gap: 0.5rem;\\r\\n\\r\\n    padding-left: 1rem;\\r\\n\\r\\n    padding-right: 1rem;\\r\\n\\r\\n    padding-top: 1.5rem\\n}\\r\\n\\r\\n        .ds-card > .header:empty {\\r\\n\\r\\n    padding: 0px;\\r\\n\\r\\n    display: none\\n}\\r\\n\\r\\n        .ds-card > .body {\\r\\n\\r\\n    display: flex;\\r\\n\\r\\n    flex-direction: column;\\r\\n\\r\\n    overflow-y: auto;\\r\\n\\r\\n    padding: 1rem\\n}\\r\\n\\r\\n        .ds-card > .body:empty {\\r\\n\\r\\n    padding: 0px;\\r\\n\\r\\n    display: none\\n}\\r\\n\\r\\n        .ds-card > .footer {\\r\\n\\r\\n    display: flex;\\r\\n\\r\\n    flex-direction: row;\\r\\n\\r\\n    justify-content: flex-end;\\r\\n\\r\\n    gap: 0.5rem;\\r\\n\\r\\n    padding-left: 1rem;\\r\\n\\r\\n    padding-right: 1rem;\\r\\n\\r\\n    padding-top: 1.5rem;\\r\\n\\r\\n    padding-bottom: 0.75rem\\n}\\r\\n\\r\\n        .ds-card > .footer:empty {\\r\\n\\r\\n    padding: 0px;\\r\\n\\r\\n    display: none\\n}\\r\\n\\r\\n        .ds-card > .fallback {\\r\\n\\r\\n    position: relative;\\r\\n\\r\\n    display: flex;\\r\\n\\r\\n    height: 100%;\\r\\n\\r\\n    width: 100%;\\r\\n\\r\\n    flex-direction: column\\n}\\r\\n\\r\\n        .ds-card > .fallback:empty {\\r\\n\\r\\n    display: none\\n}\\r\\n</style>"],"names":[],"mappings":"AA2BI,sCAAS,CAET,QAAQ,CAAE,QAAQ,CAElB,YAAY,CAAE,GAAG,CAEjB,mBAAmB,CAAE,CAAC,CAEtB,YAAY,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,mBAAmB,CAAC,CAAC,CAE1D,eAAe,CAAE,CAAC,CAElB,gBAAgB,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,eAAe,CAAC,CAAC,CAE1D,WAAW,CAAE,kEAAkE,CAE/E,mBAAmB,CAAE,8EAA8E,CAEnG,UAAU,CAAE,IAAI,uBAAuB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,gBAAgB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,WAAW,CAAC,CAEvG,aAAa,CAAE,IAAI,aAAa,CAAC;AACrC,CAEQ,QAAQ,wCAAW,CAEvB,YAAY,CAAE,IAAI;AACtB,CAEQ,QAAQ,wCAAW,CAEvB,WAAW,CAAE,SAAS,CAEtB,mBAAmB,CAAE,SAAS,CAE9B,UAAU,CAAE,IAAI,uBAAuB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,gBAAgB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,WAAW,CAAC;AAC3G,CAEQ,uBAAQ,CAAG,sBAAQ,CAEvB,OAAO,CAAE,IAAI,CAEb,cAAc,CAAE,MAAM,CAEtB,GAAG,CAAE,MAAM,CAEX,YAAY,CAAE,IAAI,CAElB,aAAa,CAAE,IAAI,CAEnB,WAAW,CAAE,MAAM;AACvB,CAEQ,uBAAQ,CAAG,sBAAO,MAAO,CAE7B,OAAO,CAAE,GAAG,CAEZ,OAAO,CAAE,IAAI;AACjB,CAEQ,uBAAQ,CAAG,oBAAM,CAErB,OAAO,CAAE,IAAI,CAEb,cAAc,CAAE,MAAM,CAEtB,UAAU,CAAE,IAAI,CAEhB,OAAO,CAAE,IAAI;AACjB,CAEQ,uBAAQ,CAAG,oBAAK,MAAO,CAE3B,OAAO,CAAE,GAAG,CAEZ,OAAO,CAAE,IAAI;AACjB,CAEQ,uBAAQ,CAAG,sBAAQ,CAEvB,OAAO,CAAE,IAAI,CAEb,cAAc,CAAE,GAAG,CAEnB,eAAe,CAAE,QAAQ,CAEzB,GAAG,CAAE,MAAM,CAEX,YAAY,CAAE,IAAI,CAElB,aAAa,CAAE,IAAI,CAEnB,WAAW,CAAE,MAAM,CAEnB,cAAc,CAAE,OAAO;AAC3B,CAEQ,uBAAQ,CAAG,sBAAO,MAAO,CAE7B,OAAO,CAAE,GAAG,CAEZ,OAAO,CAAE,IAAI;AACjB,CAEQ,uBAAQ,CAAG,wBAAU,CAEzB,QAAQ,CAAE,QAAQ,CAElB,OAAO,CAAE,IAAI,CAEb,MAAM,CAAE,IAAI,CAEZ,KAAK,CAAE,IAAI,CAEX,cAAc,CAAE,MAAM;AAC1B,CAEQ,uBAAQ,CAAG,wBAAS,MAAO,CAE/B,OAAO,CAAE,IAAI;AACjB"}'
};
const NO_BORDER = 1;
const NO_SHADOW = 2;
const Card = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { class: className = "" } = $$props;
  let { flags = 0 } = $$props;
  let { style = "" } = $$props;
  if ($$props.class === void 0 && $$bindings.class && className !== void 0) $$bindings.class(className);
  if ($$props.flags === void 0 && $$bindings.flags && flags !== void 0) $$bindings.flags(flags);
  if ($$props.style === void 0 && $$bindings.style && style !== void 0) $$bindings.style(style);
  $$result.css.add(css$2);
  return `<div class="${escape(null_to_empty("ds-card" + cn(className) + cn(flags & NO_BORDER ? "no-border" : "") + cn(flags & NO_SHADOW ? "no-shadow" : "")), true) + " svelte-1byia0x"}"${add_attribute("style", style, 0)}><div class="header svelte-1byia0x">${slots.header ? slots.header({}) : ``}</div> <div class="body svelte-1byia0x">${slots.default ? slots.default({}) : ``}</div> <div class="footer svelte-1byia0x">${slots.footer ? slots.footer({}) : ``}</div> <div class="fallback svelte-1byia0x">${slots.fallback ? slots.fallback({}) : ``}</div> </div>`;
});
const css$1 = {
  code: '.input-container.svelte-svly5k.svelte-svly5k{--input-main-color:var(--primary);--input-background-color:var(--base1);--input-border-color:var(--base3);--input-shadow-color:var(--primary);display:flex;width:100%;flex-direction:column;gap:0.5rem}.input-container.transparent.svelte-svly5k>.input-container__inner.svelte-svly5k{background-color:transparent}.input-container.underlined.svelte-svly5k>.input-container__inner.svelte-svly5k{border-radius:0px;border-top-width:0px;border-left-width:0px;border-right-width:0px}.input-container.underlined.svelte-svly5k>.input-container__inner.svelte-svly5k:focus-within{box-shadow:none}.input-container.required.svelte-svly5k>label::after{margin-left:0.125rem;--tw-text-opacity:1;color:rgb(var(--error) / var(--tw-text-opacity));--tw-content:"*";content:var(--tw-content)}.input-container__inner.svelte-svly5k.svelte-svly5k{margin-top:0.5rem;display:flex;width:100%;flex-direction:row;align-items:center;overflow:hidden;border-width:1px;border-color:rgb(var(--input-border-color));background-color:rgb(var(--input-background-color));border-radius:var(--rounded-input);transition-property:box-shadow;transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1);transition-duration:150ms;transition-duration:var(--animation-input)}.input-container__inner.svelte-svly5k.svelte-svly5k:focus-within{--input-focus-shadow:0 0 0 0.25rem rgb(var(--input-shadow-color) / .3);box-shadow:var(--input-focus-shadow)}.input-container__inner.svelte-svly5k>input.svelte-svly5k{--input-autofill-shadow:0 0 0 1000px rgb(var(--input-main-color)) inset;width:100%;-webkit-appearance:none;-moz-appearance:none;appearance:none;border-width:0px;background-color:transparent;padding:0.75rem;outline:2px solid transparent !important;outline-offset:2px !important}.input-container__inner.svelte-svly5k>input.svelte-svly5k:-webkit-autofill{-webkit-box-shadow:var(--input-autofill-shadow) !important}.input-container__inner.svelte-svly5k>.show-password.svelte-svly5k{padding:0.75rem;color:rgb(var(--input-main-color));outline:2px solid transparent !important;outline-offset:2px !important;transition-property:transform;transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1);transition-duration:150ms;transition-duration:var(--animation-input)}.input-container__inner.svelte-svly5k>.show-password.svelte-svly5k:focus{--tw-scale-x:1.1;--tw-scale-y:1.1;transform:translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))}.input-container.primary.svelte-svly5k.svelte-svly5k{--input-main-color:var(--primary);--input-border-color:var(--primary);--input-shadow-color:var(--primary)}.input-container.secondary.svelte-svly5k.svelte-svly5k{--input-main-color:var(--secondary);--input-border-color:var(--secondary);--input-shadow-color:var(--secondary)}.input-container.neutral.svelte-svly5k.svelte-svly5k{--input-main-color:var(--neutral);--input-border-color:var(--neutral);--input-shadow-color:var(--neutral)}.input-container.success.svelte-svly5k.svelte-svly5k{--input-main-color:var(--success);--input-border-color:var(--success);--input-shadow-color:var(--success)}.input-container.info.svelte-svly5k.svelte-svly5k{--input-main-color:var(--info);--input-border-color:var(--info);--input-shadow-color:var(--info)}.input-container.warning.svelte-svly5k.svelte-svly5k{--input-main-color:var(--warning);--input-border-color:var(--warning);--input-shadow-color:var(--warning)}.input-container.error.svelte-svly5k.svelte-svly5k{--input-main-color:var(--error);--input-border-color:var(--error);--input-shadow-color:var(--error)}.input-container.smoke.svelte-svly5k.svelte-svly5k{--input-main-color:var(--smoke);--input-border-color:var(--smoke);--input-shadow-color:var(--smoke)}',
  map: '{"version":3,"file":"TextInput.svelte","sources":["TextInput.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">export const REQUIRED = 1;\\nexport const SECURE = 2;\\nexport const SHOW_PASSWORD_BTN = 4;\\nexport const TRANSPARENT = 8;\\nexport const UNDERLINE = 16;\\nexport const DEFAULT = \\"\\";\\nexport const PRIMARY = \\"primary\\";\\nexport const SECONDARY = \\"secondary\\";\\nexport const NEUTRAL = \\"neutral\\";\\nexport const SUCCESS = \\"success\\";\\nexport const INFO = \\"info\\";\\nexport const WARNING = \\"warning\\";\\nexport const ERROR = \\"error\\";\\nexport const SMOKE = \\"smoke\\";\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { fade, scale } from \\"svelte/transition\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nlet showPassword = false;\\nexport const SetValue = function(val) {\\n  value = val ? String(val) : \\"\\";\\n};\\nexport const GetValue = function() {\\n  return String(value);\\n};\\nexport const Value = function(val) {\\n  if (val) SetValue(val);\\n  return GetValue();\\n};\\nlet className = \\"\\";\\nexport { className as class };\\nexport let style = \\"\\";\\nexport let name;\\nexport let label = \\"\\";\\nexport let min = null;\\nexport let max = null;\\nexport let error = null;\\nexport let palette = DEFAULT;\\nexport let flags = 0 | SHOW_PASSWORD_BTN;\\nexport let value = \\"\\";\\n$: label = label.trim();\\n$: min = min == null ? null : min < 0 ? 0 : min;\\n$: max = max == null ? null : max < 0 ? 0 : max;\\n<\/script>\\r\\n\\r\\n<div \\r\\n    class={\\"input-container\\" + cn(className) + cn(palette)} \\r\\n    class:required={flags & REQUIRED} \\r\\n    class:transparent={flags & TRANSPARENT} \\r\\n    class:underlined={flags & UNDERLINE}\\r\\n    style={style}\\r\\n>\\r\\n    <slot name=\\"label\\">\\r\\n        {#if label && label != \\"\\"}\\r\\n            <label for={name}>{label}</label>\\r\\n        {/if}\\r\\n    </slot>\\r\\n\\r\\n    <div class=\\"input-container__inner\\">\\r\\n        <input {...{\\r\\n            name: name,\\r\\n            type: (flags & SECURE) && !showPassword ? \\"password\\" : \\"text\\",\\r\\n            minlength: min,\\r\\n            maxlength: max,\\r\\n        }} bind:value={value}>\\r\\n\\r\\n        {#if (flags & SHOW_PASSWORD_BTN) && (flags & SECURE) && value.length > 0}\\r\\n            <button type=\\"button\\" transition:scale class=\\"show-password\\" on:click={() => showPassword = !showPassword}>\\r\\n                {#if showPassword}\\r\\n                    <i class=\\"fa-solid fa-eye-slash\\"></i>\\r\\n                {:else}\\r\\n                    <i class=\\"fa-solid fa-eye\\"></i>\\r\\n                {/if}\\r\\n            </button>\\r\\n        {/if}\\r\\n    </div>\\r\\n\\r\\n    {#if error}\\r\\n        <span class=\\"text-error\\" transition:fade>{ error }</span>\\r\\n    {/if}\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .input-container {\\r\\n        --input-main-color: var(--primary);\\r\\n        --input-background-color: var(--base1);\\r\\n        --input-border-color: var(--base3);\\r\\n        --input-shadow-color: var(--primary);\\r\\n        display: flex;\\r\\n        width: 100%;\\r\\n        flex-direction: column;\\r\\n        gap: 0.5rem;\\r\\n    }\\r\\n\\r\\n        .input-container.transparent >  .input-container__inner {\\r\\n        background-color: transparent;\\r\\n}\\r\\n\\r\\n        .input-container.underlined > .input-container__inner {\\r\\n        border-radius: 0px;\\r\\n        border-top-width: 0px;\\r\\n        border-left-width: 0px;\\r\\n        border-right-width: 0px;\\r\\n}\\r\\n\\r\\n        .input-container.underlined > .input-container__inner:focus-within {\\r\\n                box-shadow: none;\\r\\n            }\\r\\n\\r\\n        .input-container.required > :global(label)::after {\\r\\n        margin-left: 0.125rem;\\r\\n        --tw-text-opacity: 1;\\r\\n        color: rgb(var(--error) / var(--tw-text-opacity));\\r\\n        --tw-content: \\"*\\";\\r\\n        content: var(--tw-content);\\r\\n}\\r\\n\\r\\n        .input-container__inner {\\r\\n        margin-top: 0.5rem;\\r\\n        display: flex;\\r\\n        width: 100%;\\r\\n        flex-direction: row;\\r\\n        align-items: center;\\r\\n        overflow: hidden;\\r\\n        border-width: 1px;\\r\\n        border-color: rgb(var(--input-border-color));\\r\\n        background-color: rgb(var(--input-background-color));\\r\\n        border-radius: var(--rounded-input);\\r\\n        transition-property: box-shadow;\\r\\n        transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\\r\\n        transition-duration: 150ms;\\r\\n        transition-duration: var(--animation-input);\\r\\n}\\r\\n\\r\\n        .input-container__inner:focus-within {\\r\\n                --input-focus-shadow: 0 0 0 0.25rem rgb(var(--input-shadow-color) / .3);\\r\\n\\r\\n                box-shadow: var(--input-focus-shadow);\\r\\n            }\\r\\n\\r\\n        .input-container__inner > input {\\r\\n                --input-autofill-shadow: 0 0 0 1000px rgb(var(--input-main-color)) inset;\\r\\n                width: 100%;\\r\\n                -webkit-appearance: none;\\r\\n                   -moz-appearance: none;\\r\\n                        appearance: none;\\r\\n                border-width: 0px;\\r\\n                background-color: transparent;\\r\\n                padding: 0.75rem;\\r\\n                outline: 2px solid transparent !important;\\r\\n                outline-offset: 2px !important;\\r\\n            }\\r\\n\\r\\n        .input-container__inner > input:-webkit-autofill {\\r\\n                    -webkit-box-shadow: var(--input-autofill-shadow) !important;\\r\\n                }\\r\\n\\r\\n        .input-container__inner > .show-password {\\r\\n        padding: 0.75rem;\\r\\n        color: rgb(var(--input-main-color));\\r\\n        outline: 2px solid transparent !important;\\r\\n        outline-offset: 2px !important;\\r\\n        transition-property: transform;\\r\\n        transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\\r\\n        transition-duration: 150ms;\\r\\n        transition-duration: var(--animation-input);\\r\\n}\\r\\n\\r\\n        .input-container__inner > .show-password:focus {\\r\\n        --tw-scale-x: 1.1;\\r\\n        --tw-scale-y: 1.1;\\r\\n        transform: translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y));\\r\\n}\\r\\n\\r\\n        .input-container.primary {\\r\\n            --input-main-color: var(--primary);\\r\\n            --input-border-color: var(--primary);\\r\\n            --input-shadow-color: var(--primary);\\r\\n        }\\r\\n\\r\\n        .input-container.secondary {\\r\\n            --input-main-color: var(--secondary);\\r\\n            --input-border-color: var(--secondary);\\r\\n            --input-shadow-color: var(--secondary);\\r\\n        }\\r\\n\\r\\n        .input-container.neutral {\\r\\n            --input-main-color: var(--neutral);\\r\\n            --input-border-color: var(--neutral);\\r\\n            --input-shadow-color: var(--neutral);\\r\\n        }\\r\\n\\r\\n        .input-container.success {\\r\\n            --input-main-color: var(--success);\\r\\n            --input-border-color: var(--success);\\r\\n            --input-shadow-color: var(--success);\\r\\n        }\\r\\n\\r\\n        .input-container.info {\\r\\n            --input-main-color: var(--info);\\r\\n            --input-border-color: var(--info);\\r\\n            --input-shadow-color: var(--info);\\r\\n        }\\r\\n\\r\\n        .input-container.warning {\\r\\n            --input-main-color: var(--warning);\\r\\n            --input-border-color: var(--warning);\\r\\n            --input-shadow-color: var(--warning);\\r\\n        }\\r\\n\\r\\n        .input-container.error {\\r\\n            --input-main-color: var(--error);\\r\\n            --input-border-color: var(--error);\\r\\n            --input-shadow-color: var(--error);\\r\\n        }\\r\\n\\r\\n        .input-container.smoke {\\r\\n            --input-main-color: var(--smoke);\\r\\n            --input-border-color: var(--smoke);\\r\\n            --input-shadow-color: var(--smoke);\\r\\n        }\\r\\n</style>"],"names":[],"mappings":"AAmFI,4CAAiB,CACb,kBAAkB,CAAE,cAAc,CAClC,wBAAwB,CAAE,YAAY,CACtC,oBAAoB,CAAE,YAAY,CAClC,oBAAoB,CAAE,cAAc,CACpC,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,IAAI,CACX,cAAc,CAAE,MAAM,CACtB,GAAG,CAAE,MACT,CAEI,gBAAgB,0BAAY,CAAI,qCAAwB,CACxD,gBAAgB,CAAE,WAC1B,CAEQ,gBAAgB,yBAAW,CAAG,qCAAwB,CACtD,aAAa,CAAE,GAAG,CAClB,gBAAgB,CAAE,GAAG,CACrB,iBAAiB,CAAE,GAAG,CACtB,kBAAkB,CAAE,GAC5B,CAEQ,gBAAgB,yBAAW,CAAG,qCAAuB,aAAc,CAC3D,UAAU,CAAE,IAChB,CAEJ,gBAAgB,uBAAS,CAAW,KAAM,OAAQ,CAClD,WAAW,CAAE,QAAQ,CACrB,iBAAiB,CAAE,CAAC,CACpB,KAAK,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,iBAAiB,CAAC,CAAC,CACjD,YAAY,CAAE,GAAG,CACjB,OAAO,CAAE,IAAI,YAAY,CACjC,CAEQ,mDAAwB,CACxB,UAAU,CAAE,MAAM,CAClB,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,IAAI,CACX,cAAc,CAAE,GAAG,CACnB,WAAW,CAAE,MAAM,CACnB,QAAQ,CAAE,MAAM,CAChB,YAAY,CAAE,GAAG,CACjB,YAAY,CAAE,IAAI,IAAI,oBAAoB,CAAC,CAAC,CAC5C,gBAAgB,CAAE,IAAI,IAAI,wBAAwB,CAAC,CAAC,CACpD,aAAa,CAAE,IAAI,eAAe,CAAC,CACnC,mBAAmB,CAAE,UAAU,CAC/B,0BAA0B,CAAE,aAAa,GAAG,CAAC,CAAC,CAAC,CAAC,CAAC,GAAG,CAAC,CAAC,CAAC,CAAC,CACxD,mBAAmB,CAAE,KAAK,CAC1B,mBAAmB,CAAE,IAAI,iBAAiB,CAClD,CAEQ,mDAAuB,aAAc,CAC7B,oBAAoB,CAAE,iDAAiD,CAEvE,UAAU,CAAE,IAAI,oBAAoB,CACxC,CAEJ,qCAAuB,CAAG,mBAAM,CACxB,uBAAuB,CAAE,+CAA+C,CACxE,KAAK,CAAE,IAAI,CACX,kBAAkB,CAAE,IAAI,CACrB,eAAe,CAAE,IAAI,CAChB,UAAU,CAAE,IAAI,CACxB,YAAY,CAAE,GAAG,CACjB,gBAAgB,CAAE,WAAW,CAC7B,OAAO,CAAE,OAAO,CAChB,OAAO,CAAE,GAAG,CAAC,KAAK,CAAC,WAAW,CAAC,UAAU,CACzC,cAAc,CAAE,GAAG,CAAC,UACxB,CAEJ,qCAAuB,CAAG,mBAAK,iBAAkB,CACrC,kBAAkB,CAAE,IAAI,uBAAuB,CAAC,CAAC,UACrD,CAER,qCAAuB,CAAG,4BAAe,CACzC,OAAO,CAAE,OAAO,CAChB,KAAK,CAAE,IAAI,IAAI,kBAAkB,CAAC,CAAC,CACnC,OAAO,CAAE,GAAG,CAAC,KAAK,CAAC,WAAW,CAAC,UAAU,CACzC,cAAc,CAAE,GAAG,CAAC,UAAU,CAC9B,mBAAmB,CAAE,SAAS,CAC9B,0BAA0B,CAAE,aAAa,GAAG,CAAC,CAAC,CAAC,CAAC,CAAC,GAAG,CAAC,CAAC,CAAC,CAAC,CACxD,mBAAmB,CAAE,KAAK,CAC1B,mBAAmB,CAAE,IAAI,iBAAiB,CAClD,CAEQ,qCAAuB,CAAG,4BAAc,MAAO,CAC/C,YAAY,CAAE,GAAG,CACjB,YAAY,CAAE,GAAG,CACjB,SAAS,CAAE,UAAU,IAAI,gBAAgB,CAAC,CAAC,CAAC,IAAI,gBAAgB,CAAC,CAAC,CAAC,OAAO,IAAI,WAAW,CAAC,CAAC,CAAC,MAAM,IAAI,WAAW,CAAC,CAAC,CAAC,MAAM,IAAI,WAAW,CAAC,CAAC,CAAC,OAAO,IAAI,YAAY,CAAC,CAAC,CAAC,OAAO,IAAI,YAAY,CAAC,CACtM,CAEQ,gBAAgB,oCAAS,CACrB,kBAAkB,CAAE,cAAc,CAClC,oBAAoB,CAAE,cAAc,CACpC,oBAAoB,CAAE,cAC1B,CAEA,gBAAgB,sCAAW,CACvB,kBAAkB,CAAE,gBAAgB,CACpC,oBAAoB,CAAE,gBAAgB,CACtC,oBAAoB,CAAE,gBAC1B,CAEA,gBAAgB,oCAAS,CACrB,kBAAkB,CAAE,cAAc,CAClC,oBAAoB,CAAE,cAAc,CACpC,oBAAoB,CAAE,cAC1B,CAEA,gBAAgB,oCAAS,CACrB,kBAAkB,CAAE,cAAc,CAClC,oBAAoB,CAAE,cAAc,CACpC,oBAAoB,CAAE,cAC1B,CAEA,gBAAgB,iCAAM,CAClB,kBAAkB,CAAE,WAAW,CAC/B,oBAAoB,CAAE,WAAW,CACjC,oBAAoB,CAAE,WAC1B,CAEA,gBAAgB,oCAAS,CACrB,kBAAkB,CAAE,cAAc,CAClC,oBAAoB,CAAE,cAAc,CACpC,oBAAoB,CAAE,cAC1B,CAEA,gBAAgB,kCAAO,CACnB,kBAAkB,CAAE,YAAY,CAChC,oBAAoB,CAAE,YAAY,CAClC,oBAAoB,CAAE,YAC1B,CAEA,gBAAgB,kCAAO,CACnB,kBAAkB,CAAE,YAAY,CAChC,oBAAoB,CAAE,YAAY,CAClC,oBAAoB,CAAE,YAC1B"}'
};
const REQUIRED$1 = 1;
const SECURE = 2;
const SHOW_PASSWORD_BTN = 4;
const TRANSPARENT$1 = 8;
const UNDERLINE$1 = 16;
const DEFAULT$1 = "";
const TextInput = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let showPassword = false;
  const SetValue = function(val) {
    value = val ? String(val) : "";
  };
  const GetValue = function() {
    return String(value);
  };
  const Value = function(val) {
    if (val) SetValue(val);
    return GetValue();
  };
  let { class: className = "" } = $$props;
  let { style = "" } = $$props;
  let { name: name2 } = $$props;
  let { label = "" } = $$props;
  let { min = null } = $$props;
  let { max = null } = $$props;
  let { error = null } = $$props;
  let { palette = DEFAULT$1 } = $$props;
  let { flags = 0 | SHOW_PASSWORD_BTN } = $$props;
  let { value = "" } = $$props;
  if ($$props.SetValue === void 0 && $$bindings.SetValue && SetValue !== void 0) $$bindings.SetValue(SetValue);
  if ($$props.GetValue === void 0 && $$bindings.GetValue && GetValue !== void 0) $$bindings.GetValue(GetValue);
  if ($$props.Value === void 0 && $$bindings.Value && Value !== void 0) $$bindings.Value(Value);
  if ($$props.class === void 0 && $$bindings.class && className !== void 0) $$bindings.class(className);
  if ($$props.style === void 0 && $$bindings.style && style !== void 0) $$bindings.style(style);
  if ($$props.name === void 0 && $$bindings.name && name2 !== void 0) $$bindings.name(name2);
  if ($$props.label === void 0 && $$bindings.label && label !== void 0) $$bindings.label(label);
  if ($$props.min === void 0 && $$bindings.min && min !== void 0) $$bindings.min(min);
  if ($$props.max === void 0 && $$bindings.max && max !== void 0) $$bindings.max(max);
  if ($$props.error === void 0 && $$bindings.error && error !== void 0) $$bindings.error(error);
  if ($$props.palette === void 0 && $$bindings.palette && palette !== void 0) $$bindings.palette(palette);
  if ($$props.flags === void 0 && $$bindings.flags && flags !== void 0) $$bindings.flags(flags);
  if ($$props.value === void 0 && $$bindings.value && value !== void 0) $$bindings.value(value);
  $$result.css.add(css$1);
  label = label.trim();
  min = min == null ? null : min < 0 ? 0 : min;
  max = max == null ? null : max < 0 ? 0 : max;
  return `<div class="${[
    escape(null_to_empty("input-container" + cn(className) + cn(palette)), true) + " svelte-svly5k",
    (flags & REQUIRED$1 ? "required" : "") + " " + (flags & TRANSPARENT$1 ? "transparent" : "") + " " + (flags & UNDERLINE$1 ? "underlined" : "")
  ].join(" ").trim()}"${add_attribute("style", style, 0)}>${slots.label ? slots.label({}) : ` ${label && label != "" ? `<label${add_attribute("for", name2, 0)}>${escape(label)}</label>` : ``} `} <div class="input-container__inner svelte-svly5k"><input${spread(
    [
      escape_object({
        name: name2,
        type: flags & SECURE && !showPassword ? "password" : "text",
        minlength: min,
        maxlength: max
      })
    ],
    { classes: "svelte-svly5k" }
  )}${add_attribute("value", value, 0)}> ${flags & SHOW_PASSWORD_BTN && flags & SECURE && value.length > 0 ? `<button type="button" class="show-password svelte-svly5k">${`<i class="fa-solid fa-eye"></i>`}</button>` : ``}</div> ${error ? `<span class="text-error">${escape(error)}</span>` : ``} </div>`;
});
const css = {
  code: '.select-container.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--primary);--select-background-color:var(--base1);--select-border-color:var(--base3);--select-shadow-color:var(--primary);display:flex;width:100%;flex-direction:column;gap:0.5rem}.select-container.transparent.svelte-6e6oqn>.select-container__inner.svelte-6e6oqn.svelte-6e6oqn{background-color:transparent}.select-container.underlined.svelte-6e6oqn>.select-container__inner.svelte-6e6oqn.svelte-6e6oqn{border-radius:0px;border-top-width:0px;border-left-width:0px;border-right-width:0px}.select-container.underlined.svelte-6e6oqn>.select-container__inner.svelte-6e6oqn.svelte-6e6oqn:focus-within{box-shadow:none}.select-container.required.svelte-6e6oqn>label::after{margin-left:0.125rem;--tw-text-opacity:1;color:rgb(var(--error) / var(--tw-text-opacity));--tw-content:"*";content:var(--tw-content)}.select-container__inner.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{margin-top:0.5rem;display:flex;width:100%;flex-direction:row;align-items:center;overflow:hidden;border-width:1px;border-color:rgb(var(--select-border-color));background-color:rgb(var(--select-background-color));border-radius:var(--rounded-input);transition-property:box-shadow;transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1);transition-duration:150ms;transition-duration:var(--animation-input)}.select-container__inner.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn:focus-within{--select-focus-shadow:0 0 0 0.25rem rgba(var(--select-shadow-color) / .3);box-shadow:var(--select-focus-shadow)}.select-container__inner.svelte-6e6oqn>span.svelte-6e6oqn.svelte-6e6oqn{cursor:pointer;padding:0.75rem;color:rgb(var(--select-main-color));transition-property:transform;transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1);transition-duration:150ms;transition-duration:var(--animation-input)}.select-container__inner.svelte-6e6oqn>.selected-options.svelte-6e6oqn.svelte-6e6oqn{display:flex;width:100%;flex-wrap:wrap;align-items:center;gap:0.75rem}.select-container__inner.svelte-6e6oqn>.selected-options.svelte-6e6oqn>input.svelte-6e6oqn{width:100%;-webkit-appearance:none;-moz-appearance:none;appearance:none;border-width:0px;background-color:transparent;padding:0.75rem;outline:2px solid transparent !important;outline-offset:2px !important}.select-container.primary.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--primary);--select-border-color:var(--primary);--select-shadow-color:var(--primary)}.select-container.secondary.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--secondary);--select-border-color:var(--secondary);--select-shadow-color:var(--secondary)}.select-container.neutral.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--neutral);--select-border-color:var(--neutral);--select-shadow-color:var(--neutral)}.select-container.success.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--success);--select-border-color:var(--success);--select-shadow-color:var(--success)}.select-container.info.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--info);--select-border-color:var(--info);--select-shadow-color:var(--info)}.select-container.warning.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--warning);--select-border-color:var(--warning);--select-shadow-color:var(--warning)}.select-container.error.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--error);--select-border-color:var(--error);--select-shadow-color:var(--error)}.select-container.smoke.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--smoke);--select-border-color:var(--smoke);--select-shadow-color:var(--smoke)}.select-dropdown.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--primary);--select-background-color:var(--base1);--select-text-color:var(--basec);z-index:var(--dropdown-z-index);display:flex;max-height:300px;flex-direction:column;overflow:hidden;background-color:rgb(var(--select-background-color));color:rgb(var(--select-text-color));--tw-shadow:0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);--tw-shadow-colored:0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);border-radius:var(--rounded-box)}.select-dropdown__container.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{display:flex;max-height:100%;flex-direction:column;gap:0px;overflow-y:auto}.select-dropdown.svelte-6e6oqn .option.svelte-6e6oqn.svelte-6e6oqn{position:relative;cursor:pointer;scroll-margin-top:0.5rem;scroll-margin-bottom:0.5rem;padding-top:0.75rem;padding-bottom:0.75rem;padding-left:1rem;padding-right:1rem}.select-dropdown.svelte-6e6oqn .option.svelte-6e6oqn.svelte-6e6oqn:hover{background-color:rgb(var(--select-main-color)/.1)}.select-dropdown.svelte-6e6oqn .option[data-highlighted].svelte-6e6oqn.svelte-6e6oqn{background-color:rgb(var(--select-main-color)/.1);color:rgb(var(--select-main-color))}.select-dropdown.svelte-6e6oqn .option[data-disabled].svelte-6e6oqn.svelte-6e6oqn{opacity:0.5}.select-dropdown.primary.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--primary)}.select-dropdown.secondary.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--secondary)}.select-dropdown.neutral.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--neutral)}.select-dropdown.success.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--success)}.select-dropdown.info.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--info)}.select-dropdown.warning.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--warning)}.select-dropdown.error.svelte-6e6oqn.svelte-6e6oqn.svelte-6e6oqn{--select-main-color:var(--error)}',
  map: '{"version":3,"file":"Select.svelte","sources":["Select.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\">export const REQUIRED = 1;\\nexport const MULTIPLE = 2;\\nexport const TRANSPARENT = 4;\\nexport const UNDERLINE = 8;\\nexport const DEFAULT = \\"\\";\\nexport const PRIMARY = \\"primary\\";\\nexport const SECONDARY = \\"secondary\\";\\nexport const NEUTRAL = \\"neutral\\";\\nexport const SUCCESS = \\"success\\";\\nexport const INFO = \\"info\\";\\nexport const WARNING = \\"warning\\";\\nexport const ERROR = \\"error\\";\\nexport const SMOKE = \\"smoke\\";\\n<\/script>\\r\\n\\r\\n<script lang=\\"ts\\">import { fade, fly } from \\"svelte/transition\\";\\nimport { createCombobox, melt } from \\"@melt-ui/svelte\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nconst toOption = (opt) => ({\\n  value: opt,\\n  label: opt.label,\\n  disabled: false\\n});\\nlet selectOptions = [];\\nconst LoadSelectOptions = function() {\\n  let temp = [];\\n  options.forEach((opt) => {\\n    if (typeof opt == \\"object\\") {\\n      temp.push({\\n        value: opt[valueKey],\\n        label: opt[labelKey]\\n      });\\n      return;\\n    }\\n    temp.push({\\n      value: opt,\\n      label: `${opt}`\\n    });\\n  });\\n  selectOptions = temp;\\n};\\nexport const SetValue = function(val) {\\n  value = val ?? null;\\n};\\nexport const GetValue = function() {\\n  return value;\\n};\\nexport const Value = function(val) {\\n  if (val) SetValue(val);\\n  return GetValue();\\n};\\nconst UpdateSelection = function() {\\n  if (Array.isArray(value)) {\\n    const opt2 = selectOptions.filter((opt3) => value.includes(opt3.value));\\n    selected.set(opt2.map((opt3) => toOption(opt3)));\\n    return;\\n  }\\n  const opt = selectOptions.find((opt2) => opt2.value == value);\\n  if (!opt) return;\\n  selected.set(toOption(opt));\\n};\\nconst UpdateValue = function() {\\n  if (!$selected) {\\n    if (multiple) value = [];\\n    else value = null;\\n    return;\\n  }\\n  if (Array.isArray($selected)) {\\n    value = $selected.map((v) => v.value.value);\\n    return;\\n  }\\n  value = $selected?.value.value;\\n};\\nlet className = \\"\\";\\nexport { className as class };\\nexport let dropdownClass = \\"\\";\\nexport let style = \\"\\";\\nexport let dropdownStyle = \\"\\";\\nexport let name;\\nexport let label = \\"\\";\\nexport let min = null;\\nexport let max = null;\\nexport let error = null;\\nexport let palette = DEFAULT;\\nexport let flags = 0;\\nexport let searchText = \\"Search options...\\";\\nexport let noResultsText = \\"No results found\\";\\nexport let valueKey = \\"value\\";\\nexport let labelKey = \\"label\\";\\nexport let options = [];\\nexport let value = null;\\n$: options, LoadSelectOptions();\\n$: label = label.trim();\\n$: min = min == null ? null : min < 0 ? 0 : min;\\n$: max = max == null ? null : max < 0 ? 0 : max;\\n$: multiple = (flags & MULTIPLE) == 1;\\n$: ({\\n  elements: { menu, input, option },\\n  states: { open, inputValue, touchedInput, selected },\\n  helpers: { isSelected }\\n} = createCombobox({\\n  forceVisible: true,\\n  multiple\\n}));\\n$: if (!$open && !Array.isArray($selected)) {\\n  $inputValue = $selected?.label ?? \\"\\";\\n}\\n$: filteredItems = $touchedInput ? selectOptions.filter(({ label: label2 }) => {\\n  const normalizedInput = $inputValue.toLowerCase();\\n  return label2.toLowerCase().includes(normalizedInput);\\n}) : selectOptions;\\n$: $selected, UpdateValue();\\n$: value, UpdateSelection();\\n<\/script>\\r\\n\\r\\n<div \\r\\n    class={\\"select-container\\" + cn(className) + cn(palette)} \\r\\n    class:required={flags & REQUIRED}\\r\\n    class:transparent={flags & TRANSPARENT}\\r\\n    class:underlined={flags & UNDERLINE}\\r\\n    style={style}\\r\\n>\\r\\n    <slot name=\\"label\\">\\r\\n        {#if label && label != \\"\\"}\\r\\n            <label for={name}>{label}</label>\\r\\n        {/if}\\r\\n    </slot>\\r\\n\\r\\n    <div class=\\"select-container__inner\\">\\r\\n        <div class=\\"selected-options\\">\\r\\n            <input\\r\\n                {...$input} use:$input.action\\r\\n                placeholder={searchText}\\r\\n            />\\r\\n        </div>\\r\\n        <!-- svelte-ignore a11y-click-events-have-key-events -->\\r\\n        <!-- svelte-ignore a11y-no-static-element-interactions -->\\r\\n        <span on:click={() => $open = !$open} class={\\"fa-solid fa-caret-down\\" + ($open ? cn(\\"rotate-180\\") : \\"\\")}></span>\\r\\n    </div>\\r\\n\\r\\n    {#if error}\\r\\n        <span class=\\"text-error\\" transition:fade>{ error }</span>\\r\\n    {/if}\\r\\n\\r\\n    {#if $open}\\r\\n        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->\\r\\n        <ul\\r\\n            class={\\"select-dropdown\\" + cn(dropdownClass) + cn(palette)}\\r\\n            {...$menu} use:$menu.action\\r\\n            transition:fly={{ duration: 150, y: -5 }}\\r\\n            style={dropdownStyle}\\r\\n            tabindex=\\"-1\\"\\r\\n        >\\r\\n            <div\\r\\n            class=\\"select-dropdown__container\\"\\r\\n            tabindex=\\"-1\\"\\r\\n            >\\r\\n                {#each filteredItems as item, idx (idx)}\\r\\n                    {@const __MELTUI_BUILDER_0__ = $option(toOption(item))}\\n                    <li\\r\\n                        {...__MELTUI_BUILDER_0__} use:__MELTUI_BUILDER_0__.action\\r\\n                        class=\\"option\\"\\r\\n                    >\\r\\n                        {#if $isSelected(item)}\\r\\n                            <div class=\\"absolute left-2 top-1/2 -translate-y-1/2 z-10 text-primary\\">\\r\\n                                <i class=\\"fa-solid fa-check\\"></i>\\r\\n                            </div>\\r\\n                        {/if}\\r\\n                        <div class=\\"pl-4\\">\\r\\n                            <span>{item.label}</span>\\r\\n                        </div>\\r\\n                    </li>\\r\\n                {:else}\\r\\n                    <li class=\\"relative cursor-pointer rounded-md py-1 pl-8 pr-4\\">\\r\\n                        { noResultsText }\\r\\n                    </li>\\r\\n                {/each}\\r\\n            </div>\\r\\n        </ul>\\r\\n    {/if}\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    .select-container {\\r\\n        --select-main-color: var(--primary);\\r\\n        --select-background-color: var(--base1);\\r\\n        --select-border-color: var(--base3);\\r\\n        --select-shadow-color: var(--primary);\\r\\n        display: flex;\\r\\n        width: 100%;\\r\\n        flex-direction: column;\\r\\n        gap: 0.5rem;\\r\\n    }\\r\\n\\r\\n        .select-container.transparent >  .select-container__inner {\\r\\n        background-color: transparent;\\r\\n}\\r\\n\\r\\n        .select-container.underlined > .select-container__inner {\\r\\n        border-radius: 0px;\\r\\n        border-top-width: 0px;\\r\\n        border-left-width: 0px;\\r\\n        border-right-width: 0px;\\r\\n}\\r\\n\\r\\n        .select-container.underlined > .select-container__inner:focus-within {\\r\\n                box-shadow: none;\\r\\n            }\\r\\n\\r\\n        .select-container.required > :global(label)::after {\\r\\n        margin-left: 0.125rem;\\r\\n        --tw-text-opacity: 1;\\r\\n        color: rgb(var(--error) / var(--tw-text-opacity));\\r\\n        --tw-content: \\"*\\";\\r\\n        content: var(--tw-content);\\r\\n}\\r\\n\\r\\n        .select-container__inner {\\r\\n        margin-top: 0.5rem;\\r\\n        display: flex;\\r\\n        width: 100%;\\r\\n        flex-direction: row;\\r\\n        align-items: center;\\r\\n        overflow: hidden;\\r\\n        border-width: 1px;\\r\\n        border-color: rgb(var(--select-border-color));\\r\\n        background-color: rgb(var(--select-background-color));\\r\\n        border-radius: var(--rounded-input);\\r\\n        transition-property: box-shadow;\\r\\n        transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\\r\\n        transition-duration: 150ms;\\r\\n        transition-duration: var(--animation-input);\\r\\n}\\r\\n\\r\\n        .select-container__inner:focus-within{\\r\\n                --select-focus-shadow: 0 0 0 0.25rem rgba(var(--select-shadow-color) / .3);\\r\\n\\r\\n                box-shadow: var(--select-focus-shadow);\\r\\n            }\\r\\n\\r\\n        .select-container__inner > span {\\r\\n        cursor: pointer;\\r\\n        padding: 0.75rem;\\r\\n        color: rgb(var(--select-main-color));\\r\\n        transition-property: transform;\\r\\n        transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\\r\\n        transition-duration: 150ms;\\r\\n        transition-duration: var(--animation-input);\\r\\n}\\r\\n\\r\\n        .select-container__inner > .selected-options {\\r\\n        display: flex;\\r\\n        width: 100%;\\r\\n        flex-wrap: wrap;\\r\\n        align-items: center;\\r\\n        gap: 0.75rem;\\r\\n}\\r\\n\\r\\n        .select-container__inner > .selected-options > input {\\r\\n        width: 100%;\\r\\n        -webkit-appearance: none;\\r\\n           -moz-appearance: none;\\r\\n                appearance: none;\\r\\n        border-width: 0px;\\r\\n        background-color: transparent;\\r\\n        padding: 0.75rem;\\r\\n        outline: 2px solid transparent !important;\\r\\n        outline-offset: 2px !important;\\r\\n}\\r\\n\\r\\n        .select-container.primary {\\r\\n            --select-main-color: var(--primary);\\r\\n            --select-border-color: var(--primary);\\r\\n            --select-shadow-color: var(--primary);\\r\\n        }\\r\\n\\r\\n        .select-container.secondary {\\r\\n            --select-main-color: var(--secondary);\\r\\n            --select-border-color: var(--secondary);\\r\\n            --select-shadow-color: var(--secondary);\\r\\n        }\\r\\n\\r\\n        .select-container.neutral {\\r\\n            --select-main-color: var(--neutral);\\r\\n            --select-border-color: var(--neutral);\\r\\n            --select-shadow-color: var(--neutral);\\r\\n        }\\r\\n\\r\\n        .select-container.success {\\r\\n            --select-main-color: var(--success);\\r\\n            --select-border-color: var(--success);\\r\\n            --select-shadow-color: var(--success);\\r\\n        }\\r\\n\\r\\n        .select-container.info {\\r\\n            --select-main-color: var(--info);\\r\\n            --select-border-color: var(--info);\\r\\n            --select-shadow-color: var(--info);\\r\\n        }\\r\\n\\r\\n        .select-container.warning {\\r\\n            --select-main-color: var(--warning);\\r\\n            --select-border-color: var(--warning);\\r\\n            --select-shadow-color: var(--warning);\\r\\n        }\\r\\n\\r\\n        .select-container.error {\\r\\n            --select-main-color: var(--error);\\r\\n            --select-border-color: var(--error);\\r\\n            --select-shadow-color: var(--error);\\r\\n        }\\r\\n\\r\\n        .select-container.smoke {\\r\\n            --select-main-color: var(--smoke);\\r\\n            --select-border-color: var(--smoke);\\r\\n            --select-shadow-color: var(--smoke);\\r\\n        }\\r\\n\\r\\n    .select-dropdown {\\r\\n        --select-main-color: var(--primary);\\r\\n        --select-background-color: var(--base1);\\r\\n        --select-text-color: var(--basec);\\r\\n        z-index: var(--dropdown-z-index);\\r\\n        display: flex;\\r\\n        max-height: 300px;\\r\\n        flex-direction: column;\\r\\n        overflow: hidden;\\r\\n        background-color: rgb(var(--select-background-color));\\r\\n        color: rgb(var(--select-text-color));\\r\\n        --tw-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);\\r\\n        --tw-shadow-colored: 0 10px 15px -3px var(--tw-shadow-color), 0 4px 6px -4px var(--tw-shadow-color);\\r\\n        box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);\\r\\n        border-radius: var(--rounded-box);\\r\\n    }\\r\\n\\r\\n    .select-dropdown__container {\\r\\n        display: flex;\\r\\n        max-height: 100%;\\r\\n        flex-direction: column;\\r\\n        gap: 0px;\\r\\n        overflow-y: auto;\\r\\n}\\r\\n\\r\\n    .select-dropdown .option {\\r\\n        position: relative;\\r\\n        cursor: pointer;\\r\\n        scroll-margin-top: 0.5rem;\\r\\n        scroll-margin-bottom: 0.5rem;\\r\\n        padding-top: 0.75rem;\\r\\n        padding-bottom: 0.75rem;\\r\\n        padding-left: 1rem;\\r\\n        padding-right: 1rem;\\r\\n}\\r\\n\\r\\n    .select-dropdown .option:hover {\\r\\n        background-color: rgb(var(--select-main-color)/.1);\\r\\n}\\r\\n\\r\\n    .select-dropdown .option[data-highlighted] {\\r\\n        background-color: rgb(var(--select-main-color)/.1);\\r\\n        color: rgb(var(--select-main-color));\\r\\n}\\r\\n\\r\\n    .select-dropdown .option[data-disabled] {\\r\\n        opacity: 0.5;\\r\\n}\\r\\n\\r\\n    .select-dropdown.primary {\\r\\n            --select-main-color: var(--primary);\\r\\n        }\\r\\n\\r\\n    .select-dropdown.secondary {\\r\\n            --select-main-color: var(--secondary);\\r\\n        }\\r\\n\\r\\n    .select-dropdown.neutral {\\r\\n            --select-main-color: var(--neutral);\\r\\n        }\\r\\n\\r\\n    .select-dropdown.success {\\r\\n            --select-main-color: var(--success);\\r\\n        }\\r\\n\\r\\n    .select-dropdown.info {\\r\\n            --select-main-color: var(--info);\\r\\n        }\\r\\n\\r\\n    .select-dropdown.warning {\\r\\n            --select-main-color: var(--warning);\\r\\n        }\\r\\n\\r\\n    .select-dropdown.error {\\r\\n            --select-main-color: var(--error);\\r\\n        }\\r\\n</style>"],"names":[],"mappings":"AAuLI,2DAAkB,CACd,mBAAmB,CAAE,cAAc,CACnC,yBAAyB,CAAE,YAAY,CACvC,qBAAqB,CAAE,YAAY,CACnC,qBAAqB,CAAE,cAAc,CACrC,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,IAAI,CACX,cAAc,CAAE,MAAM,CACtB,GAAG,CAAE,MACT,CAEI,iBAAiB,0BAAY,CAAI,oDAAyB,CAC1D,gBAAgB,CAAE,WAC1B,CAEQ,iBAAiB,yBAAW,CAAG,oDAAyB,CACxD,aAAa,CAAE,GAAG,CAClB,gBAAgB,CAAE,GAAG,CACrB,iBAAiB,CAAE,GAAG,CACtB,kBAAkB,CAAE,GAC5B,CAEQ,iBAAiB,yBAAW,CAAG,oDAAwB,aAAc,CAC7D,UAAU,CAAE,IAChB,CAEJ,iBAAiB,uBAAS,CAAW,KAAM,OAAQ,CACnD,WAAW,CAAE,QAAQ,CACrB,iBAAiB,CAAE,CAAC,CACpB,KAAK,CAAE,IAAI,IAAI,OAAO,CAAC,CAAC,CAAC,CAAC,IAAI,iBAAiB,CAAC,CAAC,CACjD,YAAY,CAAE,GAAG,CACjB,OAAO,CAAE,IAAI,YAAY,CACjC,CAEQ,kEAAyB,CACzB,UAAU,CAAE,MAAM,CAClB,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,IAAI,CACX,cAAc,CAAE,GAAG,CACnB,WAAW,CAAE,MAAM,CACnB,QAAQ,CAAE,MAAM,CAChB,YAAY,CAAE,GAAG,CACjB,YAAY,CAAE,IAAI,IAAI,qBAAqB,CAAC,CAAC,CAC7C,gBAAgB,CAAE,IAAI,IAAI,yBAAyB,CAAC,CAAC,CACrD,aAAa,CAAE,IAAI,eAAe,CAAC,CACnC,mBAAmB,CAAE,UAAU,CAC/B,0BAA0B,CAAE,aAAa,GAAG,CAAC,CAAC,CAAC,CAAC,CAAC,GAAG,CAAC,CAAC,CAAC,CAAC,CACxD,mBAAmB,CAAE,KAAK,CAC1B,mBAAmB,CAAE,IAAI,iBAAiB,CAClD,CAEQ,kEAAwB,aAAa,CAC7B,qBAAqB,CAAE,mDAAmD,CAE1E,UAAU,CAAE,IAAI,qBAAqB,CACzC,CAEJ,sCAAwB,CAAG,gCAAK,CAChC,MAAM,CAAE,OAAO,CACf,OAAO,CAAE,OAAO,CAChB,KAAK,CAAE,IAAI,IAAI,mBAAmB,CAAC,CAAC,CACpC,mBAAmB,CAAE,SAAS,CAC9B,0BAA0B,CAAE,aAAa,GAAG,CAAC,CAAC,CAAC,CAAC,CAAC,GAAG,CAAC,CAAC,CAAC,CAAC,CACxD,mBAAmB,CAAE,KAAK,CAC1B,mBAAmB,CAAE,IAAI,iBAAiB,CAClD,CAEQ,sCAAwB,CAAG,6CAAkB,CAC7C,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,IAAI,CACX,SAAS,CAAE,IAAI,CACf,WAAW,CAAE,MAAM,CACnB,GAAG,CAAE,OACb,CAEQ,sCAAwB,CAAG,+BAAiB,CAAG,mBAAM,CACrD,KAAK,CAAE,IAAI,CACX,kBAAkB,CAAE,IAAI,CACrB,eAAe,CAAE,IAAI,CAChB,UAAU,CAAE,IAAI,CACxB,YAAY,CAAE,GAAG,CACjB,gBAAgB,CAAE,WAAW,CAC7B,OAAO,CAAE,OAAO,CAChB,OAAO,CAAE,GAAG,CAAC,KAAK,CAAC,WAAW,CAAC,UAAU,CACzC,cAAc,CAAE,GAAG,CAAC,UAC5B,CAEQ,iBAAiB,kDAAS,CACtB,mBAAmB,CAAE,cAAc,CACnC,qBAAqB,CAAE,cAAc,CACrC,qBAAqB,CAAE,cAC3B,CAEA,iBAAiB,oDAAW,CACxB,mBAAmB,CAAE,gBAAgB,CACrC,qBAAqB,CAAE,gBAAgB,CACvC,qBAAqB,CAAE,gBAC3B,CAEA,iBAAiB,kDAAS,CACtB,mBAAmB,CAAE,cAAc,CACnC,qBAAqB,CAAE,cAAc,CACrC,qBAAqB,CAAE,cAC3B,CAEA,iBAAiB,kDAAS,CACtB,mBAAmB,CAAE,cAAc,CACnC,qBAAqB,CAAE,cAAc,CACrC,qBAAqB,CAAE,cAC3B,CAEA,iBAAiB,+CAAM,CACnB,mBAAmB,CAAE,WAAW,CAChC,qBAAqB,CAAE,WAAW,CAClC,qBAAqB,CAAE,WAC3B,CAEA,iBAAiB,kDAAS,CACtB,mBAAmB,CAAE,cAAc,CACnC,qBAAqB,CAAE,cAAc,CACrC,qBAAqB,CAAE,cAC3B,CAEA,iBAAiB,gDAAO,CACpB,mBAAmB,CAAE,YAAY,CACjC,qBAAqB,CAAE,YAAY,CACnC,qBAAqB,CAAE,YAC3B,CAEA,iBAAiB,gDAAO,CACpB,mBAAmB,CAAE,YAAY,CACjC,qBAAqB,CAAE,YAAY,CACnC,qBAAqB,CAAE,YAC3B,CAEJ,0DAAiB,CACb,mBAAmB,CAAE,cAAc,CACnC,yBAAyB,CAAE,YAAY,CACvC,mBAAmB,CAAE,YAAY,CACjC,OAAO,CAAE,IAAI,kBAAkB,CAAC,CAChC,OAAO,CAAE,IAAI,CACb,UAAU,CAAE,KAAK,CACjB,cAAc,CAAE,MAAM,CACtB,QAAQ,CAAE,MAAM,CAChB,gBAAgB,CAAE,IAAI,IAAI,yBAAyB,CAAC,CAAC,CACrD,KAAK,CAAE,IAAI,IAAI,mBAAmB,CAAC,CAAC,CACpC,WAAW,CAAE,kEAAkE,CAC/E,mBAAmB,CAAE,8EAA8E,CACnG,UAAU,CAAE,IAAI,uBAAuB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,gBAAgB,CAAC,UAAU,CAAC,CAAC,CAAC,IAAI,WAAW,CAAC,CACvG,aAAa,CAAE,IAAI,aAAa,CACpC,CAEA,qEAA4B,CACxB,OAAO,CAAE,IAAI,CACb,UAAU,CAAE,IAAI,CAChB,cAAc,CAAE,MAAM,CACtB,GAAG,CAAE,GAAG,CACR,UAAU,CAAE,IACpB,CAEI,8BAAgB,CAAC,mCAAQ,CACrB,QAAQ,CAAE,QAAQ,CAClB,MAAM,CAAE,OAAO,CACf,iBAAiB,CAAE,MAAM,CACzB,oBAAoB,CAAE,MAAM,CAC5B,WAAW,CAAE,OAAO,CACpB,cAAc,CAAE,OAAO,CACvB,YAAY,CAAE,IAAI,CAClB,aAAa,CAAE,IACvB,CAEI,8BAAgB,CAAC,mCAAO,MAAO,CAC3B,gBAAgB,CAAE,IAAI,IAAI,mBAAmB,CAAC,CAAC,EAAE,CACzD,CAEI,8BAAgB,CAAC,OAAO,CAAC,gBAAgB,6BAAE,CACvC,gBAAgB,CAAE,IAAI,IAAI,mBAAmB,CAAC,CAAC,EAAE,CAAC,CAClD,KAAK,CAAE,IAAI,IAAI,mBAAmB,CAAC,CAC3C,CAEI,8BAAgB,CAAC,OAAO,CAAC,aAAa,6BAAE,CACpC,OAAO,CAAE,GACjB,CAEI,gBAAgB,kDAAS,CACjB,mBAAmB,CAAE,cACzB,CAEJ,gBAAgB,oDAAW,CACnB,mBAAmB,CAAE,gBACzB,CAEJ,gBAAgB,kDAAS,CACjB,mBAAmB,CAAE,cACzB,CAEJ,gBAAgB,kDAAS,CACjB,mBAAmB,CAAE,cACzB,CAEJ,gBAAgB,+CAAM,CACd,mBAAmB,CAAE,WACzB,CAEJ,gBAAgB,kDAAS,CACjB,mBAAmB,CAAE,cACzB,CAEJ,gBAAgB,gDAAO,CACf,mBAAmB,CAAE,YACzB"}'
};
const REQUIRED = 1;
const MULTIPLE = 2;
const TRANSPARENT = 4;
const UNDERLINE = 8;
const DEFAULT = "";
const Select = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let multiple;
  let menu;
  let input;
  let option;
  let open;
  let inputValue;
  let touchedInput;
  let selected;
  let isSelected;
  let filteredItems;
  let $selected, $$unsubscribe_selected = noop, $$subscribe_selected = () => ($$unsubscribe_selected(), $$unsubscribe_selected = subscribe(selected, ($$value) => $selected = $$value), selected);
  let $inputValue, $$unsubscribe_inputValue = noop, $$subscribe_inputValue = () => ($$unsubscribe_inputValue(), $$unsubscribe_inputValue = subscribe(inputValue, ($$value) => $inputValue = $$value), inputValue);
  let $touchedInput, $$unsubscribe_touchedInput = noop, $$subscribe_touchedInput = () => ($$unsubscribe_touchedInput(), $$unsubscribe_touchedInput = subscribe(touchedInput, ($$value) => $touchedInput = $$value), touchedInput);
  let $open, $$unsubscribe_open = noop, $$subscribe_open = () => ($$unsubscribe_open(), $$unsubscribe_open = subscribe(open, ($$value) => $open = $$value), open);
  let $input, $$unsubscribe_input = noop, $$subscribe_input = () => ($$unsubscribe_input(), $$unsubscribe_input = subscribe(input, ($$value) => $input = $$value), input);
  let $menu, $$unsubscribe_menu = noop, $$subscribe_menu = () => ($$unsubscribe_menu(), $$unsubscribe_menu = subscribe(menu, ($$value) => $menu = $$value), menu);
  let $option, $$unsubscribe_option = noop, $$subscribe_option = () => ($$unsubscribe_option(), $$unsubscribe_option = subscribe(option, ($$value) => $option = $$value), option);
  let $isSelected, $$unsubscribe_isSelected = noop, $$subscribe_isSelected = () => ($$unsubscribe_isSelected(), $$unsubscribe_isSelected = subscribe(isSelected, ($$value) => $isSelected = $$value), isSelected);
  const toOption = (opt) => ({
    value: opt,
    label: opt.label,
    disabled: false
  });
  let selectOptions = [];
  const LoadSelectOptions = function() {
    let temp = [];
    options.forEach((opt) => {
      if (typeof opt == "object") {
        temp.push({
          value: opt[valueKey],
          label: opt[labelKey]
        });
        return;
      }
      temp.push({ value: opt, label: `${opt}` });
    });
    selectOptions = temp;
  };
  const SetValue = function(val) {
    value = val ?? null;
  };
  const GetValue = function() {
    return value;
  };
  const Value = function(val) {
    if (val) SetValue(val);
    return GetValue();
  };
  const UpdateSelection = function() {
    if (Array.isArray(value)) {
      const opt2 = selectOptions.filter((opt3) => value.includes(opt3.value));
      selected.set(opt2.map((opt3) => toOption(opt3)));
      return;
    }
    const opt = selectOptions.find((opt2) => opt2.value == value);
    if (!opt) return;
    selected.set(toOption(opt));
  };
  const UpdateValue = function() {
    if (!$selected) {
      if (multiple) value = [];
      else value = null;
      return;
    }
    if (Array.isArray($selected)) {
      value = $selected.map((v) => v.value.value);
      return;
    }
    value = $selected?.value.value;
  };
  let { class: className = "" } = $$props;
  let { dropdownClass = "" } = $$props;
  let { style = "" } = $$props;
  let { dropdownStyle = "" } = $$props;
  let { name: name2 } = $$props;
  let { label = "" } = $$props;
  let { min = null } = $$props;
  let { max = null } = $$props;
  let { error = null } = $$props;
  let { palette = DEFAULT } = $$props;
  let { flags = 0 } = $$props;
  let { searchText = "Search options..." } = $$props;
  let { noResultsText = "No results found" } = $$props;
  let { valueKey = "value" } = $$props;
  let { labelKey = "label" } = $$props;
  let { options = [] } = $$props;
  let { value = null } = $$props;
  if ($$props.SetValue === void 0 && $$bindings.SetValue && SetValue !== void 0) $$bindings.SetValue(SetValue);
  if ($$props.GetValue === void 0 && $$bindings.GetValue && GetValue !== void 0) $$bindings.GetValue(GetValue);
  if ($$props.Value === void 0 && $$bindings.Value && Value !== void 0) $$bindings.Value(Value);
  if ($$props.class === void 0 && $$bindings.class && className !== void 0) $$bindings.class(className);
  if ($$props.dropdownClass === void 0 && $$bindings.dropdownClass && dropdownClass !== void 0) $$bindings.dropdownClass(dropdownClass);
  if ($$props.style === void 0 && $$bindings.style && style !== void 0) $$bindings.style(style);
  if ($$props.dropdownStyle === void 0 && $$bindings.dropdownStyle && dropdownStyle !== void 0) $$bindings.dropdownStyle(dropdownStyle);
  if ($$props.name === void 0 && $$bindings.name && name2 !== void 0) $$bindings.name(name2);
  if ($$props.label === void 0 && $$bindings.label && label !== void 0) $$bindings.label(label);
  if ($$props.min === void 0 && $$bindings.min && min !== void 0) $$bindings.min(min);
  if ($$props.max === void 0 && $$bindings.max && max !== void 0) $$bindings.max(max);
  if ($$props.error === void 0 && $$bindings.error && error !== void 0) $$bindings.error(error);
  if ($$props.palette === void 0 && $$bindings.palette && palette !== void 0) $$bindings.palette(palette);
  if ($$props.flags === void 0 && $$bindings.flags && flags !== void 0) $$bindings.flags(flags);
  if ($$props.searchText === void 0 && $$bindings.searchText && searchText !== void 0) $$bindings.searchText(searchText);
  if ($$props.noResultsText === void 0 && $$bindings.noResultsText && noResultsText !== void 0) $$bindings.noResultsText(noResultsText);
  if ($$props.valueKey === void 0 && $$bindings.valueKey && valueKey !== void 0) $$bindings.valueKey(valueKey);
  if ($$props.labelKey === void 0 && $$bindings.labelKey && labelKey !== void 0) $$bindings.labelKey(labelKey);
  if ($$props.options === void 0 && $$bindings.options && options !== void 0) $$bindings.options(options);
  if ($$props.value === void 0 && $$bindings.value && value !== void 0) $$bindings.value(value);
  $$result.css.add(css);
  {
    LoadSelectOptions();
  }
  label = label.trim();
  min = min == null ? null : min < 0 ? 0 : min;
  max = max == null ? null : max < 0 ? 0 : max;
  multiple = (flags & MULTIPLE) == 1;
  $$subscribe_menu({ elements: { menu, input, option }, states: { open, inputValue, touchedInput, selected }, helpers: { isSelected } } = createCombobox({ forceVisible: true, multiple }), $$subscribe_input(), $$subscribe_option(), $$subscribe_open(), $$subscribe_inputValue(), $$subscribe_touchedInput(), $$subscribe_selected(), $$subscribe_isSelected());
  {
    if (!$open && !Array.isArray($selected)) {
      set_store_value(inputValue, $inputValue = $selected?.label ?? "", $inputValue);
    }
  }
  filteredItems = $touchedInput ? selectOptions.filter(({ label: label2 }) => {
    const normalizedInput = $inputValue.toLowerCase();
    return label2.toLowerCase().includes(normalizedInput);
  }) : selectOptions;
  {
    UpdateValue();
  }
  {
    UpdateSelection();
  }
  $$unsubscribe_selected();
  $$unsubscribe_inputValue();
  $$unsubscribe_touchedInput();
  $$unsubscribe_open();
  $$unsubscribe_input();
  $$unsubscribe_menu();
  $$unsubscribe_option();
  $$unsubscribe_isSelected();
  return `<div class="${[
    escape(null_to_empty("select-container" + cn(className) + cn(palette)), true) + " svelte-6e6oqn",
    (flags & REQUIRED ? "required" : "") + " " + (flags & TRANSPARENT ? "transparent" : "") + " " + (flags & UNDERLINE ? "underlined" : "")
  ].join(" ").trim()}"${add_attribute("style", style, 0)}>${slots.label ? slots.label({}) : ` ${label && label != "" ? `<label${add_attribute("for", name2, 0)}>${escape(label)}</label>` : ``} `} <div class="select-container__inner svelte-6e6oqn"><div class="selected-options svelte-6e6oqn"><input${spread(
    [
      escape_object($input),
      {
        placeholder: escape_attribute_value(searchText)
      }
    ],
    { classes: "svelte-6e6oqn" }
  )}></div>   <span class="${escape(null_to_empty("fa-solid fa-caret-down" + ($open ? cn("rotate-180") : "")), true) + " svelte-6e6oqn"}"></span></div> ${error ? `<span class="text-error svelte-6e6oqn">${escape(error)}</span>` : ``} ${$open ? ` <ul${spread(
    [
      {
        class: escape_attribute_value("select-dropdown" + cn(dropdownClass) + cn(palette))
      },
      escape_object($menu),
      {
        style: escape_attribute_value(dropdownStyle)
      },
      { tabindex: "-1" }
    ],
    { classes: "svelte-6e6oqn" }
  )}><div class="select-dropdown__container svelte-6e6oqn" tabindex="-1">${filteredItems.length ? each(filteredItems, (item, idx) => {
    let __MELTUI_BUILDER_0__ = $option(toOption(item));
    return ` <li${spread([escape_object(__MELTUI_BUILDER_0__), { class: "option" }], { classes: "svelte-6e6oqn" })}>${$isSelected(item) ? `<div class="absolute left-2 top-1/2 -translate-y-1/2 z-10 text-primary" data-svelte-h="svelte-1y9vexb"><i class="fa-solid fa-check"></i> </div>` : ``} <div class="pl-4"><span>${escape(item.label)}</span></div> </li>`;
  }) : `<li class="relative cursor-pointer rounded-md py-1 pl-8 pr-4">${escape(noResultsText)} </li>`}</div></ul>` : ``} </div>`;
});
const subBG = "data:image/svg+xml,%3csvg%20width='615'%20height='805'%20viewBox='0%200%20615%20805'%20fill='none'%20xmlns='http://www.w3.org/2000/svg'%3e%3cg%20clip-path='url(%23clip0_220_41)'%3e%3crect%20width='615'%20height='805'%20fill='%23DBF2FF'/%3e%3cellipse%20cx='423.5'%20cy='925'%20rx='579.5'%20ry='259'%20fill='%2356C080'/%3e%3c/g%3e%3cdefs%3e%3cclipPath%20id='clip0_220_41'%3e%3crect%20width='615'%20height='805'%20fill='white'/%3e%3c/clipPath%3e%3c/defs%3e%3c/svg%3e";
const smallBush = "/_app/immutable/assets/small_bush.v070rkD3.svg";
const bigBoss = "/_app/immutable/assets/big_boss.Dkq6dD3z.svg";
let transitionDuration = 1;
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $dictionaryStore;
  let $formValues, $$unsubscribe_formValues = noop, $$subscribe_formValues = () => ($$unsubscribe_formValues(), $$unsubscribe_formValues = subscribe(formValues, ($$value) => $formValues = $$value), formValues);
  let title = "";
  let stage = "none";
  const callServiceMethod = getContext("CallServiceMethod");
  getContext("i18nLoadDictionary");
  let showClouds = false;
  let parallaxController;
  let dictionary;
  let projects = [];
  const GetProjects = async function() {
    const errorTitle = GetText("dashboard.toasts.error.title", $dictionaryStore) || "An error occured";
    let resp = await callServiceMethod({
      service: "auth",
      url: "/basic-auth/projects/select",
      showToast: true
    });
    if (!resp) {
      showToast({
        data: {
          title: errorTitle,
          description: "Failed to load project list",
          type: ERROR,
          flags: CHANGE_MAIN_COLOR
        }
      });
      return;
    }
    if (resp.code !== 200) {
      showToast({
        data: {
          title: errorTitle,
          description: resp.error?.message ?? "",
          type: ERROR,
          flags: CHANGE_MAIN_COLOR
        }
      });
      return;
    }
    projects = resp.data?.projects ?? [];
  };
  const SubmitForm = async function(data) {
    switch (stage) {
      case "auth":
        return Auth(data);
      case "project":
        return SelectProject(data);
    }
  };
  const Auth = async function(data) {
    const login = data.values["login"];
    const password = data.values["password"];
    let resp = await callServiceMethod({
      service: "auth",
      url: "/basic-auth/",
      method: "POST",
      headers: {
        "Content-Type": "application/json;charset=utf-8"
      },
      data: { username: login, password },
      showToast: true
    });
    if (!resp || resp.code !== 200) return;
    const serverStage = (resp.headers ?? {})["x-authorization-state"] ?? "auth";
    if (serverStage == "done") return window.location.replace("/");
    stage = "project";
    GetProjects();
  };
  const SelectProject = async function(data) {
    const project = data.values["project"];
    let resp = await callServiceMethod({
      service: "auth",
      url: "/basic-auth/projects/set",
      method: "POST",
      headers: {
        "Content-Type": "application/json;charset=utf-8"
      },
      data: { project_id: project },
      showToast: true
    });
    if (!resp || resp.code !== 200) return;
    window.location.replace("/");
  };
  const Error = function(data) {
    if (data.field) return form.SetFieldError(data.field, data.msg);
    showToast({
      data: {
        title: "Ошибка авторизации",
        description: data.msg,
        type: ERROR,
        flags: CHANGE_MAIN_COLOR
      }
    });
  };
  let form;
  let formFields = [];
  let formValues = writable({});
  $$subscribe_formValues();
  const UpdateFormFields = function() {
    switch (stage) {
      case "auth":
        formFields = [
          new FormField({ name: "login", required: true }),
          new FormField({ name: "password", required: true })
        ];
        break;
      case "project":
        formFields = [
          new FormField({
            name: "project",
            type: "number",
            required: true
          })
        ];
        break;
    }
  };
  let formErrors = {};
  let formErrorsUnsubscriber = null;
  onDestroy(() => {
    if (formErrorsUnsubscriber != null) formErrorsUnsubscriber();
  });
  let $$settled;
  let $$rendered;
  let previous_head = $$result.head;
  do {
    $$settled = true;
    $$result.head = previous_head;
    {
      setTimeout(() => showClouds = true, transitionDuration * 1e3);
    }
    {
      UpdateFormFields();
    }
    {
      formErrorsUnsubscriber = form?.errors?.subscribe((data) => formErrors = data);
    }
    $$rendered = `${$$result.head += `<!-- HEAD_svelte-1uo06u1_START -->${$$result.title = `<title>${escape(title)}</title>`, ""}<!-- HEAD_svelte-1uo06u1_END -->`, ""} <main class="fs-page bg-base-200"><div class="w-[90%] m-auto p-3"${add_attribute("this", parallaxController, 0)}>${validate_component(Form, "Form").$$render(
      $$result,
      {
        fields: formFields,
        OnSubmit: SubmitForm,
        error: Error,
        requiredValueText: GetText("dashboard.pages.auth.form.errors.field_is_required", $dictionaryStore),
        invalidValueText: GetText("dashboard.pages.auth.form.errors.invalid_value", $dictionaryStore),
        this: form,
        values: formValues
      },
      {
        this: ($$value) => {
          form = $$value;
          $$settled = false;
        },
        values: ($$value) => {
          formValues = $$value;
          $$settled = false;
        }
      },
      {
        default: () => {
          return `<div class="relative">${validate_component(Card, "Card").$$render(
            $$result,
            {
              class: "absolute left-1/2 -translate-x-1/2 lg:scale-[.65] 2xl:scale-[.80] origin-center hidden lg:flex",
              style: "width: 1360px;"
            },
            {},
            {
              fallback: () => {
                return `${showClouds ? `<svg class="absolute -top-32 -left-64 z-10" width="740" height="290" viewBox="0 0 740 290" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M93.6837 189.447C81.908 186.013 68.3333 187.975 58.3567 195.335C48.3801 202.695 42.4923 215.452 44.1278 227.718C35.9502 226.41 27.4456 225.265 19.4317 226.9C11.2541 228.536 3.40364 233.769 0.786826 241.62C0.459724 242.765 0.132603 243.91 0.459705 244.891C1.11391 247.344 4.22141 248.162 6.83823 248.489C39.3848 252.251 79.9454 253.886 111.511 242.928C139.805 232.952 111.838 194.681 93.6837 189.447Z" fill="rgb(var(--primary-fg))"></path><path d="M696.043 248.979C703.73 248.816 711.908 248.489 718.613 244.563C725.319 240.638 729.898 231.97 726.627 224.937C723.52 218.395 715.179 216.269 707.982 215.451C695.389 214.307 682.469 215.124 670.202 218.068C671.511 207.765 663.333 197.297 653.193 196.152C642.889 194.844 632.586 203.185 631.441 213.325C634.221 205.638 631.768 196.48 626.207 190.428C620.646 184.377 612.632 180.942 604.618 179.47C596.277 177.998 586.791 179.143 581.394 185.685C578.614 189.12 577.469 193.372 576.161 197.624C568.147 222.975 543.941 251.76 586.628 251.596C622.936 251.596 659.572 249.961 696.043 248.979Z" fill="white"></path><path d="M596.275 221.176C589.733 210.381 570.598 208.255 563.074 219.213C568.308 203.512 571.088 186.34 566.345 170.475C561.766 154.611 548.191 140.545 531.673 139.564C515.154 138.419 498.472 154.12 501.743 170.312C487.35 154.938 465.107 147.578 444.5 151.34C474.43 134.003 491.112 96.3866 484.079 62.5316C479.827 42.2513 465.107 21.1532 446.299 8.39624C457.748 9.05044 468.869 11.8308 479.173 16.9009C497.654 25.8962 512.374 42.5784 519.079 61.8774C522.023 70.2185 525.948 83.7932 522.677 92.6249C519.897 100.639 511.556 104.237 509.103 112.905C520.878 98.8398 540.504 93.7698 558.004 98.8399C575.504 103.91 589.733 117.648 598.401 133.676C613.284 160.498 613.448 195.989 596.275 221.176Z" fill="white"></path><path d="M601.019 242.601C596.112 245.708 590.224 246.526 584.337 247.344C552.117 251.596 518.589 255.685 487.351 246.363C479.991 244.236 472.468 241.292 465.108 243.255C461.183 244.236 457.912 246.526 454.477 248.652C413.917 273.512 359.945 274.657 318.24 251.76C311.043 247.834 304.01 243.092 295.833 241.456C288.146 239.984 280.296 241.62 272.609 243.092C216.838 253.232 159.432 253.395 103.661 243.582C101.371 243.255 99.0814 242.764 97.4459 241.292C91.7217 236.386 97.7731 227.391 103.824 223.138C127.866 206.293 158.287 198.769 187.235 202.694C183.964 185.849 194.759 168.676 209.642 160.008C224.361 151.339 242.352 149.704 259.361 150.358C258.871 134.657 267.048 118.793 280.296 110.125C293.543 101.456 311.207 100.148 325.436 106.854C334.922 74.4706 354.221 44.0501 382.515 25.7325C401.16 13.6297 424.057 7.25121 446.136 8.55962C465.108 21.3166 479.664 42.4146 483.916 62.6949C491.113 96.55 474.267 134.167 444.337 151.503C464.945 147.741 487.351 155.101 501.58 170.475C498.309 154.283 514.991 138.746 531.51 139.727C548.028 140.872 561.603 154.938 566.182 170.638C570.925 186.503 567.982 203.676 562.911 219.377C595.294 186.339 599.22 223.465 600.037 222.321C603.145 222.811 606.252 224.774 607.234 227.881C609.36 233.278 605.925 239.493 601.019 242.601Z" fill="rgb(var(--primary-fg))"></path><path d="M251.349 186.34C268.685 186.012 285.367 199.424 288.802 216.433C300.087 213.98 312.68 217.741 320.694 226.246C328.708 234.751 331.979 247.18 329.036 258.465C252.33 247.508 174.153 270.241 97.1204 262.064C86.8167 260.919 74.5505 257.484 72.2608 247.508C70.6253 240.311 75.0411 233.279 79.9476 227.718C98.1017 207.111 123.943 195.662 151.092 193.863C179.223 192.064 192.961 201.713 215.204 217.414C219.947 200.732 233.031 186.667 251.349 186.34Z" fill="rgb(var(--primary))"></path><path d="M632.912 277.928C629.477 282.998 624.571 284.633 619.828 285.614C617.047 286.105 614.267 286.432 611.323 286.759C611.159 286.269 610.996 285.941 610.996 285.451C606.417 270.241 592.678 256.993 576.814 256.993C566.183 256.993 556.534 262.717 547.538 268.605C545.576 269.913 543.122 271.058 541.323 269.75C540.342 268.932 540.015 267.787 539.524 266.479C538.216 262.717 535.599 260.1 532.328 258.301C537.398 257.32 542.468 257.647 547.538 257.32C548.847 257.32 550.319 256.993 551.3 255.685C552.281 254.213 552.445 251.923 552.772 249.96C553.917 243.418 557.842 238.512 562.094 236.386C566.347 234.259 570.926 234.259 575.505 234.423C587.117 234.75 594.314 239.657 603.963 247.998C612.631 255.521 626.206 250.614 633.239 258.792C637.001 263.208 636.346 272.857 632.912 277.928Z" fill="white"></path><path d="M156.488 252.414C146.839 246.199 134.245 245.054 123.778 249.47C116.909 252.414 110.858 257.647 108.404 264.68C105.951 271.713 107.586 280.381 113.474 284.96C118.381 288.885 125.414 289.54 131.792 289.049C151.745 287.577 171.371 277.764 191.161 281.199C206.371 283.815 225.67 275.801 202.119 268.605C193.614 265.988 184.455 269.423 175.951 266.315C168.264 263.535 163.03 256.666 156.488 252.414Z" fill="rgb(var(--primary-fg))"></path><path d="M611.489 286.596C604.947 287.087 598.241 286.923 591.699 286.759C569.129 286.105 546.559 285.615 523.989 284.96C520.228 284.797 515.485 283.325 514.994 277.764C514.83 275.474 515.485 273.185 516.302 271.386C519.41 264.189 525.134 260.101 530.695 258.629C531.349 258.465 532.003 258.302 532.494 258.138C535.765 259.937 538.218 262.554 539.69 266.316C540.181 267.46 540.508 268.769 541.489 269.587C543.288 270.895 545.905 269.75 547.704 268.442C556.536 262.554 566.349 256.83 576.98 256.83C592.844 256.83 606.582 270.077 611.162 285.287C611.325 285.778 611.325 286.269 611.489 286.596Z" fill="rgb(var(--primary-fg))"></path><path d="M408.686 172.602C412.611 160.826 424.714 152.321 436.98 152.485C424.551 159.845 415.392 172.275 411.957 186.34C405.088 184.705 402.144 176.854 397.401 171.62C389.878 163.279 378.429 162.135 368.289 160.172C382.027 152.158 404.434 153.957 408.686 172.602Z" fill="white"></path><path d="M213.569 235.732C200.322 224.12 181.677 219.213 164.504 222.975C169.083 216.433 178.242 214.143 186.093 215.615C193.943 217.087 200.976 221.503 207.845 225.755C205.392 224.284 212.751 212.017 214.06 211.036C217.167 217.251 213.569 228.699 213.569 235.732Z" fill="rgb(var(--primary-fg))"></path><path d="M606.909 177.671H605.928C606.091 158.7 602.82 142.017 596.115 128.279C588.264 112.088 575.344 99.4942 560.624 93.9335C544.105 87.7185 525.461 90.3353 513.031 100.803C507.634 105.382 503.708 111.106 501.419 117.485L500.438 117.158C502.727 110.616 506.816 104.728 512.377 99.9848C524.97 89.1905 544.106 86.4101 560.951 92.7886C575.998 98.5129 589.082 111.106 597.096 127.625C603.638 141.69 606.909 158.372 606.909 177.671Z" fill="white"></path><path d="M599.054 182.414L599.545 176.526H600.526L600.362 178.652C609.848 164.096 612.629 144.961 607.722 128.442C602.488 111.269 588.914 96.3864 572.231 89.6808C564.545 86.5733 555.713 84.9378 546.063 84.7743C534.288 84.6107 522.839 86.2462 511.718 90.1714L511.391 89.1901C522.676 85.4284 534.288 83.6294 546.063 83.7929C555.876 83.9565 564.872 85.7555 572.559 88.6995C589.568 95.405 603.47 110.615 608.703 128.115C613.937 145.615 610.666 165.895 600.199 180.779L599.054 182.414Z" fill="white"></path><path d="M716.488 236.059C715.507 236.059 714.525 236.059 713.544 235.896V234.914C718.123 235.242 722.539 234.751 726.955 233.77C728.1 233.442 729.245 233.115 730.226 232.461C731.208 231.643 731.862 230.335 732.025 228.699C732.516 224.774 730.553 220.685 726.955 217.578C723.848 214.961 720.086 213.162 715.507 211.527C709.946 209.564 703.24 207.438 696.371 207.601C689.175 207.765 682.796 210.382 678.871 214.798L678.054 215.779V214.471C677.563 206.457 675.11 198.77 670.857 192.064C670.039 190.756 669.222 189.611 667.913 188.793C666.278 187.812 664.315 187.648 662.68 187.648C647.142 187.812 632.75 199.915 629.97 215.125L628.988 214.961C631.932 199.26 646.815 186.831 662.843 186.667C664.642 186.667 666.768 186.831 668.731 187.975C670.203 188.793 671.021 190.265 672.002 191.574C676.254 198.116 678.708 205.639 679.362 213.326C683.614 209.237 689.829 206.784 696.862 206.62C703.895 206.457 710.764 208.583 716.488 210.545C721.231 212.181 725.156 213.98 728.264 216.76C732.025 220.031 734.151 224.611 733.661 228.863C733.497 230.171 733.007 232.134 731.535 233.279C730.39 234.26 729.081 234.587 727.937 234.914C723.684 235.569 720.086 236.059 716.488 236.059Z" fill="white"></path><path d="M542.794 259.283L541.812 259.119C543.121 243.909 555.06 230.334 569.943 227.063C573.541 226.246 576.485 226.246 579.102 226.9C582.21 227.718 584.663 229.517 586.625 231.152C592.023 235.405 596.602 240.802 599.873 246.853L600.364 246.69L601.018 249.143L600.037 249.47C596.602 242.764 591.696 236.713 585.808 231.97C584.009 230.498 581.555 228.699 578.611 228.045C576.158 227.391 573.378 227.391 569.943 228.208C555.714 231.316 544.102 244.4 542.794 259.283Z" fill="white"></path><path d="M413.266 174.4H412.285C412.775 166.059 408.196 157.554 401 153.302C393.803 149.049 384.154 149.213 377.121 153.629C373.85 155.592 371.233 158.535 368.617 161.152L367.799 160.498C370.416 157.718 373.196 154.937 376.467 152.811C383.827 148.068 393.967 148.068 401.49 152.484C408.196 156.409 412.612 163.769 413.266 171.619H413.593L413.266 174.4Z" fill="white"></path><path d="M249.874 163.933H248.893C249.71 149.868 250.364 135.312 255.762 121.737C261.813 106.69 273.098 96.0595 286.019 93.2791C293.705 91.6436 302.21 93.4427 308.916 98.0221C313.659 101.293 316.93 105.709 318.565 110.616C321.673 86.0829 327.233 66.9475 335.902 50.4289C347.023 29.3308 363.378 13.3028 382.35 5.45239C403.775 -3.5429 428.471 -1.25319 445.317 11.1767L444.663 11.9944C427.981 -0.108337 403.612 -2.39804 382.677 6.4337C364.032 14.2841 347.677 29.985 336.719 50.9195C327.888 67.9288 322.327 87.7184 319.219 113.559L318.729 117.321L318.238 113.559C317.42 107.835 313.822 102.438 308.425 98.8399C301.883 94.424 293.706 92.6249 286.346 94.2605C273.752 97.0408 262.794 107.345 256.907 122.064C252.327 133.513 251.182 145.779 250.365 158.045L249.874 163.933Z" fill="rgb(var(--primary-fg))"></path><path d="M9.12956 242.764C7.82115 242.764 6.51271 242.601 5.2043 242.11C2.75104 241.292 0.951981 239.493 0.297778 237.204C-0.683527 233.932 0.951961 230.498 2.42392 228.208C9.1295 217.905 23.522 213.816 34.6434 218.886C32.5173 209.727 34.3163 199.423 39.55 190.592C45.1107 181.269 53.6153 174.727 63.2648 172.437C72.7508 170.148 83.5452 172.11 92.5405 177.835C101.699 183.559 108.078 192.391 110.04 201.877L109.059 202.04C107.096 192.718 100.881 184.213 92.0497 178.652C83.218 173.092 72.7508 171.129 63.592 173.419C54.2696 175.708 45.9284 182.087 40.5312 191.082C35.1341 200.077 33.4986 210.545 35.9518 219.704L36.2789 220.848L35.2977 220.358C24.5033 214.797 9.94731 218.559 3.40528 228.699C2.09688 230.825 0.624899 233.932 1.44265 236.876C2.09686 239.33 4.22298 240.638 5.69493 241.129C8.96595 242.274 13.3818 241.619 16.3257 239.493L16.98 240.311C14.6903 241.947 11.9099 242.764 9.12956 242.764Z" fill="rgb(var(--primary-fg))"></path><path d="M205.229 275.802L204.902 271.55L205.883 271.386V272.204C206.21 270.732 205.72 269.097 204.738 267.788C203.593 266.153 201.794 265.008 199.014 264.027C192.472 261.737 185.112 261.737 178.57 264.354L177.916 264.517L177.752 263.863C176.771 257.975 172.519 252.578 165.486 247.508C159.762 243.583 150.93 239.003 140.626 239.167C131.14 239.33 122.636 243.91 118.22 250.942C115.93 254.704 114.622 259.284 114.949 263.863H113.968C113.64 259.12 114.949 254.377 117.402 250.288C121.981 242.928 130.977 238.185 140.626 238.022C151.094 237.858 160.252 242.601 166.14 246.527C173.173 251.433 177.425 256.994 178.734 262.882C185.276 260.428 192.799 260.428 199.341 262.718C202.285 263.699 204.248 265.008 205.556 266.807C207.028 268.77 207.519 271.877 205.883 274.167L205.229 275.802Z" fill="rgb(var(--primary-fg))"></path><path d="M270.319 238.349L269.665 238.022C261.487 232.788 250.693 231.643 241.534 234.914L241.207 233.933C250.366 230.498 261.16 231.643 269.501 236.713C273.099 223.302 285.693 212.835 299.595 212.017C300.576 212.017 301.557 211.854 301.721 211.199V211.036C301.557 210.872 301.23 210.709 301.23 210.709H300.249C300.249 210.218 300.412 209.891 300.739 209.727C301.23 209.564 301.884 209.727 302.375 210.218C302.866 210.545 302.866 211.199 302.702 211.526C302.211 212.835 300.412 212.998 299.595 212.998C285.856 213.98 273.426 224.283 270.155 237.694L270.319 238.349Z" fill="rgb(var(--primary-fg))"></path><path d="M590.225 224.284L588.099 217.087L589.081 216.76L590.062 220.195C591.207 215.452 592.025 210.709 591.207 205.966C590.389 201.223 587.118 195.662 581.721 194.354C575.342 192.718 568.964 197.625 564.711 201.387L562.095 203.676L563.894 200.732C573.38 184.868 574.525 164.588 566.674 147.742C562.422 138.583 556.207 132.041 549.011 129.261C544.268 127.461 538.38 126.971 531.674 127.789C520.226 129.424 510.413 134.331 503.707 141.69C496.347 150.032 493.567 161.48 496.511 171.13L495.53 171.457C492.422 161.48 495.366 149.541 502.889 141.036C509.595 133.513 519.735 128.443 531.347 126.807C538.216 125.826 544.268 126.317 549.174 128.279C556.534 131.223 563.076 137.929 567.328 147.251C574.852 163.443 574.198 182.905 565.856 198.606C570.109 195.172 575.669 191.737 581.557 193.209C587.445 194.681 590.88 200.405 591.861 205.639C592.842 211.2 591.697 216.76 590.389 222.157L590.225 224.284Z" fill="rgb(var(--primary-fg))"></path><path d="M240.881 164.751C234.993 147.414 235.32 129.751 241.698 115.195C248.731 98.8397 263.778 87.064 279.969 85.265C297.469 83.3024 315.296 93.7696 320.694 109.307L319.712 109.634C314.479 94.5874 297.142 84.2837 280.133 86.2463C264.268 88.0454 249.549 99.6575 242.68 115.685C236.465 130.078 236.138 147.414 241.862 164.587L240.881 164.751Z" fill="rgb(var(--primary-fg))"></path><path d="M344.412 193.536L342.94 191.41C338.197 184.868 330.51 180.779 322.332 180.289C314.155 179.961 306.141 183.232 300.744 189.284L299.926 188.63C305.487 182.251 313.828 178.817 322.332 179.144C330.019 179.471 337.379 183.233 342.449 189.12C341.958 183.887 344.575 177.999 349.482 172.602C359.622 161.971 374.832 156.737 389.225 159.027L389.061 160.008C374.832 157.882 359.949 162.952 350.136 173.419C347.519 176.2 341.795 183.396 343.594 191.083L344.412 193.536Z" fill="white"></path><path d="M708.473 244.4L708.309 243.419C714.687 241.947 721.229 240.311 727.444 238.349C731.37 237.204 735.786 235.568 737.748 232.134C739.057 230.008 739.384 227.064 738.729 223.629C737.421 215.942 732.351 208.419 724.501 202.531C719.758 199.097 712.561 194.844 704.384 195.171C691.627 195.662 681.814 206.783 673.636 217.087L672.818 216.433C681.159 205.966 690.973 194.681 704.22 194.027C712.561 193.699 720.085 197.952 724.991 201.55C733.169 207.601 738.402 215.288 739.711 223.302C740.365 227.064 739.874 230.171 738.566 232.461C736.276 236.386 731.697 237.858 727.608 239.167C721.393 241.293 715.015 242.928 708.473 244.4Z" fill="white"></path><path d="M215.692 216.923C212.421 204.657 218.145 190.428 229.594 182.414C239.571 175.381 253.8 173.092 267.374 176.526C272.281 177.671 275.715 179.307 278.169 181.596C280.785 183.886 282.748 186.994 284.22 191.246C286.837 198.442 287.327 206.456 285.692 213.979L284.711 213.816C286.183 206.62 285.692 198.769 283.239 191.736C281.767 187.648 279.804 184.54 277.514 182.414C275.061 180.288 271.79 178.652 267.047 177.508C253.799 174.237 240.061 176.363 230.248 183.232C219.29 190.919 213.73 204.657 216.837 216.596L215.692 216.923Z" fill="rgb(var(--primary))"></path><path d="M328.709 256.666L327.892 256.175L330.999 249.797L331.163 249.96C332.471 246.035 332.962 241.619 332.307 237.203C331.326 230.334 327.891 224.119 322.821 220.357C317.751 216.596 310.882 214.96 304.013 215.941C297.144 216.923 290.929 220.357 287.167 225.427L286.35 224.773C290.275 219.54 296.817 215.778 303.849 214.797C311.046 213.815 318.242 215.451 323.476 219.54C328.709 223.465 332.471 230.007 333.452 237.04C334.433 244.072 332.635 251.268 328.709 256.666Z" fill="rgb(var(--primary))"></path><path d="M101.373 216.924C101.046 213.325 103.335 210.218 106.116 207.438C110.041 203.349 114.13 199.424 118.382 195.662C122.634 191.9 127.377 188.302 132.938 186.176C141.606 182.905 151.419 183.886 160.742 185.685C166.303 186.667 172.845 187.975 178.569 190.919C184.784 194.19 189.2 198.933 190.999 204.494C191.326 205.475 191.653 205.966 192.144 205.966L191.98 205.638H192.144C192.307 205.638 192.307 205.639 192.307 205.475L191.326 204.984C191.653 204.494 191.98 204.494 192.307 204.494C192.798 204.657 193.125 205.311 192.961 205.966C192.961 206.456 192.471 206.783 191.98 206.783C190.672 206.783 190.181 205.148 189.854 204.494C188.218 199.26 183.966 194.681 177.915 191.573C172.354 188.629 165.976 187.321 160.415 186.34C151.256 184.704 141.606 183.559 133.102 186.83C127.705 188.793 123.125 192.554 118.873 196.153C114.621 199.751 110.532 203.676 106.607 207.765C104.153 210.381 101.864 213.162 102.191 216.433L101.373 216.924Z" fill="rgb(var(--primary))"></path><path d="M218.474 210.872C217.656 203.839 220.273 196.316 225.506 191.41C230.576 186.503 238.263 183.723 248.731 182.741C251.838 182.414 255.927 182.251 259.852 183.069C270.81 185.031 279.641 194.517 281.113 205.638L280.132 205.802C278.824 195.171 270.156 186.176 259.688 184.213C255.763 183.559 251.838 183.723 248.894 183.886C238.59 184.704 231.23 187.484 226.324 192.227C221.417 196.97 218.801 204.167 219.618 210.872H218.474Z" fill="rgb(var(--primary))"></path><path d="M86.4908 255.357C78.4768 255.357 70.4629 253.068 63.9209 248.652C62.1218 247.343 59.6685 245.381 59.1779 242.6C58.8508 240.801 59.505 239.166 59.9956 237.694C66.047 222.156 82.8927 211.853 99.4113 213.488C100.393 194.68 121.654 179.306 140.463 177.18C158.78 175.054 178.243 186.175 186.911 203.675L185.93 204.166C177.589 187.157 158.453 176.362 140.626 178.325C121.981 180.451 101.047 195.661 100.393 214.143V214.633H99.9021C83.547 212.834 66.8647 222.974 60.9769 238.185C60.4863 239.656 59.9956 241.128 60.1592 242.6C60.4863 245.054 62.776 246.853 64.4115 247.998C73.8975 254.54 86.8179 256.339 97.7758 252.904L98.1031 253.885C94.3414 254.867 90.416 255.357 86.4908 255.357Z" fill="rgb(var(--primary))"></path></svg> <svg class="absolute -top-24 -right-64 z-10" width="756" height="213" viewBox="0 0 756 213" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M387.913 103.117C370.754 106.145 349.304 101.351 331.388 100.594C332.65 95.5469 331.893 90.2477 329.117 85.9578C325.079 79.6492 316.752 76.3688 309.182 78.1352C320.537 60.9758 319.78 36.4985 307.163 20.0961C300.35 11.5164 292.527 6.21721 283.947 3.44143C297.574 1.92737 311.705 4.95545 323.565 11.5164C342.491 21.8625 356.118 40.536 362.931 61.2282C362.426 59.4617 380.595 62.4898 382.361 63.2469C389.679 66.7797 393.717 72.079 396.745 79.1446C403.558 94.2852 402.297 100.594 387.913 103.117Z" fill="white"></path><path d="M369.238 86.2107C372.519 85.4537 375.799 85.2013 379.08 84.949C375.799 85.4537 372.519 85.9583 369.238 86.2107Z" fill="white"></path><path d="M731.099 67.5365C730.595 72.0787 725.548 74.6022 721.258 76.3686C690.725 88.2287 657.92 94.5373 625.115 94.7896C628.648 110.687 617.545 127.09 603.414 135.165C589.282 143.24 572.375 145.258 556.225 147.025C558.244 155.857 553.45 164.689 546.132 170.493C538.561 162.165 525.944 158.38 515.346 162.165C526.701 154.847 538.561 147.025 544.87 135.165C551.431 123.304 550.421 106.145 539.318 98.5748C528.467 91.2568 513.327 96.0514 503.233 104.126C502.981 93.2756 493.392 83.6865 482.541 82.4248C471.69 81.1631 460.839 87.2194 454.783 96.5561C448.979 105.893 447.97 117.501 450.493 128.099C439.895 110.94 424.754 96.5561 406.838 86.967C417.184 88.9858 427.278 92.5186 436.614 97.3131C444.437 79.649 456.045 63.2467 471.943 52.3959C473.457 51.3865 474.971 50.3771 476.485 49.3678C487.336 42.5545 498.943 38.0123 512.065 39.5264C523.421 41.0404 534.271 47.0967 541.842 55.6764C548.403 63.2467 550.926 71.8264 553.197 81.4155C553.702 83.1819 554.964 86.21 556.225 87.2193C560.768 80.6584 562.786 72.3311 571.366 69.5553C579.189 67.0318 587.768 68.2936 595.843 69.0506C602.909 69.5553 609.47 69.0506 616.536 68.5459C648.836 66.0225 680.631 60.2186 712.679 57.4428C721.258 57.4428 732.109 59.9662 731.099 67.5365Z" fill="white"></path><path d="M546.383 170.744C543.355 173.015 540.075 175.034 536.542 176.044C524.429 180.081 511.055 178.315 498.438 176.801C387.407 163.174 274.861 176.044 162.821 178.819C139.857 179.324 116.894 179.576 94.4355 175.034C116.137 155.604 145.157 146.267 174.176 142.986C203.196 139.454 232.215 141.472 261.487 141.22C265.02 141.22 268.805 141.22 272.085 139.706C275.871 138.192 278.646 134.911 281.422 131.883C298.582 114.724 325.582 108.415 348.546 115.481C344.256 107.658 348.293 97.3123 355.359 92.0131C359.396 88.985 363.939 87.2185 368.986 86.2091C372.266 85.7044 375.546 85.4521 378.827 84.9474H379.079C379.584 84.9474 380.341 84.9474 380.846 84.9474C389.678 84.6951 398.762 85.4521 407.594 87.2185C425.511 96.5552 440.651 110.939 451.25 128.351C448.726 117.752 449.736 106.144 455.539 96.8076C461.343 87.4709 472.447 81.4146 483.297 82.6763C494.148 83.9381 503.737 93.5271 503.989 104.378C514.083 96.0506 529.224 91.5084 540.075 98.8263C551.178 106.397 552.187 123.556 545.626 135.416C539.065 147.276 527.205 155.099 516.102 162.417C526.196 158.379 538.813 162.417 546.383 170.744Z" fill="rgb(var(--primary-fg))"></path><path d="M434.344 126.332C427.783 128.099 421.475 129.865 415.419 132.893C409.362 135.921 404.063 140.463 401.035 146.52C402.045 138.445 405.83 130.874 412.138 125.828C418.194 120.781 426.522 118.005 434.597 119.014C440.653 119.771 447.466 122.799 450.747 128.351C444.691 125.828 440.905 124.566 434.344 126.332Z" fill="white"></path><path d="M156.509 162.669C152.472 151.566 142.63 142.482 131.022 139.706C119.415 136.678 106.545 139.706 97.4606 147.529C92.9185 151.314 89.3857 156.865 89.6381 162.669C84.5912 159.893 78.2826 160.146 73.2357 162.921C68.1888 165.697 65.413 171.754 65.6654 177.305C60.8709 174.277 55.5716 172.006 50.02 172.763C44.4685 173.52 39.1693 178.314 39.674 183.866C71.7217 183.866 103.517 183.866 135.565 183.866C154.995 183.866 174.426 183.866 193.856 183.866C197.894 183.866 203.95 184.875 207.735 183.866C213.286 182.352 211.015 180.081 206.978 176.296C201.426 170.744 194.108 165.445 186.79 162.921C184.519 162.164 155.5 160.146 156.509 162.669Z" fill="rgb(var(--primary))"></path><path d="M331.386 100.593C330.124 100.593 328.863 100.593 327.601 100.593C308.171 100.593 288.74 101.35 269.562 103.369C227.168 108.163 184.774 119.519 142.381 115.481C136.577 114.977 129.007 111.949 129.764 106.145C130.521 102.107 134.81 100.088 138.596 99.079C148.689 96.3033 159.288 94.7892 169.886 95.0415C177.204 68.2931 211.775 52.9002 236.505 65.5173C230.953 44.5728 242.309 21.1048 260.982 10.254C267.796 6.21649 275.366 3.94541 283.441 2.93604C292.021 5.45947 300.096 11.0111 306.656 19.5907C319.274 35.9931 320.283 60.4704 308.675 77.6298C315.993 75.8634 324.321 79.1439 328.61 85.4525C331.891 90.247 332.648 95.5462 331.386 100.593Z" fill="rgb(var(--primary-fg))"></path><path d="M445.698 94.7883L444.436 93.022C436.361 82.9282 419.959 81.1618 409.865 88.9845L408.855 87.7227C418.949 79.6477 435.352 81.1618 444.184 90.4985C446.455 55.9274 476.736 24.8891 511.307 22.3657C522.158 21.6087 531.999 23.6274 539.065 28.1696C555.215 38.5157 558.495 60.7219 552.944 79.6477C566.57 57.9461 580.702 48.3571 595.842 50.6282C602.151 51.6375 607.702 54.6657 613.002 57.6938C615.273 58.9556 617.796 60.2173 620.067 61.479C644.292 73.0868 674.321 66.7782 696.275 61.9836L696.527 63.4977C674.321 68.2922 644.04 74.6009 619.31 62.7407C616.787 61.479 614.516 60.2173 611.992 58.9556C606.693 55.9274 601.141 52.8993 595.338 51.8899C579.188 49.3664 564.047 61.2266 549.411 88.4797L547.897 87.7227C557.234 68.0399 556.224 40.7868 538.055 29.1789C529.223 23.6274 518.625 22.8704 511.055 23.6274C476.484 26.1509 446.455 57.4414 445.445 92.2649L445.698 94.7883Z" fill="white"></path><path d="M575.907 151.567C567.579 151.567 559 149.548 551.682 145.258L552.439 143.996C568.841 153.333 590.795 151.819 605.683 139.959C620.572 128.351 627.133 107.407 621.833 89.2379L620.572 84.6956L623.347 88.4808C629.151 96.8081 646.058 101.855 656.152 103.369C667.255 104.883 678.611 102.612 686.433 100.846C693.499 99.0793 701.322 96.8082 707.378 91.7613C713.939 86.4621 717.219 78.8917 716.21 72.0785L717.724 71.8262C718.986 79.1441 715.2 87.4715 708.387 93.023C702.078 98.3223 693.751 100.593 686.686 102.36C678.863 104.378 667.255 106.65 655.9 104.883C646.815 103.621 631.675 99.0793 624.357 92.0137C628.142 109.93 621.329 129.865 606.693 141.221C597.861 148.034 587.01 151.567 575.907 151.567Z" fill="white"></path><path d="M391.948 73.8449C391.696 66.2746 387.406 58.4519 380.34 53.4051C373.275 48.3582 364.947 46.3394 357.377 48.3582L355.863 48.8629L356.368 47.3488C358.639 40.0308 356.872 31.1988 351.573 23.8809C347.283 18.077 340.975 13.0301 331.89 8.23552C323.311 3.94568 306.908 -2.36289 293.786 4.95508L293.029 3.69337C306.908 -4.12929 323.815 2.17927 332.647 6.9738C341.732 11.7683 348.293 16.8152 352.835 22.8714C358.134 30.1894 360.153 39.0215 358.134 46.3395C365.704 44.8254 374.032 46.8441 381.097 51.891C388.415 57.1902 392.957 65.2652 393.21 73.5926L391.948 73.8449Z" fill="white"></path><path d="M385.637 143.24L384.123 142.987C385.637 132.641 392.955 123.052 402.544 118.762C412.133 114.472 423.993 115.482 432.825 121.033C433.33 121.538 434.339 121.79 434.592 121.79H434.844C435.097 121.538 435.097 121.538 435.097 121.538L433.583 121.79C433.583 121.033 433.835 120.529 434.339 120.529C435.096 120.276 435.854 120.781 436.358 121.538C436.611 122.043 436.611 122.8 436.106 123.304C434.592 124.566 432.825 123.304 432.068 122.8C423.741 117.248 412.385 116.491 403.301 120.529C393.964 124.314 387.151 133.398 385.637 143.24Z" fill="white"></path><path d="M651.607 75.3589L650.598 74.0972C666.243 58.9565 685.673 48.1057 706.87 42.5542C717.973 39.5261 726.805 39.0214 734.628 41.0401C744.722 43.3112 752.797 50.3768 755.32 58.7042L753.806 59.2089C751.535 51.3862 743.712 44.8253 734.376 42.5542C726.805 40.7878 718.226 41.2925 707.375 44.0683C686.178 49.6198 667 60.4706 651.607 75.3589Z" fill="white"></path><path d="M104.025 104.631L103.268 103.117C123.707 92.2658 148.185 89.7423 170.391 96.0509C166.858 83.1814 172.157 68.0408 183.26 60.4705C194.111 52.9002 210.009 52.9002 220.86 60.4705C219.598 43.0588 223.888 27.1611 233.477 17.0674C244.075 5.7119 264.263 0.665051 277.385 11.2635L276.375 12.5252C263.758 2.43146 244.58 7.22594 234.486 18.0767C225.15 28.1705 220.86 44.3205 222.626 61.9845L222.878 63.751L221.364 62.4893C211.018 54.1619 195.121 53.6572 184.27 61.2275C173.419 68.7978 168.372 83.9385 172.41 96.5557L172.914 98.0697L171.4 97.565C149.194 91.5088 124.717 93.7799 104.025 104.631Z" fill="rgb(var(--primary-fg))"></path><path d="M477.238 212.886L476.985 211.877C472.948 200.521 462.349 190.932 448.47 186.39C435.853 182.1 421.974 181.848 408.852 181.848C395.478 181.848 381.852 182.352 368.73 182.605C325.831 183.866 281.671 184.876 239.277 175.287L239.53 173.773C281.671 183.362 325.831 182.352 368.73 181.091C381.852 180.838 395.478 180.334 408.852 180.334C422.227 180.334 436.106 180.586 448.975 184.876C462.854 189.67 473.705 199.259 477.995 210.363C497.93 196.736 525.688 196.484 546.128 209.605L545.37 210.867C525.435 197.998 497.677 198.502 478.247 212.381L477.238 212.886Z" fill="rgb(var(--primary-fg))"></path><path d="M253.414 147.276L252.404 146.015C263.507 135.669 276.125 123.808 291.77 118.257C305.144 113.462 324.322 113.715 335.93 125.575L336.435 120.78L337.949 121.033L337.192 129.612L335.93 128.098C325.079 115.481 305.649 114.976 292.275 119.771C276.882 125.323 264.265 136.93 253.414 147.276Z" fill="rgb(var(--primary-fg))"></path><path d="M531.998 140.463C531.746 140.463 531.493 140.463 531.241 140.463V138.949C537.045 139.201 543.101 136.678 547.896 131.883C552.942 126.836 555.718 120.023 555.466 113.462C554.709 100.593 543.353 88.9848 529.979 87.2184C517.11 85.452 503.483 92.2653 495.913 103.873C494.903 105.387 494.146 106.901 493.137 108.415L491.623 107.658C492.38 106.144 493.389 104.63 494.399 103.116C502.221 91.0036 516.605 83.938 529.979 85.7044C544.363 87.4708 556.223 99.5833 556.98 113.462C557.232 120.528 554.456 127.846 549.157 133.145C544.363 137.94 538.054 140.463 531.998 140.463Z" fill="rgb(var(--primary-fg))"></path><path d="M21.7596 206.83C19.4885 206.83 17.2176 206.578 14.9465 206.073C8.38552 204.559 0.815149 199.26 0.0581178 191.185C-0.44657 185.633 2.32927 180.082 8.63787 174.025C11.666 171.25 15.1988 168.221 19.4886 167.969C20.498 167.969 21.5073 167.969 22.7691 167.969C25.2925 168.221 27.8159 168.221 29.0777 166.455C29.8347 165.446 30.0871 164.184 30.0871 162.67C30.0871 162.417 30.0871 161.913 30.0871 161.66C31.3488 152.324 40.4331 145.763 48.7604 144.249C56.3308 142.987 64.9105 144.501 74.4995 148.791C75.7613 149.296 76.7706 149.8 77.78 149.548C79.0417 149.296 79.7988 148.286 80.8081 147.277L81.0604 147.024C88.8831 138.445 100.743 133.903 112.099 134.912C123.454 135.921 134.305 142.735 140.361 152.576L139.099 153.333C133.296 143.996 122.95 137.435 111.846 136.426C100.996 135.417 89.3878 139.706 82.0698 148.034L81.8176 148.286C80.8082 149.548 79.5464 150.81 77.78 151.062C76.2659 151.314 74.7518 150.557 73.4901 150.053C63.901 145.763 55.8261 144.249 48.5082 145.51C40.9378 146.772 32.3581 152.828 31.0964 161.408C31.0964 161.66 31.0964 161.913 31.0964 162.417C30.844 163.932 30.8441 165.698 29.5824 166.96C27.816 169.231 24.7878 168.978 22.0121 168.978C21.0027 168.978 19.9933 168.978 18.9839 168.978C15.1987 169.231 11.9182 172.007 9.14243 174.782C3.33853 180.586 0.562867 185.885 1.06755 190.68C1.57224 197.493 8.38543 202.54 14.694 204.054C21.5073 205.568 29.0776 204.054 35.6385 202.792L92.1636 191.185L92.4159 192.699L35.891 204.307C31.8535 205.821 26.8065 206.83 21.7596 206.83Z" fill="rgb(var(--primary))"></path><path d="M222.624 178.567L221.109 178.315C222.876 169.987 217.324 161.66 210.763 157.623C204.707 154.09 196.632 152.576 186.286 153.333C175.688 154.09 163.827 159.137 163.323 168.726H161.809C162.313 158.127 174.426 152.576 186.286 151.819C196.884 151.062 205.212 152.576 211.773 156.361C218.586 160.398 224.642 169.483 222.624 178.567Z" fill="rgb(var(--primary))"></path></svg>` : ``} <div class="flex flex-row" data-svelte-h="svelte-1qyjo4z"><img class="rounded-l-box select-none" draggable="false" width="615" height="805"${add_attribute("src", subBG, 0)} alt=""> <img class="absolute bottom-4 left-4 z-[3] select-none" draggable="false"${add_attribute("src", smallBush, 0)} alt=""> <img class="absolute top-[93px] left-[67px] z-[2] select-none" draggable="false"${add_attribute("src", bigBoss, 0)} alt=""> <svg class="absolute top-0 right-0 rounded-r-box" width="844" height="805" viewBox="0 0 844 805" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M98.5 0H843.5V805H0L98.5 0Z" fill="rgb(var(--neutral))"></path><rect x="39.5" y="557.5" width="132" height="34" stroke="rgb(var(--neutral-fg)/.2)" stroke-width="7"></rect><rect x="36.5" y="601.5" width="138" height="40" fill="rgb(var(--neutral-fg)/.2)" stroke="rgb(var(--neutral-fg)/.2)"></rect><path d="M66.1379 722.157L175.162 722.157L175.162 678.119L66.1379 678.119L66.1379 722.157Z" fill="rgb(var(--neutral-fg)/.2)"></path><path d="M54.9919 722.157L60.3438 722.157L60.3438 678.119L54.9919 678.119L54.9919 722.157Z" fill="rgb(var(--neutral-fg)/.2))"></path><path d="M45.4197 722.157L50.7715 722.157L50.7715 678.119L45.4197 678.119L45.4197 722.157Z" fill="rgb(var(--neutral-fg)/.2)"></path><path d="M35.9997 722.157L41.3516 722.157L41.3516 678.119L35.9997 678.119L35.9997 722.157Z" fill="rgb(var(--neutral-fg)/.2)"></path><path d="M167.19 716.244L169.637 716.244L169.637 683.369L167.19 683.369L167.19 716.244Z" fill="#261F4D"></path><path d="M158.425 716.244L160.871 716.244L160.871 683.369L158.425 683.369L158.425 716.244Z" fill="#261F4D"></path><path d="M149.659 716.244L152.105 716.244L152.105 683.369L149.659 683.369L149.659 716.244Z" fill="#261F4D"></path><path d="M141.048 716.244L143.494 716.244L143.494 683.369L141.048 683.369L141.048 716.244Z" fill="#261F4D"></path><path d="M132.284 716.244L134.73 716.244L134.73 683.369L132.284 683.369L132.284 716.244Z" fill="#261F4D"></path><path d="M123.518 716.244L125.965 716.244L125.965 683.369L123.518 683.369L123.518 716.244Z" fill="#261F4D"></path><path d="M114.757 716.244L117.203 716.244L117.203 683.369L114.757 683.369L114.757 716.244Z" fill="#261F4D"></path><path d="M105.991 716.244L108.438 716.244L108.438 683.369L105.991 683.369L105.991 716.244Z" fill="#261F4D"></path><path d="M97.2234 716.244L99.6699 716.244L99.6699 683.369L97.2234 683.369L97.2234 716.244Z" fill="#261F4D"></path><path d="M88.616 716.244L91.0625 716.244L91.0625 683.369L88.616 683.369L88.616 716.244Z" fill="#261F4D"></path><path d="M79.8504 716.244L82.2969 716.244L82.2969 683.369L79.8504 683.369L79.8504 716.244Z" fill="#261F4D"></path><path d="M71.0846 716.244L73.5312 716.244L73.5313 683.369L71.0846 683.369L71.0846 716.244Z" fill="#261F4D"></path><path d="M36.0286 778.342L175.176 778.342L175.176 734.61L36.0286 734.61L36.0286 778.342Z" fill="rgb(var(--neutral-fg)/.2)"></path><path class="animate-pulse" d="M161.101 746.937C163.296 746.937 165.076 745.294 165.076 743.267C165.076 741.24 163.296 739.597 161.101 739.597C158.905 739.597 157.125 741.24 157.125 743.267C157.125 745.294 158.905 746.937 161.101 746.937Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M143.21 746.937C145.406 746.937 147.186 745.294 147.186 743.267C147.186 741.24 145.406 739.597 143.21 739.597C141.014 739.597 139.234 741.24 139.234 743.267C139.234 745.294 141.014 746.937 143.21 746.937Z" fill="#FFCC14"></path><path class="animate-pulse" d="M125.165 746.937C127.361 746.937 129.141 745.294 129.141 743.267C129.141 741.24 127.361 739.597 125.165 739.597C122.969 739.597 121.189 741.24 121.189 743.267C121.189 745.294 122.969 746.937 125.165 746.937Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M107.274 746.937C109.47 746.937 111.25 745.294 111.25 743.267C111.25 741.24 109.47 739.597 107.274 739.597C105.079 739.597 103.299 741.24 103.299 743.267C103.299 745.294 105.079 746.937 107.274 746.937Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M89.3858 746.937C91.5814 746.937 93.3615 745.294 93.3615 743.267C93.3615 741.24 91.5814 739.597 89.3858 739.597C87.1901 739.597 85.4102 741.24 85.4102 743.267C85.4102 745.294 87.1901 746.937 89.3858 746.937Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M71.3408 746.937C73.5365 746.937 75.3164 745.294 75.3164 743.267C75.3164 741.24 73.5365 739.597 71.3408 739.597C69.1451 739.597 67.3652 741.24 67.3652 743.267C67.3652 745.294 69.1451 746.937 71.3408 746.937Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M53.4522 746.937C55.6478 746.937 57.4279 745.294 57.4279 743.267C57.4279 741.24 55.6478 739.597 53.4522 739.597C51.2565 739.597 49.4766 741.24 49.4766 743.267C49.4766 745.294 51.2565 746.937 53.4522 746.937Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M161.101 760.24C163.296 760.24 165.076 758.597 165.076 756.57C165.076 754.544 163.296 752.901 161.101 752.901C158.905 752.901 157.125 754.544 157.125 756.57C157.125 758.597 158.905 760.24 161.101 760.24Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M143.21 760.24C145.406 760.24 147.186 758.597 147.186 756.57C147.186 754.544 145.406 752.901 143.21 752.901C141.014 752.901 139.234 754.544 139.234 756.57C139.234 758.597 141.014 760.24 143.21 760.24Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M125.165 760.24C127.361 760.24 129.141 758.597 129.141 756.57C129.141 754.544 127.361 752.901 125.165 752.901C122.969 752.901 121.189 754.544 121.189 756.57C121.189 758.597 122.969 760.24 125.165 760.24Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M107.274 760.24C109.47 760.24 111.25 758.597 111.25 756.57C111.25 754.544 109.47 752.901 107.274 752.901C105.079 752.901 103.299 754.544 103.299 756.57C103.299 758.597 105.079 760.24 107.274 760.24Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M89.3858 760.24C91.5814 760.24 93.3615 758.597 93.3615 756.57C93.3615 754.544 91.5814 752.901 89.3858 752.901C87.1901 752.901 85.4102 754.544 85.4102 756.57C85.4102 758.597 87.1901 760.24 89.3858 760.24Z" fill="#FFCC14"></path><path class="animate-pulse" d="M71.3408 760.24C73.5365 760.24 75.3164 758.597 75.3164 756.57C75.3164 754.544 73.5365 752.901 71.3408 752.901C69.1451 752.901 67.3652 754.544 67.3652 756.57C67.3652 758.597 69.1451 760.24 71.3408 760.24Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M53.4522 760.24C55.6478 760.24 57.4279 758.597 57.4279 756.57C57.4279 754.544 55.6478 752.901 53.4522 752.901C51.2565 752.901 49.4766 754.544 49.4766 756.57C49.4766 758.597 51.2565 760.24 53.4522 760.24Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M161.101 773.39C163.296 773.39 165.076 771.747 165.076 769.721C165.076 767.694 163.296 766.051 161.101 766.051C158.905 766.051 157.125 767.694 157.125 769.721C157.125 771.747 158.905 773.39 161.101 773.39Z" fill="#FFCC14"></path><path class="animate-pulse" d="M143.21 773.39C145.406 773.39 147.186 771.747 147.186 769.721C147.186 767.694 145.406 766.051 143.21 766.051C141.014 766.051 139.234 767.694 139.234 769.721C139.234 771.747 141.014 773.39 143.21 773.39Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M125.165 773.39C127.361 773.39 129.141 771.747 129.141 769.721C129.141 767.694 127.361 766.051 125.165 766.051C122.969 766.051 121.189 767.694 121.189 769.721C121.189 771.747 122.969 773.39 125.165 773.39Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M107.274 773.39C109.47 773.39 111.25 771.747 111.25 769.721C111.25 767.694 109.47 766.051 107.274 766.051C105.079 766.051 103.299 767.694 103.299 769.721C103.299 771.747 105.079 773.39 107.274 773.39Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M89.3858 773.39C91.5814 773.39 93.3615 771.747 93.3615 769.721C93.3615 767.694 91.5814 766.051 89.3858 766.051C87.1901 766.051 85.4102 767.694 85.4102 769.721C85.4102 771.747 87.1901 773.39 89.3858 773.39Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M71.3408 773.39C73.5365 773.39 75.3164 771.747 75.3164 769.721C75.3164 767.694 73.5365 766.051 71.3408 766.051C69.1451 766.051 67.3652 767.694 67.3652 769.721C67.3652 771.747 69.1451 773.39 71.3408 773.39Z" fill="#6EC5BD"></path><path class="animate-pulse" d="M53.4522 773.39C55.6478 773.39 57.4279 771.747 57.4279 769.721C57.4279 767.694 55.6478 766.051 53.4522 766.051C51.2565 766.051 49.4766 767.694 49.4766 769.721C49.4766 771.747 51.2565 773.39 53.4522 773.39Z" fill="#FFCC14"></path></svg></div> ${stage == "auth" ? `<div class="flex flex-col gap-3 absolute top-1/2 -translate-y-1/2 right-0 z-10 pl-[279px] pr-[148.5px] text-neutral-content scale-125" style="width: 844px;"><span class="font-semibold text-3xl">${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.pages.auth.form.title",
                    source: dictionary,
                    placeholderWidth: "100%",
                    placeholderHeight: "36px"
                  },
                  {},
                  {}
                )}</span> <span class="font-semibold">${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.pages.auth.form.description",
                    source: dictionary,
                    placeholderWidth: "100%"
                  },
                  {},
                  {}
                )}</span> ${validate_component(TextInput, "TextInput").$$render(
                  $$result,
                  {
                    name: "login",
                    flags: REQUIRED$1 | TRANSPARENT$1 | UNDERLINE$1,
                    error: formErrors["login"],
                    style: "--input-shadow-color: var(--neutral-fg);--input-main-color: var(--neutral-fg)",
                    value: $formValues["login"]
                  },
                  {
                    value: ($$value) => {
                      $formValues["login"] = $$value;
                      $$settled = false;
                    }
                  },
                  {
                    label: () => {
                      return `<label for="login">${validate_component(Text, "Text").$$render(
                        $$result,
                        {
                          key: "dashboard.pages.auth.form.inputs.username.label",
                          source: dictionary
                        },
                        {},
                        {}
                      )}</label>`;
                    }
                  }
                )} ${validate_component(TextInput, "TextInput").$$render(
                  $$result,
                  {
                    name: "password",
                    flags: REQUIRED$1 | SECURE | SHOW_PASSWORD_BTN | TRANSPARENT$1 | UNDERLINE$1,
                    error: formErrors["password"],
                    style: "--input-shadow-color: var(--neutral-fg);--input-main-color: var(--neutral-fg)",
                    value: $formValues["password"]
                  },
                  {
                    value: ($$value) => {
                      $formValues["password"] = $$value;
                      $$settled = false;
                    }
                  },
                  {
                    label: () => {
                      return `<label for="password">${validate_component(Text, "Text").$$render(
                        $$result,
                        {
                          key: "dashboard.pages.auth.form.inputs.password.label",
                          source: dictionary
                        },
                        {},
                        {}
                      )}</label>`;
                    }
                  }
                )} ${validate_component(Button, "Button").$$render($$result, { flags: SUBMIT, class: "mt-[85px]" }, {}, {
                  default: () => {
                    return `${validate_component(Text, "Text").$$render(
                      $$result,
                      {
                        key: "dashboard.pages.auth.form.buttons.log_in.text",
                        source: dictionary,
                        placeholderWidth: "50px"
                      },
                      {},
                      {}
                    )}`;
                  }
                })}</div>` : ``} ${stage == "project" ? `<div class="flex flex-col gap-3 absolute top-1/2 -translate-y-1/2 right-0 z-10 pl-[279px] pr-[148.5px] text-neutral-content scale-125" style="width: 844px;"><span class="text-center font-semibold text-3xl">${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.pages.auth.project-select.form.title",
                    source: dictionary,
                    placeholderWidth: "100%",
                    placeholderHeight: "36px"
                  },
                  {},
                  {}
                )}</span> ${validate_component(Select, "Select").$$render(
                  $$result,
                  {
                    name: "project",
                    options: projects,
                    searchText: GetText("dashboard.pages.auth.project-select.form.inputs.select.label", $dictionaryStore),
                    noResultsText: GetText("dashboard.pages.auth.project-select.form.inputs.select.not_found", $dictionaryStore),
                    valueKey: "id",
                    labelKey: "name",
                    error: formErrors["project"],
                    flags: REQUIRED | TRANSPARENT | UNDERLINE,
                    style: "--select-main-color: var(--neutral-fg)",
                    value: $formValues["project"]
                  },
                  {
                    value: ($$value) => {
                      $formValues["project"] = $$value;
                      $$settled = false;
                    }
                  },
                  {}
                )} <div class="grid grid-cols-3 gap-2 mt-[85px]">${validate_component(Button, "Button").$$render($$result, { class: "col-span-2", flags: SUBMIT }, {}, {
                  default: () => {
                    return `${validate_component(Text, "Text").$$render(
                      $$result,
                      {
                        key: "dashboard.pages.auth.project-select.form.buttons.confirm.text",
                        source: dictionary,
                        placeholderWidth: "50px"
                      },
                      {},
                      {}
                    )}`;
                  }
                })} ${validate_component(Button, "Button").$$render(
                  $$result,
                  {
                    palette: SMOKE,
                    OnClick: () => window.location.replace("/logout")
                  },
                  {},
                  {
                    default: () => {
                      return `${validate_component(Text, "Text").$$render(
                        $$result,
                        {
                          key: "dashboard.pages.auth.project-select.form.buttons.logout.text",
                          source: dictionary,
                          placeholderWidth: "50px"
                        },
                        {},
                        {}
                      )}`;
                    }
                  }
                )}</div></div>` : ``}`;
              }
            }
          )} ${validate_component(Card, "Card").$$render(
            $$result,
            {
              class: "m-auto !bg-neutral !text-neutral-content lg:hidden"
            },
            {},
            {
              footer: () => {
                return `${stage == "auth" ? `${validate_component(Button, "Button").$$render($$result, { flags: SUBMIT, class: "mt-[85px]" }, {}, {
                  default: () => {
                    return `${validate_component(Text, "Text").$$render(
                      $$result,
                      {
                        key: "dashboard.pages.auth.form.buttons.log_in.text",
                        source: dictionary,
                        placeholderWidth: "50px"
                      },
                      {},
                      {}
                    )}`;
                  }
                })}` : ``} ${stage == "project" ? `<div class="grid grid-cols-3 gap-2 mt-[85px]">${validate_component(Button, "Button").$$render($$result, { class: "col-span-2", flags: SUBMIT }, {}, {
                  default: () => {
                    return `${validate_component(Text, "Text").$$render(
                      $$result,
                      {
                        key: "dashboard.pages.auth.project-select.form.buttons.confirm.text",
                        source: dictionary,
                        placeholderWidth: "50px"
                      },
                      {},
                      {}
                    )}`;
                  }
                })} ${validate_component(Button, "Button").$$render(
                  $$result,
                  {
                    palette: SMOKE,
                    OnClick: () => window.location.replace("/logout")
                  },
                  {},
                  {
                    default: () => {
                      return `${validate_component(Text, "Text").$$render(
                        $$result,
                        {
                          key: "dashboard.pages.auth.project-select.form.buttons.logout.text",
                          source: dictionary,
                          placeholderWidth: "50px"
                        },
                        {},
                        {}
                      )}`;
                    }
                  }
                )}</div>` : ``} `;
              },
              header: () => {
                return `${stage == "auth" ? `<span class="font-semibold text-3xl">${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.pages.auth.form.title",
                    source: dictionary,
                    placeholderWidth: "100%",
                    placeholderHeight: "36px"
                  },
                  {},
                  {}
                )}</span> <span class="font-semibold">${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.pages.auth.form.description",
                    source: dictionary,
                    placeholderWidth: "100%"
                  },
                  {},
                  {}
                )}</span>` : ``} ${stage == "project" ? `<span class="text-center font-semibold text-3xl mt-20">${validate_component(Text, "Text").$$render(
                  $$result,
                  {
                    key: "dashboard.pages.auth.project-select.form.title",
                    source: dictionary,
                    placeholderWidth: "100%",
                    placeholderHeight: "36px"
                  },
                  {},
                  {}
                )}</span>` : ``} `;
              },
              default: () => {
                return `${showClouds ? `<svg class="absolute -bottom-28 -left-60 z-10 hidden sm:block scale-50" width="740" height="290" viewBox="0 0 740 290" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M93.6837 189.447C81.908 186.013 68.3333 187.975 58.3567 195.335C48.3801 202.695 42.4923 215.452 44.1278 227.718C35.9502 226.41 27.4456 225.265 19.4317 226.9C11.2541 228.536 3.40364 233.769 0.786826 241.62C0.459724 242.765 0.132603 243.91 0.459705 244.891C1.11391 247.344 4.22141 248.162 6.83823 248.489C39.3848 252.251 79.9454 253.886 111.511 242.928C139.805 232.952 111.838 194.681 93.6837 189.447Z" fill="rgb(var(--primary-fg))"></path><path d="M696.043 248.979C703.73 248.816 711.908 248.489 718.613 244.563C725.319 240.638 729.898 231.97 726.627 224.937C723.52 218.395 715.179 216.269 707.982 215.451C695.389 214.307 682.469 215.124 670.202 218.068C671.511 207.765 663.333 197.297 653.193 196.152C642.889 194.844 632.586 203.185 631.441 213.325C634.221 205.638 631.768 196.48 626.207 190.428C620.646 184.377 612.632 180.942 604.618 179.47C596.277 177.998 586.791 179.143 581.394 185.685C578.614 189.12 577.469 193.372 576.161 197.624C568.147 222.975 543.941 251.76 586.628 251.596C622.936 251.596 659.572 249.961 696.043 248.979Z" fill="white"></path><path d="M596.275 221.176C589.733 210.381 570.598 208.255 563.074 219.213C568.308 203.512 571.088 186.34 566.345 170.475C561.766 154.611 548.191 140.545 531.673 139.564C515.154 138.419 498.472 154.12 501.743 170.312C487.35 154.938 465.107 147.578 444.5 151.34C474.43 134.003 491.112 96.3866 484.079 62.5316C479.827 42.2513 465.107 21.1532 446.299 8.39624C457.748 9.05044 468.869 11.8308 479.173 16.9009C497.654 25.8962 512.374 42.5784 519.079 61.8774C522.023 70.2185 525.948 83.7932 522.677 92.6249C519.897 100.639 511.556 104.237 509.103 112.905C520.878 98.8398 540.504 93.7698 558.004 98.8399C575.504 103.91 589.733 117.648 598.401 133.676C613.284 160.498 613.448 195.989 596.275 221.176Z" fill="white"></path><path d="M601.019 242.601C596.112 245.708 590.224 246.526 584.337 247.344C552.117 251.596 518.589 255.685 487.351 246.363C479.991 244.236 472.468 241.292 465.108 243.255C461.183 244.236 457.912 246.526 454.477 248.652C413.917 273.512 359.945 274.657 318.24 251.76C311.043 247.834 304.01 243.092 295.833 241.456C288.146 239.984 280.296 241.62 272.609 243.092C216.838 253.232 159.432 253.395 103.661 243.582C101.371 243.255 99.0814 242.764 97.4459 241.292C91.7217 236.386 97.7731 227.391 103.824 223.138C127.866 206.293 158.287 198.769 187.235 202.694C183.964 185.849 194.759 168.676 209.642 160.008C224.361 151.339 242.352 149.704 259.361 150.358C258.871 134.657 267.048 118.793 280.296 110.125C293.543 101.456 311.207 100.148 325.436 106.854C334.922 74.4706 354.221 44.0501 382.515 25.7325C401.16 13.6297 424.057 7.25121 446.136 8.55962C465.108 21.3166 479.664 42.4146 483.916 62.6949C491.113 96.55 474.267 134.167 444.337 151.503C464.945 147.741 487.351 155.101 501.58 170.475C498.309 154.283 514.991 138.746 531.51 139.727C548.028 140.872 561.603 154.938 566.182 170.638C570.925 186.503 567.982 203.676 562.911 219.377C595.294 186.339 599.22 223.465 600.037 222.321C603.145 222.811 606.252 224.774 607.234 227.881C609.36 233.278 605.925 239.493 601.019 242.601Z" fill="rgb(var(--primary-fg))"></path><path d="M251.349 186.34C268.685 186.012 285.367 199.424 288.802 216.433C300.087 213.98 312.68 217.741 320.694 226.246C328.708 234.751 331.979 247.18 329.036 258.465C252.33 247.508 174.153 270.241 97.1204 262.064C86.8167 260.919 74.5505 257.484 72.2608 247.508C70.6253 240.311 75.0411 233.279 79.9476 227.718C98.1017 207.111 123.943 195.662 151.092 193.863C179.223 192.064 192.961 201.713 215.204 217.414C219.947 200.732 233.031 186.667 251.349 186.34Z" fill="rgb(var(--primary))"></path><path d="M632.912 277.928C629.477 282.998 624.571 284.633 619.828 285.614C617.047 286.105 614.267 286.432 611.323 286.759C611.159 286.269 610.996 285.941 610.996 285.451C606.417 270.241 592.678 256.993 576.814 256.993C566.183 256.993 556.534 262.717 547.538 268.605C545.576 269.913 543.122 271.058 541.323 269.75C540.342 268.932 540.015 267.787 539.524 266.479C538.216 262.717 535.599 260.1 532.328 258.301C537.398 257.32 542.468 257.647 547.538 257.32C548.847 257.32 550.319 256.993 551.3 255.685C552.281 254.213 552.445 251.923 552.772 249.96C553.917 243.418 557.842 238.512 562.094 236.386C566.347 234.259 570.926 234.259 575.505 234.423C587.117 234.75 594.314 239.657 603.963 247.998C612.631 255.521 626.206 250.614 633.239 258.792C637.001 263.208 636.346 272.857 632.912 277.928Z" fill="white"></path><path d="M156.488 252.414C146.839 246.199 134.245 245.054 123.778 249.47C116.909 252.414 110.858 257.647 108.404 264.68C105.951 271.713 107.586 280.381 113.474 284.96C118.381 288.885 125.414 289.54 131.792 289.049C151.745 287.577 171.371 277.764 191.161 281.199C206.371 283.815 225.67 275.801 202.119 268.605C193.614 265.988 184.455 269.423 175.951 266.315C168.264 263.535 163.03 256.666 156.488 252.414Z" fill="rgb(var(--primary-fg))"></path><path d="M611.489 286.596C604.947 287.087 598.241 286.923 591.699 286.759C569.129 286.105 546.559 285.615 523.989 284.96C520.228 284.797 515.485 283.325 514.994 277.764C514.83 275.474 515.485 273.185 516.302 271.386C519.41 264.189 525.134 260.101 530.695 258.629C531.349 258.465 532.003 258.302 532.494 258.138C535.765 259.937 538.218 262.554 539.69 266.316C540.181 267.46 540.508 268.769 541.489 269.587C543.288 270.895 545.905 269.75 547.704 268.442C556.536 262.554 566.349 256.83 576.98 256.83C592.844 256.83 606.582 270.077 611.162 285.287C611.325 285.778 611.325 286.269 611.489 286.596Z" fill="rgb(var(--primary-fg))"></path><path d="M408.686 172.602C412.611 160.826 424.714 152.321 436.98 152.485C424.551 159.845 415.392 172.275 411.957 186.34C405.088 184.705 402.144 176.854 397.401 171.62C389.878 163.279 378.429 162.135 368.289 160.172C382.027 152.158 404.434 153.957 408.686 172.602Z" fill="white"></path><path d="M213.569 235.732C200.322 224.12 181.677 219.213 164.504 222.975C169.083 216.433 178.242 214.143 186.093 215.615C193.943 217.087 200.976 221.503 207.845 225.755C205.392 224.284 212.751 212.017 214.06 211.036C217.167 217.251 213.569 228.699 213.569 235.732Z" fill="rgb(var(--primary-fg))"></path><path d="M606.909 177.671H605.928C606.091 158.7 602.82 142.017 596.115 128.279C588.264 112.088 575.344 99.4942 560.624 93.9335C544.105 87.7185 525.461 90.3353 513.031 100.803C507.634 105.382 503.708 111.106 501.419 117.485L500.438 117.158C502.727 110.616 506.816 104.728 512.377 99.9848C524.97 89.1905 544.106 86.4101 560.951 92.7886C575.998 98.5129 589.082 111.106 597.096 127.625C603.638 141.69 606.909 158.372 606.909 177.671Z" fill="white"></path><path d="M599.054 182.414L599.545 176.526H600.526L600.362 178.652C609.848 164.096 612.629 144.961 607.722 128.442C602.488 111.269 588.914 96.3864 572.231 89.6808C564.545 86.5733 555.713 84.9378 546.063 84.7743C534.288 84.6107 522.839 86.2462 511.718 90.1714L511.391 89.1901C522.676 85.4284 534.288 83.6294 546.063 83.7929C555.876 83.9565 564.872 85.7555 572.559 88.6995C589.568 95.405 603.47 110.615 608.703 128.115C613.937 145.615 610.666 165.895 600.199 180.779L599.054 182.414Z" fill="white"></path><path d="M716.488 236.059C715.507 236.059 714.525 236.059 713.544 235.896V234.914C718.123 235.242 722.539 234.751 726.955 233.77C728.1 233.442 729.245 233.115 730.226 232.461C731.208 231.643 731.862 230.335 732.025 228.699C732.516 224.774 730.553 220.685 726.955 217.578C723.848 214.961 720.086 213.162 715.507 211.527C709.946 209.564 703.24 207.438 696.371 207.601C689.175 207.765 682.796 210.382 678.871 214.798L678.054 215.779V214.471C677.563 206.457 675.11 198.77 670.857 192.064C670.039 190.756 669.222 189.611 667.913 188.793C666.278 187.812 664.315 187.648 662.68 187.648C647.142 187.812 632.75 199.915 629.97 215.125L628.988 214.961C631.932 199.26 646.815 186.831 662.843 186.667C664.642 186.667 666.768 186.831 668.731 187.975C670.203 188.793 671.021 190.265 672.002 191.574C676.254 198.116 678.708 205.639 679.362 213.326C683.614 209.237 689.829 206.784 696.862 206.62C703.895 206.457 710.764 208.583 716.488 210.545C721.231 212.181 725.156 213.98 728.264 216.76C732.025 220.031 734.151 224.611 733.661 228.863C733.497 230.171 733.007 232.134 731.535 233.279C730.39 234.26 729.081 234.587 727.937 234.914C723.684 235.569 720.086 236.059 716.488 236.059Z" fill="white"></path><path d="M542.794 259.283L541.812 259.119C543.121 243.909 555.06 230.334 569.943 227.063C573.541 226.246 576.485 226.246 579.102 226.9C582.21 227.718 584.663 229.517 586.625 231.152C592.023 235.405 596.602 240.802 599.873 246.853L600.364 246.69L601.018 249.143L600.037 249.47C596.602 242.764 591.696 236.713 585.808 231.97C584.009 230.498 581.555 228.699 578.611 228.045C576.158 227.391 573.378 227.391 569.943 228.208C555.714 231.316 544.102 244.4 542.794 259.283Z" fill="white"></path><path d="M413.266 174.4H412.285C412.775 166.059 408.196 157.554 401 153.302C393.803 149.049 384.154 149.213 377.121 153.629C373.85 155.592 371.233 158.535 368.617 161.152L367.799 160.498C370.416 157.718 373.196 154.937 376.467 152.811C383.827 148.068 393.967 148.068 401.49 152.484C408.196 156.409 412.612 163.769 413.266 171.619H413.593L413.266 174.4Z" fill="white"></path><path d="M249.874 163.933H248.893C249.71 149.868 250.364 135.312 255.762 121.737C261.813 106.69 273.098 96.0595 286.019 93.2791C293.705 91.6436 302.21 93.4427 308.916 98.0221C313.659 101.293 316.93 105.709 318.565 110.616C321.673 86.0829 327.233 66.9475 335.902 50.4289C347.023 29.3308 363.378 13.3028 382.35 5.45239C403.775 -3.5429 428.471 -1.25319 445.317 11.1767L444.663 11.9944C427.981 -0.108337 403.612 -2.39804 382.677 6.4337C364.032 14.2841 347.677 29.985 336.719 50.9195C327.888 67.9288 322.327 87.7184 319.219 113.559L318.729 117.321L318.238 113.559C317.42 107.835 313.822 102.438 308.425 98.8399C301.883 94.424 293.706 92.6249 286.346 94.2605C273.752 97.0408 262.794 107.345 256.907 122.064C252.327 133.513 251.182 145.779 250.365 158.045L249.874 163.933Z" fill="rgb(var(--primary-fg))"></path><path d="M9.12956 242.764C7.82115 242.764 6.51271 242.601 5.2043 242.11C2.75104 241.292 0.951981 239.493 0.297778 237.204C-0.683527 233.932 0.951961 230.498 2.42392 228.208C9.1295 217.905 23.522 213.816 34.6434 218.886C32.5173 209.727 34.3163 199.423 39.55 190.592C45.1107 181.269 53.6153 174.727 63.2648 172.437C72.7508 170.148 83.5452 172.11 92.5405 177.835C101.699 183.559 108.078 192.391 110.04 201.877L109.059 202.04C107.096 192.718 100.881 184.213 92.0497 178.652C83.218 173.092 72.7508 171.129 63.592 173.419C54.2696 175.708 45.9284 182.087 40.5312 191.082C35.1341 200.077 33.4986 210.545 35.9518 219.704L36.2789 220.848L35.2977 220.358C24.5033 214.797 9.94731 218.559 3.40528 228.699C2.09688 230.825 0.624899 233.932 1.44265 236.876C2.09686 239.33 4.22298 240.638 5.69493 241.129C8.96595 242.274 13.3818 241.619 16.3257 239.493L16.98 240.311C14.6903 241.947 11.9099 242.764 9.12956 242.764Z" fill="rgb(var(--primary-fg))"></path><path d="M205.229 275.802L204.902 271.55L205.883 271.386V272.204C206.21 270.732 205.72 269.097 204.738 267.788C203.593 266.153 201.794 265.008 199.014 264.027C192.472 261.737 185.112 261.737 178.57 264.354L177.916 264.517L177.752 263.863C176.771 257.975 172.519 252.578 165.486 247.508C159.762 243.583 150.93 239.003 140.626 239.167C131.14 239.33 122.636 243.91 118.22 250.942C115.93 254.704 114.622 259.284 114.949 263.863H113.968C113.64 259.12 114.949 254.377 117.402 250.288C121.981 242.928 130.977 238.185 140.626 238.022C151.094 237.858 160.252 242.601 166.14 246.527C173.173 251.433 177.425 256.994 178.734 262.882C185.276 260.428 192.799 260.428 199.341 262.718C202.285 263.699 204.248 265.008 205.556 266.807C207.028 268.77 207.519 271.877 205.883 274.167L205.229 275.802Z" fill="rgb(var(--primary-fg))"></path><path d="M270.319 238.349L269.665 238.022C261.487 232.788 250.693 231.643 241.534 234.914L241.207 233.933C250.366 230.498 261.16 231.643 269.501 236.713C273.099 223.302 285.693 212.835 299.595 212.017C300.576 212.017 301.557 211.854 301.721 211.199V211.036C301.557 210.872 301.23 210.709 301.23 210.709H300.249C300.249 210.218 300.412 209.891 300.739 209.727C301.23 209.564 301.884 209.727 302.375 210.218C302.866 210.545 302.866 211.199 302.702 211.526C302.211 212.835 300.412 212.998 299.595 212.998C285.856 213.98 273.426 224.283 270.155 237.694L270.319 238.349Z" fill="rgb(var(--primary-fg))"></path><path d="M590.225 224.284L588.099 217.087L589.081 216.76L590.062 220.195C591.207 215.452 592.025 210.709 591.207 205.966C590.389 201.223 587.118 195.662 581.721 194.354C575.342 192.718 568.964 197.625 564.711 201.387L562.095 203.676L563.894 200.732C573.38 184.868 574.525 164.588 566.674 147.742C562.422 138.583 556.207 132.041 549.011 129.261C544.268 127.461 538.38 126.971 531.674 127.789C520.226 129.424 510.413 134.331 503.707 141.69C496.347 150.032 493.567 161.48 496.511 171.13L495.53 171.457C492.422 161.48 495.366 149.541 502.889 141.036C509.595 133.513 519.735 128.443 531.347 126.807C538.216 125.826 544.268 126.317 549.174 128.279C556.534 131.223 563.076 137.929 567.328 147.251C574.852 163.443 574.198 182.905 565.856 198.606C570.109 195.172 575.669 191.737 581.557 193.209C587.445 194.681 590.88 200.405 591.861 205.639C592.842 211.2 591.697 216.76 590.389 222.157L590.225 224.284Z" fill="rgb(var(--primary-fg))"></path><path d="M240.881 164.751C234.993 147.414 235.32 129.751 241.698 115.195C248.731 98.8397 263.778 87.064 279.969 85.265C297.469 83.3024 315.296 93.7696 320.694 109.307L319.712 109.634C314.479 94.5874 297.142 84.2837 280.133 86.2463C264.268 88.0454 249.549 99.6575 242.68 115.685C236.465 130.078 236.138 147.414 241.862 164.587L240.881 164.751Z" fill="rgb(var(--primary-fg))"></path><path d="M344.412 193.536L342.94 191.41C338.197 184.868 330.51 180.779 322.332 180.289C314.155 179.961 306.141 183.232 300.744 189.284L299.926 188.63C305.487 182.251 313.828 178.817 322.332 179.144C330.019 179.471 337.379 183.233 342.449 189.12C341.958 183.887 344.575 177.999 349.482 172.602C359.622 161.971 374.832 156.737 389.225 159.027L389.061 160.008C374.832 157.882 359.949 162.952 350.136 173.419C347.519 176.2 341.795 183.396 343.594 191.083L344.412 193.536Z" fill="white"></path><path d="M708.473 244.4L708.309 243.419C714.687 241.947 721.229 240.311 727.444 238.349C731.37 237.204 735.786 235.568 737.748 232.134C739.057 230.008 739.384 227.064 738.729 223.629C737.421 215.942 732.351 208.419 724.501 202.531C719.758 199.097 712.561 194.844 704.384 195.171C691.627 195.662 681.814 206.783 673.636 217.087L672.818 216.433C681.159 205.966 690.973 194.681 704.22 194.027C712.561 193.699 720.085 197.952 724.991 201.55C733.169 207.601 738.402 215.288 739.711 223.302C740.365 227.064 739.874 230.171 738.566 232.461C736.276 236.386 731.697 237.858 727.608 239.167C721.393 241.293 715.015 242.928 708.473 244.4Z" fill="white"></path><path d="M215.692 216.923C212.421 204.657 218.145 190.428 229.594 182.414C239.571 175.381 253.8 173.092 267.374 176.526C272.281 177.671 275.715 179.307 278.169 181.596C280.785 183.886 282.748 186.994 284.22 191.246C286.837 198.442 287.327 206.456 285.692 213.979L284.711 213.816C286.183 206.62 285.692 198.769 283.239 191.736C281.767 187.648 279.804 184.54 277.514 182.414C275.061 180.288 271.79 178.652 267.047 177.508C253.799 174.237 240.061 176.363 230.248 183.232C219.29 190.919 213.73 204.657 216.837 216.596L215.692 216.923Z" fill="rgb(var(--primary))"></path><path d="M328.709 256.666L327.892 256.175L330.999 249.797L331.163 249.96C332.471 246.035 332.962 241.619 332.307 237.203C331.326 230.334 327.891 224.119 322.821 220.357C317.751 216.596 310.882 214.96 304.013 215.941C297.144 216.923 290.929 220.357 287.167 225.427L286.35 224.773C290.275 219.54 296.817 215.778 303.849 214.797C311.046 213.815 318.242 215.451 323.476 219.54C328.709 223.465 332.471 230.007 333.452 237.04C334.433 244.072 332.635 251.268 328.709 256.666Z" fill="rgb(var(--primary))"></path><path d="M101.373 216.924C101.046 213.325 103.335 210.218 106.116 207.438C110.041 203.349 114.13 199.424 118.382 195.662C122.634 191.9 127.377 188.302 132.938 186.176C141.606 182.905 151.419 183.886 160.742 185.685C166.303 186.667 172.845 187.975 178.569 190.919C184.784 194.19 189.2 198.933 190.999 204.494C191.326 205.475 191.653 205.966 192.144 205.966L191.98 205.638H192.144C192.307 205.638 192.307 205.639 192.307 205.475L191.326 204.984C191.653 204.494 191.98 204.494 192.307 204.494C192.798 204.657 193.125 205.311 192.961 205.966C192.961 206.456 192.471 206.783 191.98 206.783C190.672 206.783 190.181 205.148 189.854 204.494C188.218 199.26 183.966 194.681 177.915 191.573C172.354 188.629 165.976 187.321 160.415 186.34C151.256 184.704 141.606 183.559 133.102 186.83C127.705 188.793 123.125 192.554 118.873 196.153C114.621 199.751 110.532 203.676 106.607 207.765C104.153 210.381 101.864 213.162 102.191 216.433L101.373 216.924Z" fill="rgb(var(--primary))"></path><path d="M218.474 210.872C217.656 203.839 220.273 196.316 225.506 191.41C230.576 186.503 238.263 183.723 248.731 182.741C251.838 182.414 255.927 182.251 259.852 183.069C270.81 185.031 279.641 194.517 281.113 205.638L280.132 205.802C278.824 195.171 270.156 186.176 259.688 184.213C255.763 183.559 251.838 183.723 248.894 183.886C238.59 184.704 231.23 187.484 226.324 192.227C221.417 196.97 218.801 204.167 219.618 210.872H218.474Z" fill="rgb(var(--primary))"></path><path d="M86.4908 255.357C78.4768 255.357 70.4629 253.068 63.9209 248.652C62.1218 247.343 59.6685 245.381 59.1779 242.6C58.8508 240.801 59.505 239.166 59.9956 237.694C66.047 222.156 82.8927 211.853 99.4113 213.488C100.393 194.68 121.654 179.306 140.463 177.18C158.78 175.054 178.243 186.175 186.911 203.675L185.93 204.166C177.589 187.157 158.453 176.362 140.626 178.325C121.981 180.451 101.047 195.661 100.393 214.143V214.633H99.9021C83.547 212.834 66.8647 222.974 60.9769 238.185C60.4863 239.656 59.9956 241.128 60.1592 242.6C60.4863 245.054 62.776 246.853 64.4115 247.998C73.8975 254.54 86.8179 256.339 97.7758 252.904L98.1031 253.885C94.3414 254.867 90.416 255.357 86.4908 255.357Z" fill="rgb(var(--primary))"></path></svg> <svg class="absolute -top-24 -right-64 z-10 hidden sm:block scale-50" width="756" height="213" viewBox="0 0 756 213" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M387.913 103.117C370.754 106.145 349.304 101.351 331.388 100.594C332.65 95.5469 331.893 90.2477 329.117 85.9578C325.079 79.6492 316.752 76.3688 309.182 78.1352C320.537 60.9758 319.78 36.4985 307.163 20.0961C300.35 11.5164 292.527 6.21721 283.947 3.44143C297.574 1.92737 311.705 4.95545 323.565 11.5164C342.491 21.8625 356.118 40.536 362.931 61.2282C362.426 59.4617 380.595 62.4898 382.361 63.2469C389.679 66.7797 393.717 72.079 396.745 79.1446C403.558 94.2852 402.297 100.594 387.913 103.117Z" fill="white"></path><path d="M369.238 86.2107C372.519 85.4537 375.799 85.2013 379.08 84.949C375.799 85.4537 372.519 85.9583 369.238 86.2107Z" fill="white"></path><path d="M731.099 67.5365C730.595 72.0787 725.548 74.6022 721.258 76.3686C690.725 88.2287 657.92 94.5373 625.115 94.7896C628.648 110.687 617.545 127.09 603.414 135.165C589.282 143.24 572.375 145.258 556.225 147.025C558.244 155.857 553.45 164.689 546.132 170.493C538.561 162.165 525.944 158.38 515.346 162.165C526.701 154.847 538.561 147.025 544.87 135.165C551.431 123.304 550.421 106.145 539.318 98.5748C528.467 91.2568 513.327 96.0514 503.233 104.126C502.981 93.2756 493.392 83.6865 482.541 82.4248C471.69 81.1631 460.839 87.2194 454.783 96.5561C448.979 105.893 447.97 117.501 450.493 128.099C439.895 110.94 424.754 96.5561 406.838 86.967C417.184 88.9858 427.278 92.5186 436.614 97.3131C444.437 79.649 456.045 63.2467 471.943 52.3959C473.457 51.3865 474.971 50.3771 476.485 49.3678C487.336 42.5545 498.943 38.0123 512.065 39.5264C523.421 41.0404 534.271 47.0967 541.842 55.6764C548.403 63.2467 550.926 71.8264 553.197 81.4155C553.702 83.1819 554.964 86.21 556.225 87.2193C560.768 80.6584 562.786 72.3311 571.366 69.5553C579.189 67.0318 587.768 68.2936 595.843 69.0506C602.909 69.5553 609.47 69.0506 616.536 68.5459C648.836 66.0225 680.631 60.2186 712.679 57.4428C721.258 57.4428 732.109 59.9662 731.099 67.5365Z" fill="white"></path><path d="M546.383 170.744C543.355 173.015 540.075 175.034 536.542 176.044C524.429 180.081 511.055 178.315 498.438 176.801C387.407 163.174 274.861 176.044 162.821 178.819C139.857 179.324 116.894 179.576 94.4355 175.034C116.137 155.604 145.157 146.267 174.176 142.986C203.196 139.454 232.215 141.472 261.487 141.22C265.02 141.22 268.805 141.22 272.085 139.706C275.871 138.192 278.646 134.911 281.422 131.883C298.582 114.724 325.582 108.415 348.546 115.481C344.256 107.658 348.293 97.3123 355.359 92.0131C359.396 88.985 363.939 87.2185 368.986 86.2091C372.266 85.7044 375.546 85.4521 378.827 84.9474H379.079C379.584 84.9474 380.341 84.9474 380.846 84.9474C389.678 84.6951 398.762 85.4521 407.594 87.2185C425.511 96.5552 440.651 110.939 451.25 128.351C448.726 117.752 449.736 106.144 455.539 96.8076C461.343 87.4709 472.447 81.4146 483.297 82.6763C494.148 83.9381 503.737 93.5271 503.989 104.378C514.083 96.0506 529.224 91.5084 540.075 98.8263C551.178 106.397 552.187 123.556 545.626 135.416C539.065 147.276 527.205 155.099 516.102 162.417C526.196 158.379 538.813 162.417 546.383 170.744Z" fill="rgb(var(--primary-fg))"></path><path d="M434.344 126.332C427.783 128.099 421.475 129.865 415.419 132.893C409.362 135.921 404.063 140.463 401.035 146.52C402.045 138.445 405.83 130.874 412.138 125.828C418.194 120.781 426.522 118.005 434.597 119.014C440.653 119.771 447.466 122.799 450.747 128.351C444.691 125.828 440.905 124.566 434.344 126.332Z" fill="white"></path><path d="M156.509 162.669C152.472 151.566 142.63 142.482 131.022 139.706C119.415 136.678 106.545 139.706 97.4606 147.529C92.9185 151.314 89.3857 156.865 89.6381 162.669C84.5912 159.893 78.2826 160.146 73.2357 162.921C68.1888 165.697 65.413 171.754 65.6654 177.305C60.8709 174.277 55.5716 172.006 50.02 172.763C44.4685 173.52 39.1693 178.314 39.674 183.866C71.7217 183.866 103.517 183.866 135.565 183.866C154.995 183.866 174.426 183.866 193.856 183.866C197.894 183.866 203.95 184.875 207.735 183.866C213.286 182.352 211.015 180.081 206.978 176.296C201.426 170.744 194.108 165.445 186.79 162.921C184.519 162.164 155.5 160.146 156.509 162.669Z" fill="rgb(var(--primary))"></path><path d="M331.386 100.593C330.124 100.593 328.863 100.593 327.601 100.593C308.171 100.593 288.74 101.35 269.562 103.369C227.168 108.163 184.774 119.519 142.381 115.481C136.577 114.977 129.007 111.949 129.764 106.145C130.521 102.107 134.81 100.088 138.596 99.079C148.689 96.3033 159.288 94.7892 169.886 95.0415C177.204 68.2931 211.775 52.9002 236.505 65.5173C230.953 44.5728 242.309 21.1048 260.982 10.254C267.796 6.21649 275.366 3.94541 283.441 2.93604C292.021 5.45947 300.096 11.0111 306.656 19.5907C319.274 35.9931 320.283 60.4704 308.675 77.6298C315.993 75.8634 324.321 79.1439 328.61 85.4525C331.891 90.247 332.648 95.5462 331.386 100.593Z" fill="rgb(var(--primary-fg))"></path><path d="M445.698 94.7883L444.436 93.022C436.361 82.9282 419.959 81.1618 409.865 88.9845L408.855 87.7227C418.949 79.6477 435.352 81.1618 444.184 90.4985C446.455 55.9274 476.736 24.8891 511.307 22.3657C522.158 21.6087 531.999 23.6274 539.065 28.1696C555.215 38.5157 558.495 60.7219 552.944 79.6477C566.57 57.9461 580.702 48.3571 595.842 50.6282C602.151 51.6375 607.702 54.6657 613.002 57.6938C615.273 58.9556 617.796 60.2173 620.067 61.479C644.292 73.0868 674.321 66.7782 696.275 61.9836L696.527 63.4977C674.321 68.2922 644.04 74.6009 619.31 62.7407C616.787 61.479 614.516 60.2173 611.992 58.9556C606.693 55.9274 601.141 52.8993 595.338 51.8899C579.188 49.3664 564.047 61.2266 549.411 88.4797L547.897 87.7227C557.234 68.0399 556.224 40.7868 538.055 29.1789C529.223 23.6274 518.625 22.8704 511.055 23.6274C476.484 26.1509 446.455 57.4414 445.445 92.2649L445.698 94.7883Z" fill="white"></path><path d="M575.907 151.567C567.579 151.567 559 149.548 551.682 145.258L552.439 143.996C568.841 153.333 590.795 151.819 605.683 139.959C620.572 128.351 627.133 107.407 621.833 89.2379L620.572 84.6956L623.347 88.4808C629.151 96.8081 646.058 101.855 656.152 103.369C667.255 104.883 678.611 102.612 686.433 100.846C693.499 99.0793 701.322 96.8082 707.378 91.7613C713.939 86.4621 717.219 78.8917 716.21 72.0785L717.724 71.8262C718.986 79.1441 715.2 87.4715 708.387 93.023C702.078 98.3223 693.751 100.593 686.686 102.36C678.863 104.378 667.255 106.65 655.9 104.883C646.815 103.621 631.675 99.0793 624.357 92.0137C628.142 109.93 621.329 129.865 606.693 141.221C597.861 148.034 587.01 151.567 575.907 151.567Z" fill="white"></path><path d="M391.948 73.8449C391.696 66.2746 387.406 58.4519 380.34 53.4051C373.275 48.3582 364.947 46.3394 357.377 48.3582L355.863 48.8629L356.368 47.3488C358.639 40.0308 356.872 31.1988 351.573 23.8809C347.283 18.077 340.975 13.0301 331.89 8.23552C323.311 3.94568 306.908 -2.36289 293.786 4.95508L293.029 3.69337C306.908 -4.12929 323.815 2.17927 332.647 6.9738C341.732 11.7683 348.293 16.8152 352.835 22.8714C358.134 30.1894 360.153 39.0215 358.134 46.3395C365.704 44.8254 374.032 46.8441 381.097 51.891C388.415 57.1902 392.957 65.2652 393.21 73.5926L391.948 73.8449Z" fill="white"></path><path d="M385.637 143.24L384.123 142.987C385.637 132.641 392.955 123.052 402.544 118.762C412.133 114.472 423.993 115.482 432.825 121.033C433.33 121.538 434.339 121.79 434.592 121.79H434.844C435.097 121.538 435.097 121.538 435.097 121.538L433.583 121.79C433.583 121.033 433.835 120.529 434.339 120.529C435.096 120.276 435.854 120.781 436.358 121.538C436.611 122.043 436.611 122.8 436.106 123.304C434.592 124.566 432.825 123.304 432.068 122.8C423.741 117.248 412.385 116.491 403.301 120.529C393.964 124.314 387.151 133.398 385.637 143.24Z" fill="white"></path><path d="M651.607 75.3589L650.598 74.0972C666.243 58.9565 685.673 48.1057 706.87 42.5542C717.973 39.5261 726.805 39.0214 734.628 41.0401C744.722 43.3112 752.797 50.3768 755.32 58.7042L753.806 59.2089C751.535 51.3862 743.712 44.8253 734.376 42.5542C726.805 40.7878 718.226 41.2925 707.375 44.0683C686.178 49.6198 667 60.4706 651.607 75.3589Z" fill="white"></path><path d="M104.025 104.631L103.268 103.117C123.707 92.2658 148.185 89.7423 170.391 96.0509C166.858 83.1814 172.157 68.0408 183.26 60.4705C194.111 52.9002 210.009 52.9002 220.86 60.4705C219.598 43.0588 223.888 27.1611 233.477 17.0674C244.075 5.7119 264.263 0.665051 277.385 11.2635L276.375 12.5252C263.758 2.43146 244.58 7.22594 234.486 18.0767C225.15 28.1705 220.86 44.3205 222.626 61.9845L222.878 63.751L221.364 62.4893C211.018 54.1619 195.121 53.6572 184.27 61.2275C173.419 68.7978 168.372 83.9385 172.41 96.5557L172.914 98.0697L171.4 97.565C149.194 91.5088 124.717 93.7799 104.025 104.631Z" fill="rgb(var(--primary-fg))"></path><path d="M477.238 212.886L476.985 211.877C472.948 200.521 462.349 190.932 448.47 186.39C435.853 182.1 421.974 181.848 408.852 181.848C395.478 181.848 381.852 182.352 368.73 182.605C325.831 183.866 281.671 184.876 239.277 175.287L239.53 173.773C281.671 183.362 325.831 182.352 368.73 181.091C381.852 180.838 395.478 180.334 408.852 180.334C422.227 180.334 436.106 180.586 448.975 184.876C462.854 189.67 473.705 199.259 477.995 210.363C497.93 196.736 525.688 196.484 546.128 209.605L545.37 210.867C525.435 197.998 497.677 198.502 478.247 212.381L477.238 212.886Z" fill="rgb(var(--primary-fg))"></path><path d="M253.414 147.276L252.404 146.015C263.507 135.669 276.125 123.808 291.77 118.257C305.144 113.462 324.322 113.715 335.93 125.575L336.435 120.78L337.949 121.033L337.192 129.612L335.93 128.098C325.079 115.481 305.649 114.976 292.275 119.771C276.882 125.323 264.265 136.93 253.414 147.276Z" fill="rgb(var(--primary-fg))"></path><path d="M531.998 140.463C531.746 140.463 531.493 140.463 531.241 140.463V138.949C537.045 139.201 543.101 136.678 547.896 131.883C552.942 126.836 555.718 120.023 555.466 113.462C554.709 100.593 543.353 88.9848 529.979 87.2184C517.11 85.452 503.483 92.2653 495.913 103.873C494.903 105.387 494.146 106.901 493.137 108.415L491.623 107.658C492.38 106.144 493.389 104.63 494.399 103.116C502.221 91.0036 516.605 83.938 529.979 85.7044C544.363 87.4708 556.223 99.5833 556.98 113.462C557.232 120.528 554.456 127.846 549.157 133.145C544.363 137.94 538.054 140.463 531.998 140.463Z" fill="rgb(var(--primary-fg))"></path><path d="M21.7596 206.83C19.4885 206.83 17.2176 206.578 14.9465 206.073C8.38552 204.559 0.815149 199.26 0.0581178 191.185C-0.44657 185.633 2.32927 180.082 8.63787 174.025C11.666 171.25 15.1988 168.221 19.4886 167.969C20.498 167.969 21.5073 167.969 22.7691 167.969C25.2925 168.221 27.8159 168.221 29.0777 166.455C29.8347 165.446 30.0871 164.184 30.0871 162.67C30.0871 162.417 30.0871 161.913 30.0871 161.66C31.3488 152.324 40.4331 145.763 48.7604 144.249C56.3308 142.987 64.9105 144.501 74.4995 148.791C75.7613 149.296 76.7706 149.8 77.78 149.548C79.0417 149.296 79.7988 148.286 80.8081 147.277L81.0604 147.024C88.8831 138.445 100.743 133.903 112.099 134.912C123.454 135.921 134.305 142.735 140.361 152.576L139.099 153.333C133.296 143.996 122.95 137.435 111.846 136.426C100.996 135.417 89.3878 139.706 82.0698 148.034L81.8176 148.286C80.8082 149.548 79.5464 150.81 77.78 151.062C76.2659 151.314 74.7518 150.557 73.4901 150.053C63.901 145.763 55.8261 144.249 48.5082 145.51C40.9378 146.772 32.3581 152.828 31.0964 161.408C31.0964 161.66 31.0964 161.913 31.0964 162.417C30.844 163.932 30.8441 165.698 29.5824 166.96C27.816 169.231 24.7878 168.978 22.0121 168.978C21.0027 168.978 19.9933 168.978 18.9839 168.978C15.1987 169.231 11.9182 172.007 9.14243 174.782C3.33853 180.586 0.562867 185.885 1.06755 190.68C1.57224 197.493 8.38543 202.54 14.694 204.054C21.5073 205.568 29.0776 204.054 35.6385 202.792L92.1636 191.185L92.4159 192.699L35.891 204.307C31.8535 205.821 26.8065 206.83 21.7596 206.83Z" fill="rgb(var(--primary))"></path><path d="M222.624 178.567L221.109 178.315C222.876 169.987 217.324 161.66 210.763 157.623C204.707 154.09 196.632 152.576 186.286 153.333C175.688 154.09 163.827 159.137 163.323 168.726H161.809C162.313 158.127 174.426 152.576 186.286 151.819C196.884 151.062 205.212 152.576 211.773 156.361C218.586 160.398 224.642 169.483 222.624 178.567Z" fill="rgb(var(--primary))"></path></svg>` : ``} ${stage == "auth" ? `<div class="flex flex-col gap-3 w-full max-w-sm mx-auto pt-10">${validate_component(TextInput, "TextInput").$$render(
                  $$result,
                  {
                    name: "login",
                    flags: REQUIRED$1 | TRANSPARENT$1 | UNDERLINE$1,
                    error: formErrors["login"],
                    style: "--input-shadow-color: var(--neutral-fg);--input-main-color: var(--neutral-fg)",
                    value: $formValues["login"]
                  },
                  {
                    value: ($$value) => {
                      $formValues["login"] = $$value;
                      $$settled = false;
                    }
                  },
                  {
                    label: () => {
                      return `<label for="login">${validate_component(Text, "Text").$$render(
                        $$result,
                        {
                          key: "dashboard.pages.auth.form.inputs.username.label",
                          source: dictionary
                        },
                        {},
                        {}
                      )}</label>`;
                    }
                  }
                )} ${validate_component(TextInput, "TextInput").$$render(
                  $$result,
                  {
                    name: "password",
                    flags: REQUIRED$1 | SECURE | SHOW_PASSWORD_BTN | TRANSPARENT$1 | UNDERLINE$1,
                    error: formErrors["password"],
                    style: "--input-shadow-color: var(--neutral-fg);--input-main-color: var(--neutral-fg)",
                    value: $formValues["password"]
                  },
                  {
                    value: ($$value) => {
                      $formValues["password"] = $$value;
                      $$settled = false;
                    }
                  },
                  {
                    label: () => {
                      return `<label for="password">${validate_component(Text, "Text").$$render(
                        $$result,
                        {
                          key: "dashboard.pages.auth.form.inputs.password.label",
                          source: dictionary
                        },
                        {},
                        {}
                      )}</label>`;
                    }
                  }
                )}</div>` : ``} ${stage == "project" ? `<div class="flex flex-col gap-3 w-full max-w-sm mx-auto pt-10">${validate_component(Select, "Select").$$render(
                  $$result,
                  {
                    name: "project",
                    options: projects,
                    searchText: GetText("dashboard.pages.auth.project-select.form.inputs.select.label", $dictionaryStore),
                    noResultsText: GetText("dashboard.pages.auth.project-select.form.inputs.select.not_found", $dictionaryStore),
                    valueKey: "id",
                    labelKey: "name",
                    error: formErrors["project"],
                    flags: REQUIRED | TRANSPARENT | UNDERLINE,
                    value: $formValues["project"]
                  },
                  {
                    value: ($$value) => {
                      $formValues["project"] = $$value;
                      $$settled = false;
                    }
                  },
                  {}
                )}</div>` : ``}`;
              }
            }
          )}</div>`;
        }
      }
    )}</div> </main>`;
  } while (!$$settled);
  $$unsubscribe_formValues();
  return $$rendered;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-C4AzNVWV.js.map
