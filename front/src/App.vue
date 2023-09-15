<script setup lang="ts">
import ChanReply from '@/components/ChanReply.vue'
</script>

<template>
  <chan-reply :id="1"></chan-reply>
  <chan-reply :id="2"></chan-reply>
  <chan-reply :id="3"></chan-reply>
  <chan-reply :id="3"></chan-reply>
</template>

<style>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
