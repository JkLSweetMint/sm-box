import { c as create_ssr_component, b as add_attribute, v as validate_component } from './ssr-C-9IsUTH.js';
import { P as PLAIN, B as Button } from './Button-atP5DRSs.js';
import './index-VQC3TRid.js';

const SomethingLost = "/_app/immutable/assets/something-lost.DDfFuHnv.png";
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `<div class="fs-page text-center"><div class="flex flex-col items-center w-[90%] m-auto p-3"><img width="800"${add_attribute("src", SomethingLost, 0)} alt="something lost"> ${validate_component(Button, "Button").$$render(
    $$result,
    {
      flags: PLAIN,
      OnClick: () => window.history.back(),
      style: "color: rgb(var(--primary))"
    },
    {},
    {
      default: () => {
        return `Back`;
      }
    }
  )}</div></div>`;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-AiBW9M0m.js.map
