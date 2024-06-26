<script lang="ts">
  import { CurveType } from "@unovis/ts";
  import { VisBulletLegend, VisTooltip, VisSingleContainer, VisDonut } from "@unovis/svelte";
  import { Donut } from "@unovis/ts";

  let data = [10, 50, 23, 100];
  const triggers = { [Donut.selectors.segment]: (d: { data: any }) => d.data };
  const value = (d: number) => d;
</script>

<VisSingleContainer {data}>
  <VisTooltip {triggers} />
  <VisDonut {value} />
</VisSingleContainer>
