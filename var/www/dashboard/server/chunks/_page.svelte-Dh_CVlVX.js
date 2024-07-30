import { c as create_ssr_component, g as getContext, e as escape, v as validate_component } from './ssr-C-9IsUTH.js';
import { T as Text } from './Text-DvReL212.js';
import './index2-DbZx0BBT.js';

const css = {
  code: "@keyframes svelte-1fg87d-tilt-n-move-shaking{0%{transform:translate(0, 0) rotate(0deg)}25%{transform:translate(5px, 5px) rotate(5deg)}50%{transform:translate(0, 0) rotate(0eg)}75%{transform:translate(-5px, 5px) rotate(-5deg)}100%{transform:translate(0, 0) rotate(0deg)}}.monitor.svelte-1fg87d{will-change:transform;animation:svelte-1fg87d-tilt-n-move-shaking .3s infinite}",
  map: '{"version":3,"file":"+page.svelte","sources":["+page.svelte"],"sourcesContent":["<script lang=\\"ts\\">import { getContext, onMount } from \\"svelte\\";\\nimport { Text, GetText } from \\"@/components/i18n/Text\\";\\nimport {} from \\"@/lib/types\\";\\nlet title = \\"\\";\\nconst SetPageTitle = function() {\\n  if (!dictionaryStore) return;\\n  title = GetText(\\"dashboard.pages.errors.403.title\\", $dictionaryStore);\\n};\\nconst i18nLoadDictionary = getContext(\\"i18nLoadDictionary\\");\\nlet dictionary;\\nlet dictionaryStore;\\nonMount(async () => {\\n  dictionary = i18nLoadDictionary([\\"dashboard.pages.errors.403\\"]);\\n  dictionary.then((res) => {\\n    dictionaryStore = res;\\n    SetPageTitle();\\n  });\\n});\\n<\/script>\\r\\n\\r\\n<svelte:head>\\r\\n    <title>{title}</title>\\r\\n</svelte:head>\\r\\n\\r\\n<div class=\\"fs-page\\">\\r\\n    <div class=\\"m-auto\\">\\r\\n        <svg class=\\"monitor mx-auto\\" width=\\"300\\" version=\\"1.1\\" xmlns=\\"http://www.w3.org/2000/svg\\" xmlns:xlink=\\"http://www.w3.org/1999/xlink\\" viewBox=\\"0 0 32 32\\" xml:space=\\"preserve\\" fill=\\"#000000\\">\\r\\n            <g stroke-width=\\"0\\"></g>\\r\\n            <g stroke-linecap=\\"round\\" stroke-linejoin=\\"round\\"></g>\\r\\n                <g>\\r\\n                    <style type=\\"text/css\\"> \\r\\n                        .isometric_een{fill:rgb(var(--neutral-fg));} \\r\\n                        .isometric_twee{fill:rgb(var(--smoke));} \\r\\n                        .isometric_vier{fill:rgb(var(--neutral));} \\r\\n                        .isometric_twaalf{fill:#569080;} \\r\\n                        .isometric_dertien{fill:#225B49;} \\r\\n                        .st0{fill:#7BD6C4;} \\r\\n                        .st1{fill:#F05A28;} \\r\\n                        .st2{fill:#FFBB33;} \\r\\n                        .st3{fill:#BE1E2D;} \\r\\n                        .st4{fill:#F29227;} \\r\\n                        .st5{fill:#FF7344;} \\r\\n                        .st6{fill:#6B9086;} \\r\\n                        .st7{fill:none;} \\r\\n                        .st8{fill:#72C0AB;} \\r\\n                        .st9{fill:#F2D76C;} \\r\\n                        .st10{fill:#F28103;} \\r\\n                        .st11{fill:#225B49;} \\r\\n                    </style>\\r\\n                <g>\\r\\n                    <path class=\\"isometric_vier\\" d=\\"M21.08,14.291c0-0.124,0,11.548,0,11.548L27,21.1v-9.987L21.08,14.291z\\"></path> \\r\\n                    <polygon class=\\"isometric_twee\\" points=\\"21.08,14.167 7.695,6.456 5,7.963 18.4,15.711 18.4,28 21.08,26.453 \\"></polygon>\\r\\n                    <path class=\\"isometric_een\\" d=\\"M8.545,6.975l7.624-2.107L27,11.113l-5.92,3.178c-0.002-0.08,0.003-0.125,0-0.124L8.545,6.975z M5,7.963L5,7.963L5,7.963v12.301L18.4,28V15.711L5,7.963z\\"></path>\\r\\n                    <path class=\\"isometric_dertien\\" d=\\"M6.546,10.791v8.828l9.642,5.567C13.221,22.645,7.166,18.618,6.546,10.791z\\"></path>\\r\\n                    <path class=\\"isometric_twaalf\\" d=\\"M16.853,16.344L6.546,10.393v0.398c0.605,7.644,6.174,11.425,9.642,14.394l0.666,0.384V16.344z\\"></path>\\r\\n                </g>\\r\\n            </g>\\r\\n        </svg>\\r\\n        <p class=\\"text-center font-semibold text-5xl\\">\\r\\n            <Text key=\\"dashboard.pages.errors.403.title\\" placeholderWidth=\\"300px\\" placeholderHeight=\\"48px\\" source={dictionary}/>\\r\\n        </p>\\r\\n        <p class=\\"text-center text-neutral-400 font-semibold text-xl mt-5\\">\\r\\n            <Text key=\\"dashboard.pages.errors.403.description\\" placeholderWidth=\\"400px\\" placeholderHeight=\\"28px\\" source={dictionary}/>\\r\\n        </p>\\r\\n    </div>\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    @keyframes tilt-n-move-shaking {\\r\\n        0% { transform: translate(0, 0) rotate(0deg); }\\r\\n        25% { transform: translate(5px, 5px) rotate(5deg); }\\r\\n        50% { transform: translate(0, 0) rotate(0eg); }\\r\\n        75% { transform: translate(-5px, 5px) rotate(-5deg); }\\r\\n        100% { transform: translate(0, 0) rotate(0deg); }\\r\\n    }\\r\\n\\r\\n    .monitor {\\r\\n        will-change: transform;\\r\\n\\r\\n        animation: tilt-n-move-shaking .3s infinite;\\r\\n    }\\r\\n</style>"],"names":[],"mappings":"AAoEI,WAAW,iCAAoB,CAC3B,EAAG,CAAE,SAAS,CAAE,UAAU,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,OAAO,IAAI,CAAG,CAC9C,GAAI,CAAE,SAAS,CAAE,UAAU,GAAG,CAAC,CAAC,GAAG,CAAC,CAAC,OAAO,IAAI,CAAG,CACnD,GAAI,CAAE,SAAS,CAAE,UAAU,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,OAAO,GAAG,CAAG,CAC9C,GAAI,CAAE,SAAS,CAAE,UAAU,IAAI,CAAC,CAAC,GAAG,CAAC,CAAC,OAAO,KAAK,CAAG,CACrD,IAAK,CAAE,SAAS,CAAE,UAAU,CAAC,CAAC,CAAC,CAAC,CAAC,CAAC,OAAO,IAAI,CAAG,CACpD,CAEA,sBAAS,CACL,WAAW,CAAE,SAAS,CAEtB,SAAS,CAAE,iCAAmB,CAAC,GAAG,CAAC,QACvC"}'
};
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let title = "";
  getContext("i18nLoadDictionary");
  let dictionary;
  $$result.css.add(css);
  return `${$$result.head += `<!-- HEAD_svelte-1uo06u1_START -->${$$result.title = `<title>${escape(title)}</title>`, ""}<!-- HEAD_svelte-1uo06u1_END -->`, ""} <div class="fs-page"><div class="m-auto"><svg class="monitor mx-auto svelte-1fg87d" width="300" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 32 32" xml:space="preserve" fill="#000000"><g stroke-width="0"></g><g stroke-linecap="round" stroke-linejoin="round"></g><g><style type="text/css">.isometric_een{fill:rgb(var(--neutral-fg));} 
                        .isometric_twee{fill:rgb(var(--smoke));} 
                        .isometric_vier{fill:rgb(var(--neutral));} 
                        .isometric_twaalf{fill:#569080;} 
                        .isometric_dertien{fill:#225B49;} 
                        .st0{fill:#7BD6C4;} 
                        .st1{fill:#F05A28;} 
                        .st2{fill:#FFBB33;} 
                        .st3{fill:#BE1E2D;} 
                        .st4{fill:#F29227;} 
                        .st5{fill:#FF7344;} 
                        .st6{fill:#6B9086;} 
                        .st7{fill:none;} 
                        .st8{fill:#72C0AB;} 
                        .st9{fill:#F2D76C;} 
                        .st10{fill:#F28103;} 
                        .st11{fill:#225B49;} 
                    </style><g><path class="isometric_vier" d="M21.08,14.291c0-0.124,0,11.548,0,11.548L27,21.1v-9.987L21.08,14.291z"></path><polygon class="isometric_twee" points="21.08,14.167 7.695,6.456 5,7.963 18.4,15.711 18.4,28 21.08,26.453 "></polygon><path class="isometric_een" d="M8.545,6.975l7.624-2.107L27,11.113l-5.92,3.178c-0.002-0.08,0.003-0.125,0-0.124L8.545,6.975z M5,7.963L5,7.963L5,7.963v12.301L18.4,28V15.711L5,7.963z"></path><path class="isometric_dertien" d="M6.546,10.791v8.828l9.642,5.567C13.221,22.645,7.166,18.618,6.546,10.791z"></path><path class="isometric_twaalf" d="M16.853,16.344L6.546,10.393v0.398c0.605,7.644,6.174,11.425,9.642,14.394l0.666,0.384V16.344z"></path></g></g></svg> <p class="text-center font-semibold text-5xl">${validate_component(Text, "Text").$$render(
    $$result,
    {
      key: "dashboard.pages.errors.403.title",
      placeholderWidth: "300px",
      placeholderHeight: "48px",
      source: dictionary
    },
    {},
    {}
  )}</p> <p class="text-center text-neutral-400 font-semibold text-xl mt-5">${validate_component(Text, "Text").$$render(
    $$result,
    {
      key: "dashboard.pages.errors.403.description",
      placeholderWidth: "400px",
      placeholderHeight: "28px",
      source: dictionary
    },
    {},
    {}
  )}</p></div> </div>`;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-Dh_CVlVX.js.map
