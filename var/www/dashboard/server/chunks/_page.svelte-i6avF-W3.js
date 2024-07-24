import { c as create_ssr_component, g as getContext, e as escape, v as validate_component } from './ssr-0mLmEaQb.js';
import { T as Text } from './Text-Ce7PraTT.js';
import './index2-Cn3cpumX.js';

const css = {
  code: "@keyframes svelte-8jvyl7-ghost{40%{transform:translateY(-30px)}80%{transform:translateY(0)}}.ghost.svelte-8jvyl7{animation:svelte-8jvyl7-ghost 2s infinite;will-change:transform}",
  map: '{"version":3,"file":"+page.svelte","sources":["+page.svelte"],"sourcesContent":["<script lang=\\"ts\\">import { getContext, onMount } from \\"svelte\\";\\nimport { Text, GetText } from \\"@/components/i18n/Text\\";\\nimport {} from \\"@/lib/types\\";\\nlet title = \\"\\";\\nconst SetPageTitle = function() {\\n  if (!dictionaryStore) return;\\n  title = GetText(\\"dashboard.pages.errors.50x.title\\", $dictionaryStore);\\n};\\nconst i18nLoadDictionary = getContext(\\"i18nLoadDictionary\\");\\nlet dictionary;\\nlet dictionaryStore;\\nonMount(async () => {\\n  dictionary = i18nLoadDictionary([\\"dashboard.pages.errors.50x\\"]);\\n  dictionary.then((res) => {\\n    dictionaryStore = res;\\n    SetPageTitle();\\n  });\\n});\\n<\/script>\\r\\n\\r\\n<svelte:head>\\r\\n    <title>{title}</title>\\r\\n</svelte:head>\\r\\n\\r\\n<div class=\\"fs-page\\">\\r\\n    <div class=\\"m-auto\\">\\r\\n        <svg width=\\"300\\" class=\\"ghost mx-auto\\" version=\\"1.1\\" xmlns=\\"http://www.w3.org/2000/svg\\" xmlns:xlink=\\"http://www.w3.org/1999/xlink\\" viewBox=\\"0 0 32 32\\" xml:space=\\"preserve\\" fill=\\"#000000\\">\\r\\n            <g stroke-width=\\"0\\"></g>\\r\\n            <g stroke-linecap=\\"round\\" stroke-linejoin=\\"round\\"></g>\\r\\n                <g> \\r\\n                    <style type=\\"text/css\\">\\r\\n                    .isometric_een{fill:rgb(var(--neutral-fg));}\\r\\n                    .isometric_twee{fill:rgb(var(--neutral));}\\r\\n                    .isometric_tien{fill:#7BD6C4;}\\r\\n                    .isometric_twaalf{fill:rgb(var(--neutral));}\\r\\n                    .isometric_dertien{fill:rgb(var(--neutral));}\\r\\n                    .st0{fill:#F05A28;}\\r\\n                    .st1{fill:#FFBB33;}\\r\\n                    .st2{fill:#BE1E2D;}\\r\\n                    .st3{fill:#F29227;}\\r\\n                    .st4{fill:#FF7344;}\\r\\n                    .st5{fill:#6B9086;}\\r\\n                    .st6{fill:none;}\\r\\n                    .st7{fill:#72C0AB;}\\r\\n                    .st8{fill:#AD9A74;}\\r\\n                    .st9{fill:#F2D76C;}\\r\\n                    .st10{fill:#F28103;}\\r\\n                    .st11{fill:#225B49;}\\r\\n                    .st12{fill:#7BD6C4;}\\r\\n                </style>\\r\\n                <g>\\r\\n                    <path class=\\"isometric_twee\\" d=\\"M25.993,14.094c-0.058-6.713-5.267-10.086-10.353-9.99c4.297,1.26,7.884,4.874,7.935,10.781h0.008 c0,10.272,0.072,9.587-0.113,10.552c1.733-0.167,2.349-0.772,2.53-2.686v-8.657H25.993z\\"></path>\\r\\n                    <path class=\\"isometric_een\\" d=\\"M23.575,14.885c-0.051-5.907-3.638-9.522-7.935-10.781C10.764,4.196,6.002,7.477,6,14.094v8.603 c0.225,2.396,1.092,2.803,4,2.803c1.105,0,2-0.605,2,0.5v0.035c0,2.689,8,2.753,8,0.02V26c0-0.886,1.132-0.338,3.47-0.563 c0.186-0.964,0.113-0.28,0.113-10.552H23.575z\\"></path>\\r\\n                    <path class=\\"isometric_tien\\" d=\\"M8.355,14.358c-0.009,0.952,0.701,2.287,1.699,2.638C9.066,16.649,8.349,15.312,8.355,14.358z\\"></path>\\r\\n                    <path class=\\"isometric_tien\\" d=\\"M13.746,16.487c-0.009,0.952,0.701,2.287,1.699,2.638C14.456,18.778,13.739,17.441,13.746,16.487z\\"></path>\\r\\n                    <path class=\\"isometric_dertien\\" d=\\"M13.746,16.487c0.004-0.418,0.132-0.77,0.459-0.919C13.877,15.718,13.749,16.07,13.746,16.487z\\"></path>\\r\\n                    <path class=\\"isometric_twaalf\\" d=\\"M15.445,19.125c-0.988,0.452-2.398-1.28-2.398-2.608c0-0.801,0.498-1.181,1.158-0.949 C13.182,16.036,13.984,18.612,15.445,19.125z M8.814,13.439c-0.66-0.232-1.158,0.148-1.158,0.949c0,1.328,1.411,3.06,2.398,2.608 C8.594,16.483,7.792,13.907,8.814,13.439z\\"></path>\\r\\n                    <path class=\\"isometric_dertien\\" d=\\"M15.445,19.125c-1.449-0.51-2.265-3.088-1.24-3.557C15.626,16.068,16.522,18.632,15.445,19.125z M8.814,13.439c-1.025,0.469-0.209,3.047,1.24,3.557C11.132,16.503,10.235,13.939,8.814,13.439z\\"></path>\\r\\n                    <path class=\\"isometric_dertien\\" d=\\"M8.355,14.358c0.004-0.418,0.132-0.77,0.459-0.919C8.486,13.59,8.358,13.941,8.355,14.358z\\"></path>\\r\\n                </g>\\r\\n            </g>\\r\\n        </svg>\\r\\n        <p class=\\"text-center font-semibold text-5xl\\">\\r\\n            <Text key=\\"dashboard.pages.errors.50x.title\\" placeholderWidth=\\"300px\\" placeholderHeight=\\"48px\\" source={dictionary}/>\\r\\n        </p>\\r\\n        <p class=\\"text-center text-neutral-400 font-semibold text-xl mt-5\\">\\r\\n            <Text key=\\"dashboard.pages.errors.50x.description\\" placeholderWidth=\\"400px\\" placeholderHeight=\\"28px\\" source={dictionary}/>\\r\\n        </p>\\r\\n    </div>\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    @keyframes ghost {\\r\\n        40% {\\r\\n            transform: translateY(-30px);\\r\\n        }\\r\\n        80% {\\r\\n            transform: translateY(0);\\r\\n        }\\r\\n    }\\r\\n\\r\\n    .ghost {\\r\\n        animation: ghost 2s infinite;\\r\\n\\r\\n        will-change: transform;\\r\\n    }\\r\\n</style>"],"names":[],"mappings":"AAwEI,WAAW,mBAAM,CACb,GAAI,CACA,SAAS,CAAE,WAAW,KAAK,CAC/B,CACA,GAAI,CACA,SAAS,CAAE,WAAW,CAAC,CAC3B,CACJ,CAEA,oBAAO,CACH,SAAS,CAAE,mBAAK,CAAC,EAAE,CAAC,QAAQ,CAE5B,WAAW,CAAE,SACjB"}'
};
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let title = "";
  getContext("i18nLoadDictionary");
  let dictionary;
  $$result.css.add(css);
  return `${$$result.head += `<!-- HEAD_svelte-1uo06u1_START -->${$$result.title = `<title>${escape(title)}</title>`, ""}<!-- HEAD_svelte-1uo06u1_END -->`, ""} <div class="fs-page"><div class="m-auto"><svg width="300" class="ghost mx-auto svelte-8jvyl7" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 32 32" xml:space="preserve" fill="#000000"><g stroke-width="0"></g><g stroke-linecap="round" stroke-linejoin="round"></g><g><style type="text/css">.isometric_een{fill:rgb(var(--neutral-fg));}
                    .isometric_twee{fill:rgb(var(--neutral));}
                    .isometric_tien{fill:#7BD6C4;}
                    .isometric_twaalf{fill:rgb(var(--neutral));}
                    .isometric_dertien{fill:rgb(var(--neutral));}
                    .st0{fill:#F05A28;}
                    .st1{fill:#FFBB33;}
                    .st2{fill:#BE1E2D;}
                    .st3{fill:#F29227;}
                    .st4{fill:#FF7344;}
                    .st5{fill:#6B9086;}
                    .st6{fill:none;}
                    .st7{fill:#72C0AB;}
                    .st8{fill:#AD9A74;}
                    .st9{fill:#F2D76C;}
                    .st10{fill:#F28103;}
                    .st11{fill:#225B49;}
                    .st12{fill:#7BD6C4;}
                </style><g><path class="isometric_twee" d="M25.993,14.094c-0.058-6.713-5.267-10.086-10.353-9.99c4.297,1.26,7.884,4.874,7.935,10.781h0.008 c0,10.272,0.072,9.587-0.113,10.552c1.733-0.167,2.349-0.772,2.53-2.686v-8.657H25.993z"></path><path class="isometric_een" d="M23.575,14.885c-0.051-5.907-3.638-9.522-7.935-10.781C10.764,4.196,6.002,7.477,6,14.094v8.603 c0.225,2.396,1.092,2.803,4,2.803c1.105,0,2-0.605,2,0.5v0.035c0,2.689,8,2.753,8,0.02V26c0-0.886,1.132-0.338,3.47-0.563 c0.186-0.964,0.113-0.28,0.113-10.552H23.575z"></path><path class="isometric_tien" d="M8.355,14.358c-0.009,0.952,0.701,2.287,1.699,2.638C9.066,16.649,8.349,15.312,8.355,14.358z"></path><path class="isometric_tien" d="M13.746,16.487c-0.009,0.952,0.701,2.287,1.699,2.638C14.456,18.778,13.739,17.441,13.746,16.487z"></path><path class="isometric_dertien" d="M13.746,16.487c0.004-0.418,0.132-0.77,0.459-0.919C13.877,15.718,13.749,16.07,13.746,16.487z"></path><path class="isometric_twaalf" d="M15.445,19.125c-0.988,0.452-2.398-1.28-2.398-2.608c0-0.801,0.498-1.181,1.158-0.949 C13.182,16.036,13.984,18.612,15.445,19.125z M8.814,13.439c-0.66-0.232-1.158,0.148-1.158,0.949c0,1.328,1.411,3.06,2.398,2.608 C8.594,16.483,7.792,13.907,8.814,13.439z"></path><path class="isometric_dertien" d="M15.445,19.125c-1.449-0.51-2.265-3.088-1.24-3.557C15.626,16.068,16.522,18.632,15.445,19.125z M8.814,13.439c-1.025,0.469-0.209,3.047,1.24,3.557C11.132,16.503,10.235,13.939,8.814,13.439z"></path><path class="isometric_dertien" d="M8.355,14.358c0.004-0.418,0.132-0.77,0.459-0.919C8.486,13.59,8.358,13.941,8.355,14.358z"></path></g></g></svg> <p class="text-center font-semibold text-5xl">${validate_component(Text, "Text").$$render(
    $$result,
    {
      key: "dashboard.pages.errors.50x.title",
      placeholderWidth: "300px",
      placeholderHeight: "48px",
      source: dictionary
    },
    {},
    {}
  )}</p> <p class="text-center text-neutral-400 font-semibold text-xl mt-5">${validate_component(Text, "Text").$$render(
    $$result,
    {
      key: "dashboard.pages.errors.50x.description",
      placeholderWidth: "400px",
      placeholderHeight: "28px",
      source: dictionary
    },
    {},
    {}
  )}</p></div> </div>`;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-i6avF-W3.js.map
