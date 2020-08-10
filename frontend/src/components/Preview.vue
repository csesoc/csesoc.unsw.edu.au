<!--
  Preview
  --
  This component is a preview visualiser of CSE resource links defined in the item objects passed in. 
  --
  Props:
    - items: list of times - containing the fields: id, link, src, title
-->

<template>
  <v-container>
    <v-row align='center' justify='center'>
      <v-col>
        <v-row justify='center'>
          <v-container class='preview-container'>
            <v-progress-circular v-if="loading()" indeterminate />
            <v-img v-else class='preview-img' :src="previewImg" contain data-cy="preview-image" />
          </v-container>
        </v-row>
      </v-col>
      <v-col>
        <v-container class='preview-container'>
          <v-list two-line>
            <template v-for="(item, index) in items">
              <v-list-item 
              v-if="item.link !== ''" 
              class='preview-item' 
              :href="item.link" 
              target="_blank" 
              @mouseover="previewImg = item.src" 
              data-cy="preview-item" 
              :key="item.title">
                <v-list-item-content>
                  <div v-if="item.title !== ''" class='text-h4' v-text="item.title" data-cy='preview-title' />
                  <div v-if="item.description !== ''" class='text-subtitle-1' v-text="item.description"  data-cy='preview-description' />
                </v-list-item-content>
              </v-list-item>
              <v-divider v-if="index != items.length-1" :key="index" />
            </template>
          </v-list>
          <div class='button-container'> 
            <router-link to='/resources'>
              <v-btn text>All resources ></v-btn>
            </router-link>
          </div>

        </v-container>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

export default {
  name: 'Preview',
  props: ['items'],
  data: function () {
    return {
      previewImg: '',
    }
  },
  methods: {
    loading: function() {
      if(this.items.length < 1) {
        return true
      } else if (this.previewImg === '') {
        this.previewImg = this.items[0].src
      }
      return false
    }
  }
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

.preview-item:hover {
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
