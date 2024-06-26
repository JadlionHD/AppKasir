<script lang="ts">
  import { ShoppingCart, Package, Users, LineChart, House } from "lucide-svelte";

  import { NavigationLink } from ".";
</script>

<NavigationLink
  href="/dashboard"
  class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
  badge={10}
>
  <House class="h-4 w-4" />
  Dashboard
</NavigationLink>
<NavigationLink
  href="/transaksi"
  class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
>
  <ShoppingCart class="h-4 w-4" />
  Transaksi
</NavigationLink>
<NavigationLink
  href="/products"
  class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
>
  <Package class="h-4 w-4" />
  Products
</NavigationLink>
<NavigationLink
  href="/suppliers"
  class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
>
  <Users class="h-4 w-4" />
  Suppliers
</NavigationLink>
<NavigationLink
  href="/analytics"
  class="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
>
  <LineChart class="h-4 w-4" />
  Analytics
</NavigationLink>
