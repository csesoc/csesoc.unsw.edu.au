<!--
  Footer
  --
  This component contains footer elements of the website, visible from all pages.
-->

<template>
  <v-footer color="#000033" padless>
    <v-container class="white--text" style="padding:60px 72px" fluid>
      <div class="white-text" data-cy="footer-address">
        <a href="/">
          <img class="img-size" src="@/assets/csesoc-logo-white.svg" data-cy="footer-logo" />
        </a>
        <br />B03 CSE Building K17, UNSW
        <br />
        <a href="mailto:csesoc@csesoc.org.au"> csesoc@csesoc.org.au </a>
      </div>
      © 2021 — CSESoc UNSW
    </v-container>
  </v-footer>
</template>

<script type="text/javascript">
export default {
  name: 'Footer'
};
</script>

<style lang="scss" scoped>
.white--text {
  display: flex;
  align-items: center;
  flex-direction: column;
}

.white-text {
  @extend p;
}

.white-text a {
  text-decoration: none;
  color: $light-color;
  padding-right: $space-xxs;
}

.white-text a:hover {
  font-weight: 900;
  transition-duration: 0.1s;
}

.nav-divider {
  padding-bottom: $space-xxs;
  border-bottom: 2px solid white;
}

.anchor-items {
  margin-left: -15px;
}

.img-size {
  max-width: 300px;
  max-height: 200px;
}
</style>
