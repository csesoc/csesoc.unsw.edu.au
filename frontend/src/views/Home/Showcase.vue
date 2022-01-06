<!--
  Showcase
  --
  This view corresponds to the topmost element of the landing page.
-->

<template>
  <div id="showcase" class="content">
    <v-container>
      <h1>
        {{ quote }}
      </h1>
      <h2>
        {{ tag }}
      </h2>
    </v-container>
  </div>
</template>

<script>
export default {
  data: () => ({
    quote: 'Hello World!',
    tag: 'We are the principal representative body for UNSW computing students'
  }),
  methods: {
    joinusClick() {
      window.open('https://member.arc.unsw.edu.au/s/clubdetail?clubid=0016F0000371VxAQAU', '_blank');
    }
  }
};
</script>

<style lang="scss" scoped>
#showcase {
  background-image: linear-gradient(transparent, transparent, $dark-color-1),
    url(../../assets/landing_page/showcase.png);
  background-size: cover;
  background-attachment: local;
  background-position: 75% 50%;
  height: 100vh;
  align-items: center;
  justify-content: center;
}

.btn--join-us {
  @extend .btn--lg;
  float: right;
}

.btn--join-us:hover {
  transition: 0.4s;
  background: rgba(102, 255, 255, 0.2);
}
</style>
