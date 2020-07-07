<template>
  <h1>$ {{ title }}</h1>
</template>

<style scoped>
h1 {
  font-size: 2rem;
  font-weight: bold;
  margin-bottom: 1.3rem;
  text-transform: uppercase;
}
</style>

<script>
export default {
  name: 'HeaderTitle',
  props: ['title']
};
</script>
