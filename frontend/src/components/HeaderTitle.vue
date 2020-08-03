<!--
  HeaderTitle
  --
  This component simply formats the title.
  --
  Props:
    - title: string
-->

<template>
  <h1># {{ title }}</h1>
</template>

<script>
export default {
  name: 'HeaderTitle',
  props: ['title']
};
</script>

<style scoped>
h1 {
  font-size: 60px;
  font-weight: bold;
  margin-bottom: 1.3rem;
  text-transform: lowercase;
}
</style>
