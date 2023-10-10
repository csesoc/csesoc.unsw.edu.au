<!--
    Student Resources
    --
    Links and display of important resources
-->

<template>
  <!-- Student Resources -->
  <div id="student-resources" class="content">
    <v-container class="resource--styles">
      <HeaderTitle title="STUDENT RESOURCES" />
        <v-row justify="space-between">
        <v-col cols="12" sm="12">
          <!-- Circles -->
          <a target="_blank" :href="resources.circles.link">
            <div data-cy="resources-box" class="box primary rounded-lg">
              <h2>{{ resources.circles.title }}</h2>
              <h3>{{ resources.circles.description }}</h3>
            </div>
          </a>
        </v-col>
        <v-col cols="12" sm="3">
          <!-- Structs -->
          <a target="_blank" :href="resources.structs.link">
            <div data-cy="resources-box" class="box primary rounded-lg">
              <h2>{{ resources.structs.title }}</h2>
              <h3>{{ resources.structs.description }}</h3>
            </div>
          </a>
        </v-col>

        <v-col cols="12" sm="3">
          <!-- Jobs -->
          <a target="_blank" :href="resources.jobs_board.link">
            <div data-cy="resources-box" class="box secondary rounded-lg">
              <h2>{{ resources.jobs_board.title }}</h2>
              <h3>{{ resources.jobs_board.description }}</h3>
            </div>
          </a>
        </v-col>
        <v-col cols="12" sm="3">
          <!-- Notangles -->
          <a target="_blank" :href="resources.notangles.link">
            <div data-cy="resources-box" class="box primary rounded-lg">
              <h2>{{ resources.notangles.title }}</h2>
              <h3>{{ resources.notangles.description }}</h3>
            </div>
          </a>
        </v-col>
        <v-col cols="12" sm="3">
          <!-- CSElectives -->
          <a target="_blank" :href="resources.cselectives.link">
            <div data-cy="resources-box" class="box secondary rounded-lg">
              <h2>{{ resources.cselectives.title }}</h2>
              <h3>{{ resources.cselectives.description }}</h3>
            </div>
          </a>
        </v-col>
        <v-col cols="12" sm="6">
          <!-- CompClub -->
          <a target="_blank" :href="resources.comp_club.link">
            <div data-cy="resources-box" class="box primary rounded-lg">
              <h2>{{ resources.comp_club.title }}</h2>
              <h3 class="hidden-sm-above">{{ resources.comp_club.description }}</h3>
              <div class="flex-grow-1"></div>
            </div>
          </a>
        </v-col>
        <v-col cols="12" sm="6">
          <!-- Media -->
          <a target="_blank" :href="resources.media.link">
            <div data-cy="resources-box" class="box primary rounded-lg">
              <h2>{{ resources.media.title }}</h2>
              <h3>{{ resources.media.description }}</h3>
            </div>
          </a>
        </v-col>
        <v-col cols="12" sm="6">
          <!-- First year guide -->
          <a target="_blank" :href="resources.fy_guide.link">
            <div data-cy="resources-box" class="box secondary rounded-lg">
              <h2>{{ resources.fy_guide.title }}</h2>
              <h3>{{ resources.fy_guide.description }}</h3>
            </div>
          </a>
        </v-col>
        <v-col cols="12" sm="6">
          <!-- Preenrollment guide -->
          <a target="_blank" :href="resources.pe_guide.link">
            <div data-cy="resources-box" class="box secondary rounded-lg">
              <h2>{{ resources.pe_guide.title }}</h2>
              <h3>{{ resources.pe_guide.description }}</h3>
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
    HeaderTitle
  },
  data: () => ({
    resources: {
      circles: {
        link: 'https://circles.csesoc.app/',
        title: 'Circles',
        description: 'A UNSW degree planner where you can explore and validate your degree structure.',
      },
      structs: {
        link: 'https://structs.sh',
        title: 'Structs',
        description: 'An educational data structures and algorithms platform.',
      },
      jobs_board: {
        link: 'https://jobsboard.csesoc.unsw.edu.au/',
        title: 'Jobs Board',
        description: 'A place where CSESoc students can look for relevant job opportunities.'
      },
      cselectives: {
        link: 'https://cselectives.csesoc.unsw.edu.au/',
        title: 'CSElectives',
        description: 'Read course reviews to help you pick your electives or even write your own!'
      },
      fy_guide: {
        link: 'https://media.csesoc.org.au/first-year-guide/',
        title: 'First Year Guide',
        description: 'The ultimate guide to conquering first year at CSE.'
      },
      media: {
        link: 'https://media.csesoc.org.au/',
        title: 'CSESoc Media',
        description: 'All things content'
      },
      comp_club: {
        link: 'https://compclub.org/',
        title: 'CompClub',
        description: 'Promoting computing to high school students'
      },
      notangles: {
        link: 'https://notangles.csesoc.unsw.edu.au/',
        title: 'Notangles',
        description: 'Trimester timetabling tool - no more timetable tangles!'
      },
      pe_guide: {
        link: 'https://media.csesoc.org.au/enrolment-guide/',
        title: 'Enrolment Guide',
        description: 'Learn how to get a headstart on uni'
      }
    }
  })
};
</script>

<style lang="scss" scoped>
// Student resources
.content {
  padding-top: 50px;
}
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
        overflow: auto;
        max-height: 250px;
        padding-top: 2vh;

        &.primary {
          @include linearGradient($primary-color, $secondary-color-2);
        }
        &.secondary {
          @include linearGradient($primary-color, $secondary-color-1);
        }

        h2 {
          padding-top: $space-xs/2;
          font-size: 1.8em;
          margin: 0;
          padding-left: 35px;
          padding-right: 35px;
          // margin-bottom: 0px;
        }

        h3 {
          padding-left: 35px;
          padding-right: 35px;
          font-size: 1.3em;
          // margin-bottom: 0px;
        }

        .image {
          border-bottom-left-radius: 10px;
          border-bottom-right-radius: 10px;
          top: 0;
          transition: transform 0.2s;
        }
      }

      .box:hover > .image {
        transform: scale(1.1);
      }
    }
  }
}
</style>
