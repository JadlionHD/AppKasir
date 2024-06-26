<script lang="ts">
  import { Button } from "$lib/components/ui/button";
  import axios from "axios";
  import * as Card from "$lib/components/ui/card";
  import * as Select from "$lib/components/ui/select";
  import { Input } from "$lib/components/ui/input";
  import { Label } from "$lib/components/ui/label";
  import { Typography } from "$lib/components/typography";
  import { onMount } from "svelte";

  let username: string;
  let password: string;

  const handleLogin = async () => {
    // fetch("http://localhost:80/api/login", {
    //   method: "POST",
    //   body: JSON.stringify({
    //     username,
    //     password
    //   }),
    //   headers: {
    //     Accept: "application/json",
    //     "Content-Type": "application/json"
    //   }
    // });

    const data = await axios.post("http://localhost:80/api/login", {
      username,
      password
    });

    console.log(data);
  };
</script>

<div class="flex h-screen">
  <div class="m-auto">
    <Card.Root class="w-[350px]">
      <Card.Header>
        <Card.Title class="text-2xl">Login</Card.Title>
        <Card.Description
          >Untuk memasuki AppKasir, anda perlu login terlebih dahulu.</Card.Description
        >
      </Card.Header>
      <Card.Content>
        <div>
          <div class="grid w-full items-center gap-4">
            <div class="flex flex-col space-y-1.5">
              <Label for="username">Username</Label>
              <Input id="username" bind:value={username} placeholder="Masukkan Username" />
            </div>
            <div class="flex flex-col space-y-1.5">
              <Label for="password">Password</Label>
              <Input
                id="password"
                bind:value={password}
                type="password"
                placeholder="Masukkan Password"
              />
            </div>
          </div>
        </div>
      </Card.Content>
      <Card.Footer class="flex justify-between">
        <Button class="w-full" type="submit" on:click={handleLogin}>Masuk</Button>
      </Card.Footer>
    </Card.Root>
  </div>
</div>
