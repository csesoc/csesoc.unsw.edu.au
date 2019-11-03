<template>
  <!-- Insert post here-->
  <div class="listComponent">
    <v-container class="pa-md-1 mx-lg-auto">
      <v-row>
        <v-col class="pa-2" v-if="piclist == true">
          <img v-if="image.length != 0" :src="image" />
        </v-col>
        <v-col>
          <div class="pa-md-6">
            <h2 class="listItemTitle">
              {{ title }}
            </h2>
            <h3 class="listItemSubtitle">{{ subtitle }}</h3>
            <p>
              {{ content }}
            </p>
            <div class="text-center">
              <router-link to="/post"
                ><v-btn rounded color="primary" dark style="float: left;"
                  >More -></v-btn
                ></router-link
              >
            </div>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>


<style scoped>
@import url("https://fonts.googleapis.com/css?family=Quicksand&display=swap");

h1 {
  font-size: 48px;
  font-weight: normal;
}

body {
  font-family: "Quicksand", sans-serif;
  text-align: left;
}

p {
  font-size: 14px;
  font-weight: 300;
  -webkit-font-smoothing: subpixel-antialiased;
}

p span {
  font-size: 16px;
  font-style: italic;
  display: inline-block;
}

.underline {
  border-bottom: 1px solid grey;
}

img {
  max-width: 100%;
  max-height: 100%;
}
</style>


<script>
    export default {
        name: 'ListComponent',
        data() {
            return {
                piclist: false,
                image: 'https://avatars0.githubusercontent.com/u/164179?s=280&v=4',
                title: 'Header',
                subtitle: 'Subtitle',
                content: 'This is some text.askjdbasbakc id isd ks  sj ks skdj dsl sdkl jsdlk sdlk sdk skd sdlj vld ldckhj sdk hsjjsdn sjd skd sjhds ksdk jsdk lskd sljadksa kjx slklkj sk sdk kld dk s    ds jsdklsdkdslk dslkj dkds vs dkjsd lksd lsj sdjhbcjsdc sdjhcb ',
                link: '',
            };
        },
    };
</script>
