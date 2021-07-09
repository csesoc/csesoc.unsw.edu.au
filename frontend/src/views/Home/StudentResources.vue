<!--
    Student Resources
    --
    Links and display of important resources
-->

<template>
  <!-- Student Resources -->
  <div id="student-resouces" class="content">
    <v-container class="resource--styles">
      <HeaderTitle title="STUDENT RESOURCES" />

      <v-row no-gutters justify="space-between">
        <v-col xs="12" sm="7">
          <a target="_blank" :href="resources.big_box_1.link">
          <div data-cy="resources-box" class="box big">
            <h2>{{resources.big_box_1.title}}</h2>
            <h3>{{resources.big_box_1.description}}</h3>
            <v-img data-cy="resources-img" :src="resources.big_box_1.src" contain class="image" />
          </div>
          </a>
        </v-col>
        <v-col xs="12" sm="4">
          <a target="_blank" :href="resources.small_box_1.link">
          <div data-cy="resources-box" class="box small rounded-lg hidden-xs-only">
            <h2>{{resources.small_box_1.title}}</h2>
            <v-img data-cy="resources-img" :src="resources.small_box_1.src" contain class="image" />
          </div>
          </a>
        </v-col>
        <v-col xs="12" sm="4">
          <a target="_blank" :href="resources.small_box_2.link">
          <div data-cy="resources-box" class="box small rounded-lg hidden-xs-only">
            <h2>{{resources.small_box_2.title}}</h2>
            <v-img data-cy="resources-img" :src="resources.small_box_2.src" class="image" />
          </div>
          </a>
        </v-col>
        <v-col xs="12" sm="7">
          <a target="_blank" :href="resources.small_box_2.link">
          <div data-cy="resources-box" class="box big rounded-lg">
            <h2>{{resources.big_box_2.title}}</h2>
            <h3>{{resources.big_box_2.description}}</h3>
            <v-img data-cy="resources-img" :src="resources.big_box_2.src" class="image" />
          </div>
          </a>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';

export default {
  components: {
    HeaderTitle,
  },
  data: () => ({
    resources: {
      big_box_1: {
        link: 'https://jobsboard.csesoc.unsw.edu.au/login',
        title: 'Job Board',
        description: 'A place where CSESoc students can look for relevant job opportunities.',
        src: require('../../assets/resource-job-board.png'),
      },
      small_box_1: {
        link: 'https://media.csesoc.org.au/first-year-guide/',
        title: 'First Year Guide',
        src: require('../../assets/resource-first-year-guide.png')
      },
      small_box_2: {
        link: 'https://media.csesoc.org.au/',
        title: 'CSESOC Media',
        src: require('../../assets/resource-media.png'),
      },
      big_box_2: {
        link: 'https://notangles.csesoc.unsw.edu.au/',
        title: 'Notangles',
        description: 'Trimster timetabling tool - no more timetable tangles!',
        src: require('../../assets/resource-notangles.png'),
      }
    },
  })
};
</script>

<style lang="scss" scoped>
// Student resources
.resource--styles {
  color: $light-color;

  .row {
    text-decoration: none;

    .col {
      margin-bottom: $space-xxs;
    }

    a {
      color: inherit;
      text-decoration: none;

      .box {
        border-radius: 10px;
        display: flex;
        flex-direction: column;
        height: 100%;
        overflow: hidden;

        &.big{
          @include linearGradient($primary-color, $secondary-color-2);
        }
        &.small{
          @include linearGradient($primary-color, $secondary-color-1);
        }

        h2 {
          padding-top: $space-xs;
          padding-bottom: $space-xxs;
          margin: 0;
          padding-left: 35px;
          padding-right: 35px;
        }

        h3 {
          padding-left: 35px;
          padding-right: 35px;
        }

        .image {
          border-bottom-left-radius: 10px;
          border-bottom-right-radius: 10px;
          transition: transform .2s;
        }
      }

      .box:hover > .image {
        transform: scale(1.1);
      }
    }
  }
}
</style>
