import { c as create_ssr_component, v as validate_component, b as add_attribute, d as each, e as escape } from './ssr-C-9IsUTH.js';
import './Root-DpK35LNX.js';
import './create-CrsTuLZT.js';
import { c as cn } from './index-VQC3TRid.js';
import { N as NO_RIPPLE, a as NO_ANIMATION, P as PLAIN, B as Button } from './Button-atP5DRSs.js';
import { C as COMPACT, a as Card } from './Card-BWlCOFJT.js';
import './index2-DbZx0BBT.js';

const css = {
  code: '.ds-list[data-display="list"] .ds-list__item > *{width:100%}.ds-list[data-display="list"] .ds-list__item > * .short-link-info{margin-bottom:2.5rem}.ds-list[data-display="grid"] .ds-list__item > *{width:250px;height:250px}.ds-list[data-display="grid"] .ds-list__item > * .short-link{position:absolute;top:-0.75rem;left:-1.5rem;right:-0px;width:auto}.ds-list[data-display="grid"] .ds-list__item > * .short-link>.short-link__content.svelte-1oq9iqk{width:100%;border-bottom-right-radius:0px}.ds-list[data-display="grid"] .ds-list__item > * .short-link-info{margin-top:2.5rem}.short-link.svelte-1oq9iqk{--link-background-color:var(--neutral-fg);--link-foreground-color:var(--neutral);--link-icon-shadow:0 0 0 0.25rem rgba(var(--link-foreground-color) / .3);position:relative;margin-top:0.75rem;margin-bottom:0.75rem;margin-left:0.75rem;display:flex;width:-moz-max-content;width:max-content;flex-direction:row;align-items:center}.short-link__icon.svelte-1oq9iqk{position:absolute;left:-0.75rem;z-index:10;display:flex;height:3rem;width:3rem;align-items:center;justify-content:center;border-radius:9999px;background-color:rgb(var(--link-background-color));padding:0.5rem;color:rgb(var(--link-foreground-color));box-shadow:var(--link-icon-shadow)}.short-link__content.svelte-1oq9iqk{margin-left:1.5rem;cursor:copy;-webkit-user-select:none;-moz-user-select:none;user-select:none;background-color:rgb(var(--link-background-color));padding-top:0.25rem;padding-bottom:0.25rem;padding-left:1.25rem;padding-right:0.75rem;color:rgb(var(--link-foreground-color));border-top-right-radius:var(--rounded-box);border-bottom-right-radius:var(--rounded-box)}.short-link-info.svelte-1oq9iqk{margin-top:1.25rem;display:flex;width:100%;flex-direction:column;overflow:hidden}.short-link-state-list.svelte-1oq9iqk{position:absolute;bottom:0px;left:0px;display:flex;align-items:center;overflow:hidden;border-top-left-radius:0px !important;border-top-left-radius:var(--rounded-box);border-bottom-left-radius:var(--rounded-box)}.short-link-state.svelte-1oq9iqk{display:flex;height:2.5rem;width:2.5rem;align-items:center;justify-content:center;padding:0.75rem;font-size:1.5rem;line-height:2rem}',
  map: '{"version":3,"file":"+page.svelte","sources":["+page.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\"><\/script>\\r\\n<script lang=\\"ts\\">import { CHANGE_MAIN_COLOR, showToast, SUCCESS } from \\"@/components/Toast\\";\\nimport { Button, NO_ANIMATION, NO_RIPPLE, PLAIN } from \\"@/widgets/actions/button\\";\\nimport { Card, COMPACT } from \\"@/widgets/containers/card\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nlet displayMode = \\"list\\";\\nlet links = [\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"proxy\\", useCount: -1, active: false },\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"redirect\\", useCount: -1, active: false },\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"proxy\\", useCount: -1, active: true },\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"redirect\\", useCount: -1, active: true },\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"proxy\\", useCount: -1, active: true },\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"redirect\\", useCount: -1, active: true },\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"proxy\\", useCount: -1, active: true },\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"redirect\\", useCount: 0, active: true },\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"proxy\\", useCount: 1, active: true },\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"redirect\\", useCount: 2, active: true },\\n  { id: BigInt(1), src: \\"https://www.youtube.com/watch?v=71pOiq-E_X4\\", short: \\"Fu884hxjuhjEDSU0\\", type: \\"proxy\\", useCount: 3, active: true }\\n];\\n<\/script>\\r\\n\\r\\n<div class=\\"flex flex-col gap-3 h-full max-h-full overflow-hidden\\">\\r\\n    <div class=\\"flex flex-col md:flex-row items-center justify-between mt-3\\">\\r\\n        <Card flags={COMPACT} class=\\"w-full md:mr-5\\">\\r\\n            <span>TODO</span>\\r\\n        </Card>\\r\\n        <Card flags={COMPACT} class=\\"ml-auto mt-5 md:mt-0\\">\\r\\n            <div class=\\"flex items-center bg-base-100\\" slot=\\"default\\">\\r\\n                <Button \\r\\n                    class={\\"!rounded-none w-10 h-10 !p-1 transition-all\\" + cn(displayMode == \\"list\\" ? \\"!text-primary\\" : \\"\\")}\\r\\n                    flags={NO_RIPPLE | NO_ANIMATION | PLAIN}\\r\\n                    OnClick={() => displayMode = \\"list\\"}\\r\\n                >\\r\\n                    <svg xmlns=\\"http://www.w3.org/2000/svg\\" width=\\"24\\" height=\\"24\\" viewBox=\\"0 0 24 24\\" style=\\"fill: currentColor;\\">\\r\\n                        <path d=\\"M4 6h2v2H4zm0 5h2v2H4zm0 5h2v2H4zm16-8V6H8.023v2H18.8zM8 11h12v2H8zm0 5h12v2H8z\\"></path>\\r\\n                    </svg>\\r\\n                </Button>\\r\\n                <Button \\r\\n                    class={\\"!rounded-none w-10 h-10 !p-1 transition-all\\" + cn(displayMode == \\"grid\\" ? \\"!text-primary\\" : \\"\\")}\\r\\n                    flags={NO_RIPPLE | NO_ANIMATION | PLAIN}\\r\\n                    OnClick={() => displayMode = \\"grid\\"}\\r\\n                >\\r\\n                    <svg xmlns=\\"http://www.w3.org/2000/svg\\" width=\\"24\\" height=\\"24\\" viewBox=\\"0 0 24 24\\" style=\\"fill: currentColor;\\">\\r\\n                        <path d=\\"M4 11h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zm10 0h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zM4 21h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zm10 0h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1z\\"></path>\\r\\n                    </svg>\\r\\n                </Button>\\r\\n            </div>\\r\\n        </Card>\\r\\n    </div>\\r\\n    \\r\\n    <div class=\\"ds-list py-4\\" data-display={displayMode}>\\r\\n        {#each links as item, idx (idx)}\\r\\n            {@const link = window.location.origin + \\"/urls/\\" + item.short}\\r\\n            <div class=\\"ds-list__item\\" data-id={item.id}>\\r\\n                <Card>\\r\\n                    <!-- svelte-ignore a11y-no-static-element-interactions -->\\r\\n                    <div class=\\"short-link\\">\\r\\n                        <a class=\\"short-link__icon\\" href={link} target=\\"_blank\\">\\r\\n                            <span class=\\"fa-solid fa-link\\"></span>\\r\\n                        </a>\\r\\n                        <!-- svelte-ignore a11y-click-events-have-key-events -->\\r\\n                        <span class=\\"short-link__content\\" on:click={() => {\\r\\n                            navigator.clipboard.writeText(link)\\r\\n                            showToast({\\r\\n                                data: {\\r\\n                                    title: \\"Success\\",\\r\\n                                    description: \\"Link was successfully copied to clipboard\\",\\r\\n                                    type: SUCCESS,\\r\\n                                    flags: CHANGE_MAIN_COLOR,\\r\\n                                }\\r\\n                            })\\r\\n                        }}>{item.short}</span>\\r\\n                    </div>\\r\\n                    <div class=\\"short-link-info\\">\\r\\n                        <div class=\\"flex flex-row overflow-hidden\\">\\r\\n                            <div class=\\"bg-base-300 text-black/50 p-2 px-3 rounded-l-box\\">\\r\\n                                <span class=\\"fa-solid fa-hashtag\\"></span>\\r\\n                            </div>\\r\\n                            <div class=\\"bg-base-200 flex items-center p-2 rounded-r-box overflow-hidden\\">\\r\\n                                <span class=\\"text-info truncate\\">{item.src}</span>\\r\\n                            </div>\\r\\n                        </div>\\r\\n                    </div>\\r\\n                    <div class=\\"short-link-state-list\\">\\r\\n                        {#if item.active}\\r\\n                            <span class=\\"short-link-state bg-success/20 text-success\\">\\r\\n                                <i class=\\"fa-solid fa-check\\"></i>\\r\\n                            </span>\\r\\n                        {:else}\\r\\n                            <span class=\\"short-link-state bg-error/20 text-error\\">\\r\\n                                <i class=\\"fa-solid fa-ban\\"></i>\\r\\n                            </span>\\r\\n                        {/if}\\r\\n                        <span class=\\"short-link-state bg-info/20 text-info\\">\\r\\n                            {#if item.useCount > 0}\\r\\n                                {item.useCount}\\r\\n                            {:else}\\r\\n                                <i class=\\"fa-solid fa-infinity\\"></i>\\r\\n                            {/if}\\r\\n                        </span>\\r\\n                    </div>\\r\\n                </Card>\\r\\n            </div>\\r\\n        {/each}\\r\\n    </div>\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    :global(.ds-list[data-display=\\"list\\"] .ds-list__item > *) {\\r\\n        width: 100%;\\r\\n    }\\r\\n    :global(.ds-list[data-display=\\"list\\"] .ds-list__item > * .short-link-info) {\\r\\n        margin-bottom: 2.5rem;\\r\\n}\\r\\n\\r\\n    :global(.ds-list[data-display=\\"grid\\"] .ds-list__item > *) {\\r\\n        width: 250px;\\r\\n        height: 250px;\\r\\n    }\\r\\n    :global(.ds-list[data-display=\\"grid\\"] .ds-list__item > * .short-link) {\\r\\n        position: absolute;\\r\\n        top: -0.75rem;\\r\\n        left: -1.5rem;\\r\\n        right: -0px;\\r\\n        width: auto;\\r\\n}\\r\\n    :global(.ds-list[data-display=\\"grid\\"] .ds-list__item > * .short-link) > .short-link__content {\\r\\n        width: 100%;\\r\\n        border-bottom-right-radius: 0px;\\r\\n}\\r\\n    :global(.ds-list[data-display=\\"grid\\"] .ds-list__item > * .short-link-info) {\\r\\n        margin-top: 2.5rem;\\r\\n}\\r\\n\\r\\n    .short-link {\\r\\n        --link-background-color: var(--neutral-fg);\\r\\n        --link-foreground-color: var(--neutral);\\r\\n        --link-icon-shadow: 0 0 0 0.25rem rgba(var(--link-foreground-color) / .3);\\r\\n        position: relative;\\r\\n        margin-top: 0.75rem;\\r\\n        margin-bottom: 0.75rem;\\r\\n        margin-left: 0.75rem;\\r\\n        display: flex;\\r\\n        width: -moz-max-content;\\r\\n        width: max-content;\\r\\n        flex-direction: row;\\r\\n        align-items: center;\\r\\n    }\\r\\n\\r\\n    .short-link__icon {\\r\\n        position: absolute;\\r\\n        left: -0.75rem;\\r\\n        z-index: 10;\\r\\n        display: flex;\\r\\n        height: 3rem;\\r\\n        width: 3rem;\\r\\n        align-items: center;\\r\\n        justify-content: center;\\r\\n        border-radius: 9999px;\\r\\n        background-color: rgb(var(--link-background-color));\\r\\n        padding: 0.5rem;\\r\\n        color: rgb(var(--link-foreground-color));\\r\\n\\r\\n            box-shadow: var(--link-icon-shadow);\\r\\n}\\r\\n\\r\\n    .short-link__content {\\r\\n        margin-left: 1.5rem;\\r\\n        cursor: copy;\\r\\n        -webkit-user-select: none;\\r\\n           -moz-user-select: none;\\r\\n                user-select: none;\\r\\n        background-color: rgb(var(--link-background-color));\\r\\n        padding-top: 0.25rem;\\r\\n        padding-bottom: 0.25rem;\\r\\n        padding-left: 1.25rem;\\r\\n        padding-right: 0.75rem;\\r\\n        color: rgb(var(--link-foreground-color));\\r\\n        border-top-right-radius: var(--rounded-box);\\r\\n        border-bottom-right-radius: var(--rounded-box);\\r\\n}\\r\\n\\r\\n    .short-link-info {\\r\\n        margin-top: 1.25rem;\\r\\n        display: flex;\\r\\n        width: 100%;\\r\\n        flex-direction: column;\\r\\n        overflow: hidden;\\r\\n}\\r\\n\\r\\n    .short-link-state-list {\\r\\n        position: absolute;\\r\\n        bottom: 0px;\\r\\n        left: 0px;\\r\\n        display: flex;\\r\\n        align-items: center;\\r\\n        overflow: hidden;\\r\\n        border-top-left-radius: 0px !important;\\r\\n        border-top-left-radius: var(--rounded-box);\\r\\n        border-bottom-left-radius: var(--rounded-box);\\r\\n}\\r\\n\\r\\n    .short-link-state {\\r\\n        display: flex;\\r\\n        height: 2.5rem;\\r\\n        width: 2.5rem;\\r\\n        align-items: center;\\r\\n        justify-content: center;\\r\\n        padding: 0.75rem;\\r\\n        font-size: 1.5rem;\\r\\n        line-height: 2rem;\\r\\n}\\r\\n</style>"],"names":[],"mappings":"AA4GY,gDAAkD,CACtD,KAAK,CAAE,IACX,CACQ,iEAAmE,CACvE,aAAa,CAAE,MACvB,CAEY,gDAAkD,CACtD,KAAK,CAAE,KAAK,CACZ,MAAM,CAAE,KACZ,CACQ,4DAA8D,CAClE,QAAQ,CAAE,QAAQ,CAClB,GAAG,CAAE,QAAQ,CACb,IAAI,CAAE,OAAO,CACb,KAAK,CAAE,IAAI,CACX,KAAK,CAAE,IACf,CACY,4DAA6D,CAAG,mCAAqB,CACzF,KAAK,CAAE,IAAI,CACX,0BAA0B,CAAE,GACpC,CACY,iEAAmE,CACvE,UAAU,CAAE,MACpB,CAEI,0BAAY,CACR,uBAAuB,CAAE,iBAAiB,CAC1C,uBAAuB,CAAE,cAAc,CACvC,kBAAkB,CAAE,qDAAqD,CACzE,QAAQ,CAAE,QAAQ,CAClB,UAAU,CAAE,OAAO,CACnB,aAAa,CAAE,OAAO,CACtB,WAAW,CAAE,OAAO,CACpB,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,gBAAgB,CACvB,KAAK,CAAE,WAAW,CAClB,cAAc,CAAE,GAAG,CACnB,WAAW,CAAE,MACjB,CAEA,gCAAkB,CACd,QAAQ,CAAE,QAAQ,CAClB,IAAI,CAAE,QAAQ,CACd,OAAO,CAAE,EAAE,CACX,OAAO,CAAE,IAAI,CACb,MAAM,CAAE,IAAI,CACZ,KAAK,CAAE,IAAI,CACX,WAAW,CAAE,MAAM,CACnB,eAAe,CAAE,MAAM,CACvB,aAAa,CAAE,MAAM,CACrB,gBAAgB,CAAE,IAAI,IAAI,uBAAuB,CAAC,CAAC,CACnD,OAAO,CAAE,MAAM,CACf,KAAK,CAAE,IAAI,IAAI,uBAAuB,CAAC,CAAC,CAEpC,UAAU,CAAE,IAAI,kBAAkB,CAC9C,CAEI,mCAAqB,CACjB,WAAW,CAAE,MAAM,CACnB,MAAM,CAAE,IAAI,CACZ,mBAAmB,CAAE,IAAI,CACtB,gBAAgB,CAAE,IAAI,CACjB,WAAW,CAAE,IAAI,CACzB,gBAAgB,CAAE,IAAI,IAAI,uBAAuB,CAAC,CAAC,CACnD,WAAW,CAAE,OAAO,CACpB,cAAc,CAAE,OAAO,CACvB,YAAY,CAAE,OAAO,CACrB,aAAa,CAAE,OAAO,CACtB,KAAK,CAAE,IAAI,IAAI,uBAAuB,CAAC,CAAC,CACxC,uBAAuB,CAAE,IAAI,aAAa,CAAC,CAC3C,0BAA0B,CAAE,IAAI,aAAa,CACrD,CAEI,+BAAiB,CACb,UAAU,CAAE,OAAO,CACnB,OAAO,CAAE,IAAI,CACb,KAAK,CAAE,IAAI,CACX,cAAc,CAAE,MAAM,CACtB,QAAQ,CAAE,MAClB,CAEI,qCAAuB,CACnB,QAAQ,CAAE,QAAQ,CAClB,MAAM,CAAE,GAAG,CACX,IAAI,CAAE,GAAG,CACT,OAAO,CAAE,IAAI,CACb,WAAW,CAAE,MAAM,CACnB,QAAQ,CAAE,MAAM,CAChB,sBAAsB,CAAE,GAAG,CAAC,UAAU,CACtC,sBAAsB,CAAE,IAAI,aAAa,CAAC,CAC1C,yBAAyB,CAAE,IAAI,aAAa,CACpD,CAEI,gCAAkB,CACd,OAAO,CAAE,IAAI,CACb,MAAM,CAAE,MAAM,CACd,KAAK,CAAE,MAAM,CACb,WAAW,CAAE,MAAM,CACnB,eAAe,CAAE,MAAM,CACvB,OAAO,CAAE,OAAO,CAChB,SAAS,CAAE,MAAM,CACjB,WAAW,CAAE,IACrB"}'
};
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let displayMode = "list";
  let links = [
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "proxy",
      useCount: -1,
      active: false
    },
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "redirect",
      useCount: -1,
      active: false
    },
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "proxy",
      useCount: -1,
      active: true
    },
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "redirect",
      useCount: -1,
      active: true
    },
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "proxy",
      useCount: -1,
      active: true
    },
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "redirect",
      useCount: -1,
      active: true
    },
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "proxy",
      useCount: -1,
      active: true
    },
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "redirect",
      useCount: 0,
      active: true
    },
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "proxy",
      useCount: 1,
      active: true
    },
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "redirect",
      useCount: 2,
      active: true
    },
    {
      id: BigInt(1),
      src: "https://www.youtube.com/watch?v=71pOiq-E_X4",
      short: "Fu884hxjuhjEDSU0",
      type: "proxy",
      useCount: 3,
      active: true
    }
  ];
  $$result.css.add(css);
  return `<div class="flex flex-col gap-3 h-full max-h-full overflow-hidden"><div class="flex flex-col md:flex-row items-center justify-between mt-3">${validate_component(Card, "Card").$$render($$result, { flags: COMPACT, class: "w-full md:mr-5" }, {}, {
    default: () => {
      return `<span data-svelte-h="svelte-19yhpqi">TODO</span>`;
    }
  })} ${validate_component(Card, "Card").$$render(
    $$result,
    {
      flags: COMPACT,
      class: "ml-auto mt-5 md:mt-0"
    },
    {},
    {
      default: () => {
        return `<div class="flex items-center bg-base-100" slot="default">${validate_component(Button, "Button").$$render(
          $$result,
          {
            class: "!rounded-none w-10 h-10 !p-1 transition-all" + cn(displayMode == "list" ? "!text-primary" : ""),
            flags: NO_RIPPLE | NO_ANIMATION | PLAIN,
            OnClick: () => displayMode = "list"
          },
          {},
          {
            default: () => {
              return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" style="fill: currentColor;"><path d="M4 6h2v2H4zm0 5h2v2H4zm0 5h2v2H4zm16-8V6H8.023v2H18.8zM8 11h12v2H8zm0 5h12v2H8z"></path></svg>`;
            }
          }
        )} ${validate_component(Button, "Button").$$render(
          $$result,
          {
            class: "!rounded-none w-10 h-10 !p-1 transition-all" + cn(displayMode == "grid" ? "!text-primary" : ""),
            flags: NO_RIPPLE | NO_ANIMATION | PLAIN,
            OnClick: () => displayMode = "grid"
          },
          {},
          {
            default: () => {
              return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" style="fill: currentColor;"><path d="M4 11h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zm10 0h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zM4 21h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zm10 0h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1z"></path></svg>`;
            }
          }
        )}</div>`;
      }
    }
  )}</div> <div class="ds-list py-4"${add_attribute("data-display", displayMode, 0)}>${each(links, (item, idx) => {
    let link = window.location.origin + "/urls/" + item.short;
    return ` <div class="ds-list__item"${add_attribute("data-id", item.id, 0)}>${validate_component(Card, "Card").$$render($$result, {}, {}, {
      default: () => {
        return ` <div class="short-link svelte-1oq9iqk"><a class="short-link__icon svelte-1oq9iqk"${add_attribute("href", link, 0)} target="_blank"><span class="fa-solid fa-link"></span></a>  <span class="short-link__content svelte-1oq9iqk">${escape(item.short)}</span></div> <div class="short-link-info svelte-1oq9iqk"><div class="flex flex-row overflow-hidden"><div class="bg-base-300 text-black/50 p-2 px-3 rounded-l-box" data-svelte-h="svelte-2gsc8d"><span class="fa-solid fa-hashtag"></span></div> <div class="bg-base-200 flex items-center p-2 rounded-r-box overflow-hidden"><span class="text-info truncate">${escape(item.src)}</span></div> </div></div> <div class="short-link-state-list svelte-1oq9iqk">${item.active ? `<span class="short-link-state bg-success/20 text-success svelte-1oq9iqk" data-svelte-h="svelte-oyrrhs"><i class="fa-solid fa-check"></i> </span>` : `<span class="short-link-state bg-error/20 text-error svelte-1oq9iqk" data-svelte-h="svelte-1fgwxax"><i class="fa-solid fa-ban"></i> </span>`} <span class="short-link-state bg-info/20 text-info svelte-1oq9iqk">${item.useCount > 0 ? `${escape(item.useCount)}` : `<i class="fa-solid fa-infinity"></i>`} </span></div> `;
      }
    })} </div>`;
  })}</div> </div>`;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-6qR8Pv6Q.js.map
