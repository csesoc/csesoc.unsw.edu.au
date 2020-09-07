<!--
  Showcase
  --
  This view corresponds to the topmost element of the landing page.
-->

<template>
    <div id="showcase" class="content">
        <v-container>
          <v-row no-gutters>
            <v-col class="pa-0" cols="6">
              <h1>
                We inspire the programmers of the future.
              </h1>
            </v-col>
          </v-row>
          <v-row no-gutters>
            <v-col class="pa-0" cols="12">
              <button
                class="btn--join-us"
                data-cy=joinus-button
                @click="scrollto('joinus')">
                  Join Us
              </button>
            </v-col>
          </v-row>
        </v-container>
    </div>
</template>

<style lang="scss" scoped>
#showcase {
  background-image: linear-gradient(transparent, transparent, $dark-color-1), url(../../assets/landingPageHeaderBackground.png);
  background-position: 75% 50%;
  height: 100vh;
  align-items: center;
  justify-content: center;
  background-size: 150%;
}

.btn--join-us {
  @extend .btn--lg;
  float: right;
}

.btn--join-us:hover {
  transition: 0.4s;
  background: rgba(102,255,255, 0.2);
}
</style>
