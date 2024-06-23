<script lang="ts">
  import { type VariantsType } from "./index";
  type $$Variants = VariantsType;
  export let variant: $$Variants = "p";
</script>

{#if variant === "h1"}
  <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
    <slot />
  </h1>
{:else if variant === "h2"}
  <h2
    class="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
  >
    <slot />
  </h2>
{:else if variant === "h3"}
  <h3 class="scroll-m-20 text-2xl font-semibold tracking-tight"><slot /></h3>
{:else if variant === "h4"}
  <h4 class="scroll-m-20 text-xl font-semibold tracking-tight"><slot /></h4>
{:else if variant === "p"}
  <p class="leading-7 [&:not(:first-child)]:mt-6">
    <slot />
  </p>
{:else if variant === "blockquote"}
  <blockquote class="mt-6 border-l-2 pl-6 italic">
    <slot />
  </blockquote>
{:else if variant === "inline-code"}
  <code class="relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold">
    <slot />
  </code>
{:else if variant === "lead"}
  <p class="text-xl text-muted-foreground">
    <slot />
  </p>
{:else if variant === "large"}
  <div class="text-lg font-semibold"><slot /></div>
{:else if variant === "small"}
  <small class="text-sm font-medium leading-none"><slot /></small>
{:else if variant === "muted"}
  <p class="text-sm text-muted-foreground"><slot /></p>
{/if}
