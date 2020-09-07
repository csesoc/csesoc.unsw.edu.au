<!--
  HeaderTitle
  --
  This component simply formats the title. Ensure this component
  is used for all titles to keep consistency across the site.
  --
  Props:
    - title: string
-->

<template>
  <div class="header-title">
    <h1>{{ title }}</h1>
    <h4 v-if="subtitle !== false">{{ subtitle }}</h4>
  </div>
</template>

<script>
export default {
  name: 'HeaderTitle',
  props: {
    title: {
      type: String,
      required: true,
    },
    subtitle: {
      default: false,
    }
  }
};
</script>

<style lang="scss" scoped>
</style>
