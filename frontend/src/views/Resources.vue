<template>
  <v-container>
    <h1>Heading 1</h1>
    <h1>The quick brown fox jumps over the lazy dog.</h1>
    <h2>Heading 2</h2>
    <h2>The quick brown fox jumps over the lazy dog.</h2>
    <h3>Heading 3</h3>
    <h3>The quick brown fox jumps over the lazy dog.</h3>
    <h4>Heading 4</h4>
    <h4>The quick brown fox jumps over the lazy dog.</h4>
    <p>Paragraph</p>
    <p>
      Lorem ipsum dolor sit amet, consectetur adipiscing elit. In ex odio, condimentum quis urna quis, malesuada tin
      cidunt massa. Aliquam faucibus, odio eu scelerisque cursus, nisi erat gravida mi, sit amet co
      ngue odio nunc at erat. Phasellus tincidunt interdum velit. Vivamus porttitor, mi sed gravida varius, ex li
      gula consectetur nisi, at aliquam mauris lorem sed sapien. Quisque venenatis blandit tellus, in ullamcorper lectus ultrice
      s fermentum. Vestibulum at accumsan risus. Mauris pulvinar sodales arcu, quis egestas velit iaculis sit amet. Duis rutrum u
      llamcorper risus, in viverra eros lacinia ut. Duis tristique vestibulum erat vitae tincidunt. Integer ferment
      um, felis at aliquam pulvinar, lacus sapien tincidunt odio, eget aliquam orci elit et odio. Sed pulvinar, metus nec
      consectetur consequat, sem dolor interdum elit, a vulputate purus ex eget diam
      . Vestibulum fringilla justo ac lacinia pret
      ium. Integer imperdiet nec sem non imperdiet.
    </p>
    <!-- <v-row>
      <v-col class="pa-4" cols='6'>
        Some text
      </v-col>
      <v-col class="pa-4" cols='6'>
        Some other text
      </v-col>
    </v-row>-->
    <div class="flex-grid">
      <div class="col col--md-12">100%</div>
      <div class="col col--md-6">50%</div>
      <div class="col col--md-6">50%</div>
      <div class="col col--md-4">33%</div>
      <div class="col col--md-4">33%</div>
      <div class="col col--md-4">33%</div>
      <div class="col col--md-3">25%</div>
      <div class="col col--md-3">25%</div>
      <div class="col col--md-3">25%</div>
      <div class="col col--md-3">25%</div>
      <div class="col col--md-4">33%</div>
      <div class="col col--md-8">66%</div>
      <div class="col col--md-8">66%</div>
      <div class="col col--md-4">33%</div>
      <div class="col col--md-3">25%</div>
      <div class="col col--md-6">50%</div>
      <div class="col col--md-3">25%</div>
    </div>
  </v-container>
</template>

<script>
</script>

<style lang="scss" scoped>
.col {
  background-color: $primary-color;
}
</style>
