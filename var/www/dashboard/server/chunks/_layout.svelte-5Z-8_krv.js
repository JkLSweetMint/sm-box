import { c as create_ssr_component, a as subscribe, s as setContext, v as validate_component, n as noop } from './ssr-BISMo5iU.js';
import { r as readable } from './index-DyJcXWsV.js';
import { R as Root, s as showToast, E as ERROR, C as CHANGE_MAIN_COLOR } from './Root-D8IkUrtT.js';
import { G as GetText } from './Text-BcVg2DSs.js';
import './index2-C4fvnV3D.js';

const WITH_CREDENTIALS = 1;
const GetServiceURL = (data) => `/${data.name}/api/v${data.version ?? "1.0"}`;
const Fetch = async function(data) {
  data.flags = data.flags ? data.flags : 0;
  try {
    let fetchData = {
      method: data.method ?? "GET",
      credentials: data.flags & WITH_CREDENTIALS ? "include" : "omit",
      headers: data.headers
    };
    if (fetchData.method != "GET") fetchData.body = JSON.stringify(data.data);
    let query = "";
    if (data.query) query = "?" + Object.entries(data.query).map((v) => Array.isArray(v[1]) ? v[1].map((v1) => `${v[0]}=${v1}`) : [`${v[0]}=${v[1]}`]).map((v) => v.join("&")).join("&");
    const response = await fetch(data.url + query, fetchData);
    const headers = Object.fromEntries([...response.headers.entries()]);
    const res = await response.json();
    return {
      data: res,
      meta: {
        headers
      }
    };
  } catch (err) {
    throw new Error(`Failed to call a request: ${err}`);
  }
};
const Layout = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $services, $$unsubscribe_services;
  let $dictionary, $$unsubscribe_dictionary = noop, $$subscribe_dictionary = () => ($$unsubscribe_dictionary(), $$unsubscribe_dictionary = subscribe(dictionary, ($$value) => $dictionary = $$value), dictionary);
  let dictionary = readable();
  $$subscribe_dictionary();
  const services = readable({
    "auth": GetServiceURL({ name: "authentication" }),
    "i18n": GetServiceURL({ name: "i18n" })
  });
  $$unsubscribe_services = subscribe(services, (value) => $services = value);
  setContext("services", services);
  const callServiceMethod = async function(data) {
    const errorTitle = GetText("dashboard.toasts.error.title", $dictionary) || "An error occured";
    try {
      const service = $services[data.service];
      if (!service) throw new Error(`Failed to call service method: service "${service}" is not defined`);
      let resp = await Fetch({
        url: `${service}${data.url}`,
        method: data.method,
        query: data.query,
        headers: data.headers,
        data: data.data,
        flags: WITH_CREDENTIALS
      });
      if (resp.data.code != 200 && data.showToast) {
        showToast({
          data: {
            title: errorTitle,
            description: `${resp.data.error?.message}`,
            type: ERROR,
            flags: CHANGE_MAIN_COLOR
          }
        });
      }
      resp.data.headers = resp.meta.headers;
      return resp.data;
    } catch (error) {
      if (data.showToast) {
        showToast({
          data: {
            title: errorTitle,
            description: `${error}`,
            type: ERROR,
            flags: CHANGE_MAIN_COLOR
          }
        });
      }
      return {
        code: 500,
        code_message: "",
        status: "fatal"
      };
    }
  };
  setContext("CallServiceMethod", callServiceMethod);
  const i18nLoadDictionary = async function(path) {
    let resp = await callServiceMethod({
      service: "i18n",
      url: "/texts/dictionary",
      query: { paths: path },
      showToast: true
    });
    if (!resp || resp.code != 200) return;
    return readable(resp.data?.dictionary ?? {});
  };
  setContext("i18nLoadDictionary", i18nLoadDictionary);
  $$unsubscribe_services();
  $$unsubscribe_dictionary();
  return `${validate_component(Root, "Toast").$$render($$result, {}, {}, {})} ${slots.default ? slots.default({}) : ``}`;
});

export { Layout as default };
//# sourceMappingURL=_layout.svelte-5Z-8_krv.js.map
