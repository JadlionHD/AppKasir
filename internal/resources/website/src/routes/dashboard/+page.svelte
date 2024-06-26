<script lang="ts">
  import { Badge } from "$lib/components/ui/badge/index.js";
  import { Button } from "$lib/components/ui/button/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import * as Sheet from "$lib/components/ui/sheet/index.js";
  import { Sidebar, Navbar } from "$lib/components/header";
</script>

<div class="flex items-center">
  <h1 class="text-lg font-semibold md:text-2xl">Inventory</h1>
</div>
<div class="flex flex-1 items-center justify-center rounded-lg border border-dashed shadow-sm">
  <div class="flex flex-col items-center gap-1 text-center">
    <h3 class="text-2xl font-bold tracking-tight">You have no products</h3>
    <p class="text-sm text-muted-foreground">You can start selling as soon as you add a product.</p>
    <Button class="mt-4">Add Product</Button>
  </div>
</div>
