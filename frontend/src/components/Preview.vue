<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-col>
        <v-row justify="center">
          <v-container class="preview-container">
            <router-link to="/">
              <v-img class="preview-img" :src="preview" contain/>
            </router-link>
          </v-container>
        </v-row>
      </v-col>
      <v-col>
        <v-container class="preview-container">
          <v-list two-line>
            <template v-for="(item, index) in items">
              <v-list-item class="resource-list" @mouseover="preview = item.src" :key="item.title">
                <v-list-item-content>
                  <div class="text-h4" v-text="item.title" />
                  <div class="text-subtitle-1" v-text="item.description" />
                </v-list-item-content>
              </v-list-item>
              <v-divider v-if="index != items.length-1" :key="index" />
            </template>
          </v-list>
          <div class="button-container"> 
            <router-link to="/resources">
              <v-btn text>All resources ></v-btn>
            </router-link>
          </div>

        </v-container>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import APIClient from '../utils/APIClient';

export default {
  name: 'Preview',
  data: () => ({
    items: [],
    preview: "",
  }),
  mounted() {
    APIClient.previewAPI()
      .then((responseJson) => {
        this.items = responseJson;
        this.preview = responseJson[0].src
      });
  },
};
</script>

<style scoped>
  .preview-container {
    max-width: 600px;
    max-height: 600px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }
  .preview-img {
    max-width: 400px;
    max-height: 400px;
    width: auto;
    height: auto;
  }
  .resource-list:hover {
    transition-delay: 1s;
  }
  .button-container {
    height: 100%;
    width: 100%;
  }
  button{
    float: right;
  }
</style>
