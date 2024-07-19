import { c as create_ssr_component, b as add_attribute, e as escape, n as noop, a as subscribe } from './ssr-BISMo5iU.js';
import { r as readable } from './index-DyJcXWsV.js';

const GetValue = function(key, data) {
  if (!data) return "";
  if (key.length == 0) return "";
  const chunk = data[key[0]];
  if (key.length == 1) return typeof chunk == "object" ? "{\n" + Object.entries(chunk).map((v) => `${v[0]}: ${v[1]}`).join(";\n") + "\n}" : chunk;
  if (typeof chunk == "string") return "";
  return GetValue(key.slice(1), chunk);
};
const GetText = function(key, source) {
  return GetValue(key.split("."), source);
};
const Text = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $dict, $$unsubscribe_dict = noop, $$subscribe_dict = () => ($$unsubscribe_dict(), $$unsubscribe_dict = subscribe(dict, ($$value) => $dict = $$value), dict);
  let isLoading = false;
  let dict = readable();
  $$subscribe_dict();
  const GetDictionaryRecord = async function() {
    isLoading = true;
    if (!source) return;
    if (source instanceof Promise) {
      $$subscribe_dict(dict = await source ?? readable());
      isLoading = false;
      return;
    }
    $$subscribe_dict(dict = source);
    isLoading = false;
  };
  let { key = "" } = $$props;
  let { placeholder = true } = $$props;
  let { placeholderHeight = `16px` } = $$props;
  let { placeholderWidth = `100px` } = $$props;
  let { source } = $$props;
  if ($$props.key === void 0 && $$bindings.key && key !== void 0) $$bindings.key(key);
  if ($$props.placeholder === void 0 && $$bindings.placeholder && placeholder !== void 0) $$bindings.placeholder(placeholder);
  if ($$props.placeholderHeight === void 0 && $$bindings.placeholderHeight && placeholderHeight !== void 0) $$bindings.placeholderHeight(placeholderHeight);
  if ($$props.placeholderWidth === void 0 && $$bindings.placeholderWidth && placeholderWidth !== void 0) $$bindings.placeholderWidth(placeholderWidth);
  if ($$props.source === void 0 && $$bindings.source && source !== void 0) $$bindings.source(source);
  {
    GetDictionaryRecord();
  }
  $$unsubscribe_dict();
  return `${placeholder && isLoading ? `<div class="rounded-box bg-black/20 animate-pulse duration-500"${add_attribute("style", `width: ${placeholderWidth};height: ${placeholderHeight};`, 0)}></div>` : `${escape(GetValue(key.split("."), $dict))}`}`;
});

export { GetText as G, Text as T };
//# sourceMappingURL=Text-BcVg2DSs.js.map
