<script lang="ts">
  import { page } from "$app/stores";
  import { cn } from "$lib/utils.js";
  import type { HTMLAttributes } from "svelte/elements";
  import { Badge } from "$lib/components/ui/badge/index.js";

  type $$Props = HTMLAttributes<HTMLDivElement> & {
    title?: string;
    href?: string;
    badge?: number;
  };

  let className: $$Props["class"] = undefined;
  export let href: $$Props["href"] = "#";
  export let badge: $$Props["badge"] = undefined;
  export { className as class };
  let active = $page.url.pathname === href;
</script>

<a
  {href}
  class={active
    ? "flex items-center gap-3 rounded-lg px-3 py-2 bg-muted text-primary transition-all hover:text-primary"
    : cn(className)}
>
  <slot />
  {#if badge}
    <Badge
      variant={active ? "outline" : "default"}
      class="ml-auto flex h-6 w-6 shrink-0 items-center justify-center rounded-full"
    >
      {badge}
    </Badge>
  {/if}
</a>
