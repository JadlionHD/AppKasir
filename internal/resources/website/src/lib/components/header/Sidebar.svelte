<script lang="ts">
  import { Package2, Bell } from "lucide-svelte";

  import { Button } from "$lib/components/ui/button";
  import { NavigationItems } from ".";
</script>

<div class="hidden border-r bg-muted/40 md:block">
  <div class="flex h-full max-h-screen flex-col gap-2">
    <div class="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6">
      <a href="/dashboard" class="flex items-center gap-2 font-semibold">
        <Package2 class="h-6 w-6" />
        <span class="">AppKasir</span>
      </a>
      <Button variant="outline" size="icon" class="ml-auto h-8 w-8">
        <Bell class="h-4 w-4" />
        <span class="sr-only">Toggle notifications</span>
      </Button>
    </div>
    <div class="flex-1">
      <nav class="grid items-start px-2 text-sm font-medium lg:px-4">
        <NavigationItems />
      </nav>
    </div>
  </div>
</div>
