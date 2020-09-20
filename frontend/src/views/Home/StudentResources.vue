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
          <a target="_blank" href="https://jobsboard.csesoc.unsw.edu.au/login">
          <div data-cy="resources-box" class="box big">
            <h2>Job Board</h2>
            <h3>A place where CSESoc students can look for relevant job opportunities.</h3>
            <v-img data-cy="resources-img" src="@/assets/resource-job-board.png" contain class="image" />
          </div>
          </a>
        </v-col>
        <v-col xs="12" sm="4">
          <a target="_blank" href="">
          <div data-cy="resources-box" class="box small rounded-lg">
            <h2>First Year Guide</h2>
            <v-img data-cy="resources-img" src="@/assets/resource-first-year-guide.png" contain class="image" />
          </div>
          </a>
        </v-col>
        <v-col xs="12" sm="4">
          <a target="_blank" href="https://media.csesoc.org.au/">
          <div data-cy="resources-box" class="box small rounded-lg">
            <h2>CSESoc Media</h2>
            <v-img data-cy="resources-img" src="@/assets/resource-media.png" class="image" />
          </div>
          </a>
        </v-col>
        <v-col xs="12" sm="7">
          <a target="_blank" href="">
          <div data-cy="resources-box" class="box big rounded-lg">
            <h2>Notangles</h2>
            <h3>Trimester timetabling tool - no more timetable tangles!</h3>
            <v-img data-cy="resources-img" src="@/assets/resource-notangles.png" class="image" />
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
  }
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
