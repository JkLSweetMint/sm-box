import { c as create_ssr_component, g as getContext, v as validate_component, e as escape, b as add_attribute, d as each } from './ssr-C-9IsUTH.js';
import './Root-C1R-nd4p.js';
import './create-DPwALkVX.js';
import { c as cn } from './index-VQC3TRid.js';
import { C as COMPACT, a as Card, T as TextInput, S as Select } from './Card-DpmhQPy2.js';
import { P as PLAIN, C as CIRCLE, N as NO_RIPPLE, a as NO_ANIMATION, B as Button } from './Button-atP5DRSs.js';
import './index2-DbZx0BBT.js';
import './action-DplsO8dc.js';

const css = {
  code: '.ds-list[data-display="list"] .ds-list__item > *{overflow:hidden;width:100%}.ds-list[data-display="list"] .ds-list__item > * .short-link-info{margin-bottom:2.5rem}.ds-list[data-display="grid"] .ds-list__item > *{width:250px;height:250px}.ds-list[data-display="grid"] .ds-list__item > * .short-link{position:absolute;top:-0.75rem;left:-1.5rem;right:-0px;width:auto}.ds-list[data-display="grid"] .ds-list__item > * .short-link>.short-link__content.svelte-r4g6j8.svelte-r4g6j8{width:100%;border-bottom-right-radius:0px}.ds-list[data-display="grid"] .ds-list__item > * .short-link-info{margin-top:2.5rem}.ds-list[data-display="grid"] .ds-list__item > * .short-link-bubble{top:auto;bottom:0px;right:0px;align-items:center;justify-content:center;border-top-right-radius:0px !important;border-top-left-radius:0px !important;border-bottom-left-radius:0px !important;border-radius:var(--rounded-box)}.ds-list[data-display="grid"] .ds-list__item > * .short-link-bubble>.svelte-r4g6j8.svelte-r4g6j8{margin:0px}.short-link.svelte-r4g6j8.svelte-r4g6j8{--link-background-color:var(--neutral-fg);--link-foreground-color:var(--neutral);--link-icon-shadow:0 0 0 0.25rem rgba(var(--link-foreground-color) / .3);position:relative;margin-top:0.75rem;margin-bottom:0.75rem;margin-left:0.75rem;display:flex;width:-moz-max-content;width:max-content;flex-direction:row;align-items:center}.short-link__icon.svelte-r4g6j8.svelte-r4g6j8{position:absolute;left:-0.75rem;z-index:10;display:flex;height:3rem;width:3rem;align-items:center;justify-content:center;border-radius:9999px;background-color:rgb(var(--link-background-color));padding:0.5rem;color:rgb(var(--link-foreground-color));box-shadow:var(--link-icon-shadow)}.short-link__content.svelte-r4g6j8.svelte-r4g6j8{margin-left:1.5rem;cursor:copy;-webkit-user-select:none;-moz-user-select:none;user-select:none;background-color:rgb(var(--link-background-color));padding-top:0.25rem;padding-bottom:0.25rem;padding-left:1.25rem;padding-right:0.75rem;color:rgb(var(--link-foreground-color));border-top-right-radius:var(--rounded-box);border-bottom-right-radius:var(--rounded-box)}.short-link-info.svelte-r4g6j8.svelte-r4g6j8{margin-top:1.25rem;display:flex;width:100%;flex-direction:column;gap:0.75rem;overflow:hidden}.short-link-state-list.svelte-r4g6j8.svelte-r4g6j8{position:absolute;bottom:0px;left:0px;display:flex;align-items:center;overflow:hidden;border-top-left-radius:0px !important;border-top-left-radius:var(--rounded-box);border-bottom-left-radius:var(--rounded-box)}.short-link-state.svelte-r4g6j8.svelte-r4g6j8{display:flex;height:2.5rem;width:2.5rem;align-items:center;justify-content:center;padding:0.75rem;font-size:1.5rem;line-height:2rem}.short-link-bubble.svelte-r4g6j8.svelte-r4g6j8{position:absolute;top:-0.75rem;right:-0.75rem;display:flex;height:3rem;width:3rem;overflow:hidden;border-radius:9999px;background-color:rgb(var(--primary) / 0.2);padding:0.75rem;font-weight:600;--tw-text-opacity:1;color:rgb(var(--primary) / var(--tw-text-opacity))}.short-link-bubble.svelte-r4g6j8>.svelte-r4g6j8{margin-top:0.25rem;margin-left:0.25rem}',
  map: '{"version":3,"file":"+page.svelte","sources":["+page.svelte"],"sourcesContent":["<script lang=\\"ts\\" context=\\"module\\"><\/script>\\r\\n<script lang=\\"ts\\">import { getContext, onMount } from \\"svelte\\";\\nimport { CHANGE_MAIN_COLOR, showToast, SUCCESS } from \\"@/components/Toast\\";\\nimport { TextInput } from \\"@/widgets/inputs/textInput\\";\\nimport { Select } from \\"@/widgets/inputs/select\\";\\nimport { Button, CIRCLE, NO_ANIMATION, NO_RIPPLE, PLAIN } from \\"@/widgets/actions/button\\";\\nimport { Card, COMPACT } from \\"@/widgets/containers/card\\";\\nimport { cn } from \\"@/lib/helpers\\";\\nimport { QueryEntity, QueryEntityField } from \\"@/lib/QueryEntity\\";\\nconst callServiceMethod = getContext(\\"CallServiceMethod\\");\\nlet qm;\\nlet qmData = {\\n  search: \\"\\",\\n  sortName: \\"\\",\\n  sortType: \\"\\",\\n  page: 1,\\n  perPage: 25\\n};\\nlet displayMode = \\"list\\";\\nlet links = [];\\nconst LoadLinks = async function() {\\n  let query = {};\\n  if (qmData.search && qmData.search != \\"\\") query[\\"search\\"] = qmData.search;\\n  if (qmData.sortName && qmData.sortName != \\"\\") query[\\"sort_key\\"] = qmData.sortName;\\n  if (qmData.sortType && qmData.sortType != \\"\\") query[\\"sort_type\\"] = qmData.sortType;\\n  query[\\"offset\\"] = (qmData.page - 1) * qmData.perPage;\\n  query[\\"limit\\"] = qmData.perPage;\\n  const res = await callServiceMethod({\\n    service: \\"urls\\",\\n    url: \\"/management/list\\",\\n    query,\\n    showToast: true\\n  });\\n  if (res) links = res.data?.list.map((v) => ({\\n    id: BigInt(v[\\"id\\"]),\\n    src: v[\\"source\\"],\\n    short: v[\\"reduction\\"],\\n    type: v[\\"properties\\"][\\"type\\"],\\n    useCount: v[\\"properties\\"][\\"number_of_uses\\"] || -1,\\n    active: v[\\"properties\\"][\\"active\\"] || false\\n  })) ?? [];\\n};\\nonMount(() => {\\n  qm = new QueryEntity({\\n    fields: [\\n      new QueryEntityField({ name: \\"search\\" }),\\n      new QueryEntityField({ name: \\"page\\", type: \\"number\\", default: 1 }),\\n      new QueryEntityField({ name: \\"perPage\\", type: \\"number\\", default: 25 }),\\n      new QueryEntityField({ name: \\"sortName\\" }),\\n      new QueryEntityField({ name: \\"sortType\\" }),\\n      new QueryEntityField({ name: \\"filterActive\\", type: \\"boolean\\" }),\\n      new QueryEntityField({ name: \\"filterType\\" }),\\n      new QueryEntityField({ name: \\"filterStartActive\\", type: \\"date\\" }),\\n      new QueryEntityField({ name: \\"filterStartActiveType\\" }),\\n      new QueryEntityField({ name: \\"filterEndActive\\", type: \\"date\\" }),\\n      new QueryEntityField({ name: \\"filterEndActiveType\\" }),\\n      new QueryEntityField({ name: \\"filterNumberOfUses\\", type: \\"number\\" }),\\n      new QueryEntityField({ name: \\"filterNumberOfUsesType\\" })\\n    ]\\n  });\\n  qm.OnChange((event) => {\\n    qmData = qm.toObject();\\n  });\\n  qmData = qm.Load().toObject();\\n  LoadLinks();\\n});\\n<\/script>\\r\\n\\r\\n<div class=\\"flex flex-col gap-3 h-full max-h-full\\">\\r\\n    <div class=\\"flex items-center justify-between mt-3\\">\\r\\n        <Card flags={COMPACT} class=\\"w-full\\">\\r\\n            <div class=\\"flex flex-row flex-wrap gap-3\\">\\r\\n                <div class=\\"flex flex-col max-w-xs\\">\\r\\n                    <span class=\\"text-xs text-neutral-400\\">Search</span>\\r\\n                    <TextInput \\r\\n                        name=\\"search\\" \\r\\n                        value={qmData.search} \\r\\n                        onChange={e => {\\r\\n                            if (e.origin == \\"native\\") qm.SetProperty(\\"search\\", e.value)\\r\\n                        }}\\r\\n                    >\\r\\n                        <span class=\\"fa-solid fa-magnifying-glass text-neutral-400 w-4\\" slot=\\"prefix\\"></span>\\r\\n                    </TextInput>\\r\\n                </div>\\r\\n                <div class=\\"flex flex-col max-w-xs\\">\\r\\n                    <span class=\\"text-xs text-neutral-400\\">Sort</span>\\r\\n                    <Select name=\\"sort\\" value={qmData.sortName}>\\r\\n                        <i class=\\"fa-solid fa-sort text-neutral-400\\" slot=\\"prefix\\"></i>\\r\\n                        <svelte:fragment slot=\\"suffix\\">\\r\\n                            <Button class=\\"!p-1 w-6 h-6\\" flags={PLAIN | CIRCLE}>\\r\\n                                <i class=\\"fa-solid fa-minus\\"></i>\\r\\n                            </Button>\\r\\n                        </svelte:fragment>\\r\\n                    </Select>\\r\\n                </div>\\r\\n            </div>\\r\\n        </Card>\\r\\n    </div>\\r\\n\\r\\n    <div class=\\"flex flex-row justify-between\\">\\r\\n        <Card class=\\"w-full\\" flags={COMPACT}>\\r\\n            <div class=\\"flex flex-col items-center md:flex-row md:justify-between\\">\\r\\n                <div class=\\"flex flex-col md:flex-row items-center gap-3\\">\\r\\n                    <span>Items per page:</span>\\r\\n                    <Select\\r\\n                        class=\\"max-w-20\\"\\r\\n                        name=\\"perPage\\" \\r\\n                        searchText=\\"\\" \\r\\n                        noResultsText=\\"\\" \\r\\n                        value={qmData.perPage}\\r\\n                        options={[5, 10, 15, 20, 25, 50, 100, 150, 300]}\\r\\n                    >\\r\\n                    </Select>\\r\\n                    <div class=\\"flex items-center gap-2\\">\\r\\n                        <span>{ qmData.page }</span>\\r\\n                        <span>of</span>\\r\\n                        <span>x</span>\\r\\n                    </div>\\r\\n                    <div class=\\"flex items-center\\">\\r\\n                        <Button\\r\\n                            flags={NO_RIPPLE | NO_ANIMATION | PLAIN}\\r\\n                            class=\\"!rounded-none w-10 h-10 !p-1 transition-all\\"\\r\\n                        >\\r\\n                            <span class=\\"fa-solid fa-angles-left\\"></span>\\r\\n                        </Button>\\r\\n                        <Button\\r\\n                            flags={NO_RIPPLE | NO_ANIMATION | PLAIN}\\r\\n                            class=\\"!rounded-none w-10 h-10 !p-1 transition-all\\"\\r\\n                        >\\r\\n                            <span class=\\"fa-solid fa-angle-left\\"></span>\\r\\n                        </Button>\\r\\n                        <Button\\r\\n                            flags={NO_RIPPLE | NO_ANIMATION | PLAIN}\\r\\n                            class=\\"!rounded-none w-10 h-10 !p-1 transition-all\\"\\r\\n                        >\\r\\n                            <span class=\\"fa-solid fa-angle-right\\"></span>\\r\\n                        </Button>\\r\\n                        <Button\\r\\n                            flags={NO_RIPPLE | NO_ANIMATION | PLAIN}\\r\\n                            class=\\"!rounded-none w-10 h-10 !p-1 transition-all\\"\\r\\n                        >\\r\\n                            <span class=\\"fa-solid fa-angles-right\\"></span>\\r\\n                        </Button>\\r\\n                    </div>\\r\\n                </div>\\r\\n                <div class=\\"flex items-center\\">\\r\\n                    <Button \\r\\n                        class={\\"!rounded-none w-10 h-10 !p-1 transition-all\\" + cn(displayMode == \\"list\\" ? \\"!text-primary\\" : \\"\\")}\\r\\n                        flags={NO_RIPPLE | NO_ANIMATION | PLAIN}\\r\\n                        OnClick={() => displayMode = \\"list\\"}\\r\\n                    >\\r\\n                        <svg xmlns=\\"http://www.w3.org/2000/svg\\" width=\\"24\\" height=\\"24\\" viewBox=\\"0 0 24 24\\" style=\\"fill: currentColor;\\">\\r\\n                            <path d=\\"M4 6h2v2H4zm0 5h2v2H4zm0 5h2v2H4zm16-8V6H8.023v2H18.8zM8 11h12v2H8zm0 5h12v2H8z\\"></path>\\r\\n                        </svg>\\r\\n                    </Button>\\r\\n                    <Button \\r\\n                        class={\\"!rounded-none w-10 h-10 !p-1 transition-all\\" + cn(displayMode == \\"grid\\" ? \\"!text-primary\\" : \\"\\")}\\r\\n                        flags={NO_RIPPLE | NO_ANIMATION | PLAIN}\\r\\n                        OnClick={() => displayMode = \\"grid\\"}\\r\\n                    >\\r\\n                        <svg xmlns=\\"http://www.w3.org/2000/svg\\" width=\\"24\\" height=\\"24\\" viewBox=\\"0 0 24 24\\" style=\\"fill: currentColor;\\">\\r\\n                            <path d=\\"M4 11h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zm10 0h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zM4 21h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zm10 0h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1z\\"></path>\\r\\n                        </svg>\\r\\n                    </Button>\\r\\n                </div>\\r\\n            </div>\\r\\n        </Card>\\r\\n    </div>\\r\\n    \\r\\n    <div class=\\"ds-list py-4\\" data-display={displayMode}>\\r\\n        {#each links as item, idx (idx)}\\r\\n            {@const link = window.location.origin + \\"/urls/\\" + item.short}\\r\\n            <div class=\\"ds-list__item\\" data-id={item.id}>\\r\\n                <Card>\\r\\n                    <!-- svelte-ignore a11y-no-static-element-interactions -->\\r\\n                    <div class=\\"short-link\\">\\r\\n                        <a class=\\"short-link__icon\\" href={link} target=\\"_blank\\">\\r\\n                            <span class=\\"fa-solid fa-link\\"></span>\\r\\n                        </a>\\r\\n                        <!-- svelte-ignore a11y-click-events-have-key-events -->\\r\\n                        <span class=\\"short-link__content\\" on:click={() => {\\r\\n                            navigator.clipboard.writeText(link)\\r\\n                            showToast({\\r\\n                                data: {\\r\\n                                    title: \\"Success\\",\\r\\n                                    description: \\"Link was successfully copied to clipboard\\",\\r\\n                                    type: SUCCESS,\\r\\n                                    flags: CHANGE_MAIN_COLOR,\\r\\n                                }\\r\\n                            })\\r\\n                        }}>{item.short}</span>\\r\\n                    </div>\\r\\n                    <div class=\\"short-link-info\\">\\r\\n                        <div class=\\"flex flex-row overflow-hidden\\">\\r\\n                            <div class=\\"bg-base-300 text-black/50 p-2 px-3 rounded-l-box\\">\\r\\n                                <span class=\\"fa-solid fa-hashtag\\"></span>\\r\\n                            </div>\\r\\n                            <div class=\\"bg-base-200 flex items-center p-2 rounded-r-box overflow-hidden\\">\\r\\n                                <span class=\\"text-info truncate\\">{item.src}</span>\\r\\n                            </div>\\r\\n                        </div>\\r\\n                    </div>\\r\\n                    <div class=\\"short-link-state-list\\">\\r\\n                        {#if item.active}\\r\\n                            <span class=\\"short-link-state bg-success/20 text-success\\">\\r\\n                                <i class=\\"fa-solid fa-check\\"></i>\\r\\n                            </span>\\r\\n                        {:else}\\r\\n                            <span class=\\"short-link-state bg-error/20 text-error\\">\\r\\n                                <i class=\\"fa-solid fa-ban\\"></i>\\r\\n                            </span>\\r\\n                        {/if}\\r\\n                        <span class=\\"short-link-state bg-info/20 text-info\\">\\r\\n                            {#if item.useCount > 0}\\r\\n                                {item.useCount}\\r\\n                            {:else}\\r\\n                                <i class=\\"fa-solid fa-infinity\\"></i>\\r\\n                            {/if}\\r\\n                        </span>\\r\\n                    </div>\\r\\n                    <div class=\\"short-link-bubble\\">\\r\\n                        <span>{item.id}</span>\\r\\n                    </div>\\r\\n                </Card>\\r\\n            </div>\\r\\n        {/each}\\r\\n    </div>\\r\\n</div>\\r\\n\\r\\n<style lang=\\"postcss\\">\\r\\n    :global(.ds-list[data-display=\\"list\\"] .ds-list__item > *) {\\r\\n        overflow: hidden;\\r\\n        width: 100%;\\r\\n    }\\r\\n    :global(.ds-list[data-display=\\"list\\"] .ds-list__item > * .short-link-info) {\\r\\n        margin-bottom: 2.5rem;\\r\\n}\\r\\n\\r\\n    :global(.ds-list[data-display=\\"grid\\"] .ds-list__item > *) {\\r\\n        width: 250px;\\r\\n        height: 250px;\\r\\n    }\\r\\n    :global(.ds-list[data-display=\\"grid\\"] .ds-list__item > * .short-link) {\\r\\n        position: absolute;\\r\\n        top: -0.75rem;\\r\\n        left: -1.5rem;\\r\\n        right: -0px;\\r\\n        width: auto;\\r\\n}\\r\\n    :global(.ds-list[data-display=\\"grid\\"] .ds-list__item > * .short-link) > .short-link__content {\\r\\n        width: 100%;\\r\\n        border-bottom-right-radius: 0px;\\r\\n}\\r\\n    :global(.ds-list[data-display=\\"grid\\"] .ds-list__item > * .short-link-info) {\\r\\n        margin-top: 2.5rem;\\r\\n}\\r\\n    :global(.ds-list[data-display=\\"grid\\"] .ds-list__item > * .short-link-bubble) {\\r\\n        top: auto;\\r\\n        bottom: 0px;\\r\\n        right: 0px;\\r\\n        align-items: center;\\r\\n        justify-content: center;\\r\\n        border-top-right-radius: 0px !important;\\r\\n        border-top-left-radius: 0px !important;\\r\\n        border-bottom-left-radius: 0px !important;\\r\\n        border-radius: var(--rounded-box);\\r\\n}\\r\\n    :global(.ds-list[data-display=\\"grid\\"] .ds-list__item > * .short-link-bubble) > * {\\r\\n        margin: 0px;\\r\\n}\\r\\n\\r\\n    .short-link {\\r\\n        --link-background-color: var(--neutral-fg);\\r\\n        --link-foreground-color: var(--neutral);\\r\\n        --link-icon-shadow: 0 0 0 0.25rem rgba(var(--link-foreground-color) / .3);\\r\\n        position: relative;\\r\\n        margin-top: 0.75rem;\\r\\n        margin-bottom: 0.75rem;\\r\\n        margin-left: 0.75rem;\\r\\n        display: flex;\\r\\n        width: -moz-max-content;\\r\\n        width: max-content;\\r\\n        flex-direction: row;\\r\\n        align-items: center;\\r\\n    }\\r\\n\\r\\n    .short-link__icon {\\r\\n        position: absolute;\\r\\n        left: -0.75rem;\\r\\n        z-index: 10;\\r\\n        display: flex;\\r\\n        height: 3rem;\\r\\n        width: 3rem;\\r\\n        align-items: center;\\r\\n        justify-content: center;\\r\\n        border-radius: 9999px;\\r\\n        background-color: rgb(var(--link-background-color));\\r\\n        padding: 0.5rem;\\r\\n        color: rgb(var(--link-foreground-color));\\r\\n\\r\\n            box-shadow: var(--link-icon-shadow);\\r\\n}\\r\\n\\r\\n    .short-link__content {\\r\\n        margin-left: 1.5rem;\\r\\n        cursor: copy;\\r\\n        -webkit-user-select: none;\\r\\n           -moz-user-select: none;\\r\\n                user-select: none;\\r\\n        background-color: rgb(var(--link-background-color));\\r\\n        padding-top: 0.25rem;\\r\\n        padding-bottom: 0.25rem;\\r\\n        padding-left: 1.25rem;\\r\\n        padding-right: 0.75rem;\\r\\n        color: rgb(var(--link-foreground-color));\\r\\n        border-top-right-radius: var(--rounded-box);\\r\\n        border-bottom-right-radius: var(--rounded-box);\\r\\n}\\r\\n\\r\\n    .short-link-info {\\r\\n        margin-top: 1.25rem;\\r\\n        display: flex;\\r\\n        width: 100%;\\r\\n        flex-direction: column;\\r\\n        gap: 0.75rem;\\r\\n        overflow: hidden;\\r\\n}\\r\\n\\r\\n    .short-link-state-list {\\r\\n        position: absolute;\\r\\n        bottom: 0px;\\r\\n        left: 0px;\\r\\n        display: flex;\\r\\n        align-items: center;\\r\\n        overflow: hidden;\\r\\n        border-top-left-radius: 0px !important;\\r\\n        border-top-left-radius: var(--rounded-box);\\r\\n        border-bottom-left-radius: var(--rounded-box);\\r\\n}\\r\\n\\r\\n    .short-link-state {\\r\\n        display: flex;\\r\\n        height: 2.5rem;\\r\\n        width: 2.5rem;\\r\\n        align-items: center;\\r\\n        justify-content: center;\\r\\n        padding: 0.75rem;\\r\\n        font-size: 1.5rem;\\r\\n        line-height: 2rem;\\r\\n}\\r\\n\\r\\n    .short-link-bubble {\\r\\n        position: absolute;\\r\\n        top: -0.75rem;\\r\\n        right: -0.75rem;\\r\\n        display: flex;\\r\\n        height: 3rem;\\r\\n        width: 3rem;\\r\\n        overflow: hidden;\\r\\n        border-radius: 9999px;\\r\\n        background-color: rgb(var(--primary) / 0.2);\\r\\n        padding: 0.75rem;\\r\\n        font-weight: 600;\\r\\n        --tw-text-opacity: 1;\\r\\n        color: rgb(var(--primary) / var(--tw-text-opacity));\\r\\n}\\r\\n\\r\\n    .short-link-bubble > * {\\r\\n        margin-top: 0.25rem;\\r\\n        margin-left: 0.25rem;\\r\\n}\\r\\n</style>"],"names":[],"mappings":"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"}'
};
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  getContext("CallServiceMethod");
  let qm;
  let qmData = {
    search: "",
    sortName: "",
    sortType: "",
    page: 1,
    perPage: 25
  };
  let displayMode = "list";
  let links = [];
  $$result.css.add(css);
  return `<div class="flex flex-col gap-3 h-full max-h-full svelte-r4g6j8"><div class="flex items-center justify-between mt-3 svelte-r4g6j8">${validate_component(Card, "Card").$$render($$result, { flags: COMPACT, class: "w-full" }, {}, {
    default: () => {
      return `<div class="flex flex-row flex-wrap gap-3 svelte-r4g6j8"><div class="flex flex-col max-w-xs svelte-r4g6j8"><span class="text-xs text-neutral-400 svelte-r4g6j8" data-svelte-h="svelte-9x7aqy">Search</span> ${validate_component(TextInput, "TextInput").$$render(
        $$result,
        {
          name: "search",
          value: qmData.search,
          onChange: (e) => {
            if (e.origin == "native") qm.SetProperty("search", e.value);
          }
        },
        {},
        {
          prefix: () => {
            return `<span class="fa-solid fa-magnifying-glass text-neutral-400 w-4 svelte-r4g6j8" slot="prefix"></span>`;
          }
        }
      )}</div> <div class="flex flex-col max-w-xs svelte-r4g6j8"><span class="text-xs text-neutral-400 svelte-r4g6j8" data-svelte-h="svelte-rj37zm">Sort</span> ${validate_component(Select, "Select").$$render($$result, { name: "sort", value: qmData.sortName }, {}, {
        suffix: () => {
          return `${validate_component(Button, "Button").$$render(
            $$result,
            {
              class: "!p-1 w-6 h-6",
              flags: PLAIN | CIRCLE
            },
            {},
            {
              default: () => {
                return `<i class="fa-solid fa-minus svelte-r4g6j8"></i>`;
              }
            }
          )} `;
        },
        prefix: () => {
          return `<i class="fa-solid fa-sort text-neutral-400 svelte-r4g6j8" slot="prefix"></i>`;
        }
      })}</div></div>`;
    }
  })}</div> <div class="flex flex-row justify-between svelte-r4g6j8">${validate_component(Card, "Card").$$render($$result, { class: "w-full", flags: COMPACT }, {}, {
    default: () => {
      return `<div class="flex flex-col items-center md:flex-row md:justify-between svelte-r4g6j8"><div class="flex flex-col md:flex-row items-center gap-3 svelte-r4g6j8"><span class="svelte-r4g6j8" data-svelte-h="svelte-kn1faq">Items per page:</span> ${validate_component(Select, "Select").$$render(
        $$result,
        {
          class: "max-w-20",
          name: "perPage",
          searchText: "",
          noResultsText: "",
          value: qmData.perPage,
          options: [5, 10, 15, 20, 25, 50, 100, 150, 300]
        },
        {},
        {}
      )} <div class="flex items-center gap-2 svelte-r4g6j8"><span class="svelte-r4g6j8">${escape(qmData.page)}</span> <span class="svelte-r4g6j8" data-svelte-h="svelte-jmmfr">of</span> <span class="svelte-r4g6j8" data-svelte-h="svelte-1rhnbtu">x</span></div> <div class="flex items-center svelte-r4g6j8">${validate_component(Button, "Button").$$render(
        $$result,
        {
          flags: NO_RIPPLE | NO_ANIMATION | PLAIN,
          class: "!rounded-none w-10 h-10 !p-1 transition-all"
        },
        {},
        {
          default: () => {
            return `<span class="fa-solid fa-angles-left svelte-r4g6j8"></span>`;
          }
        }
      )} ${validate_component(Button, "Button").$$render(
        $$result,
        {
          flags: NO_RIPPLE | NO_ANIMATION | PLAIN,
          class: "!rounded-none w-10 h-10 !p-1 transition-all"
        },
        {},
        {
          default: () => {
            return `<span class="fa-solid fa-angle-left svelte-r4g6j8"></span>`;
          }
        }
      )} ${validate_component(Button, "Button").$$render(
        $$result,
        {
          flags: NO_RIPPLE | NO_ANIMATION | PLAIN,
          class: "!rounded-none w-10 h-10 !p-1 transition-all"
        },
        {},
        {
          default: () => {
            return `<span class="fa-solid fa-angle-right svelte-r4g6j8"></span>`;
          }
        }
      )} ${validate_component(Button, "Button").$$render(
        $$result,
        {
          flags: NO_RIPPLE | NO_ANIMATION | PLAIN,
          class: "!rounded-none w-10 h-10 !p-1 transition-all"
        },
        {},
        {
          default: () => {
            return `<span class="fa-solid fa-angles-right svelte-r4g6j8"></span>`;
          }
        }
      )}</div></div> <div class="flex items-center svelte-r4g6j8">${validate_component(Button, "Button").$$render(
        $$result,
        {
          class: "!rounded-none w-10 h-10 !p-1 transition-all" + cn(displayMode == "list" ? "!text-primary" : ""),
          flags: NO_RIPPLE | NO_ANIMATION | PLAIN,
          OnClick: () => displayMode = "list"
        },
        {},
        {
          default: () => {
            return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" style="fill: currentColor;" class="svelte-r4g6j8"><path d="M4 6h2v2H4zm0 5h2v2H4zm0 5h2v2H4zm16-8V6H8.023v2H18.8zM8 11h12v2H8zm0 5h12v2H8z" class="svelte-r4g6j8"></path></svg>`;
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
            return `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" style="fill: currentColor;" class="svelte-r4g6j8"><path d="M4 11h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zm10 0h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zM4 21h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zm10 0h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1z" class="svelte-r4g6j8"></path></svg>`;
          }
        }
      )}</div></div>`;
    }
  })}</div> <div class="ds-list py-4 svelte-r4g6j8"${add_attribute("data-display", displayMode, 0)}>${each(links, (item, idx) => {
    let link = window.location.origin + "/urls/" + item.short;
    return ` <div class="ds-list__item svelte-r4g6j8"${add_attribute("data-id", item.id, 0)}>${validate_component(Card, "Card").$$render($$result, {}, {}, {
      default: () => {
        return ` <div class="short-link svelte-r4g6j8"><a class="short-link__icon svelte-r4g6j8"${add_attribute("href", link, 0)} target="_blank"><span class="fa-solid fa-link svelte-r4g6j8"></span></a>  <span class="short-link__content svelte-r4g6j8">${escape(item.short)}</span></div> <div class="short-link-info svelte-r4g6j8"><div class="flex flex-row overflow-hidden svelte-r4g6j8"><div class="bg-base-300 text-black/50 p-2 px-3 rounded-l-box svelte-r4g6j8" data-svelte-h="svelte-2gsc8d"><span class="fa-solid fa-hashtag svelte-r4g6j8"></span></div> <div class="bg-base-200 flex items-center p-2 rounded-r-box overflow-hidden svelte-r4g6j8"><span class="text-info truncate svelte-r4g6j8">${escape(item.src)}</span></div> </div></div> <div class="short-link-state-list svelte-r4g6j8">${item.active ? `<span class="short-link-state bg-success/20 text-success svelte-r4g6j8" data-svelte-h="svelte-oyrrhs"><i class="fa-solid fa-check svelte-r4g6j8"></i> </span>` : `<span class="short-link-state bg-error/20 text-error svelte-r4g6j8" data-svelte-h="svelte-1fgwxax"><i class="fa-solid fa-ban svelte-r4g6j8"></i> </span>`} <span class="short-link-state bg-info/20 text-info svelte-r4g6j8">${item.useCount > 0 ? `${escape(item.useCount)}` : `<i class="fa-solid fa-infinity svelte-r4g6j8"></i>`} </span></div> <div class="short-link-bubble svelte-r4g6j8"><span class="svelte-r4g6j8">${escape(item.id)}</span></div> `;
      }
    })} </div>`;
  })}</div> </div>`;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-jU1n5x17.js.map
