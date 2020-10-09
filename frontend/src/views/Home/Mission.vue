<!--
  Mission
  --
  Mission statement of CSESOC
-->

<template>
   <div id="mission" class="content">
    <v-container>
    <HeaderTitle :title="'CSESOC\'S MISSION'" />
    <v-row no-gutters justify="end">
      <v-col cols="12" class="mission--p">
        <p class="text--md">
          {{mission}}
        </p>
      </v-col>
    </v-row>
    </v-container>
  </div>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';

export default {
  name: 'Mission',
  components: {
    HeaderTitle
  },
  data: () => ({
    mission: '"To empower every person and every organization on the planet to achieve more. We strive to create local opportunity, growth, and impact in every country around the world."',
  }),
};
</script>
<style lang="scss" scoped>
#mission {
//   background-image: linear-gradient($dark-color-1, transparent, $dark-color-1), url(../../assets/mission.jpg);
  height: 100vh;
//   background-size: 50%;
//   padding-top: $space-sm;
}
</style>
