<template>
  <v-app id="inspire">
    <v-navigation-drawer :dark="dark" v-model="drawer" app clipped>
      <v-list subheader>
        <v-subheader>Overview</v-subheader>
        <v-list-item v-for="item in menus[0]" :key="item.text" link @click="navTo(item.link)">
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ item.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-divider />
        <v-subheader>System</v-subheader>
        <v-list-item v-for="item in menus[1]" :key="item.text" link @click="navTo(item.link)">
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ item.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-divider />
        <v-subheader>Experiment</v-subheader>
        <v-list-item v-for="item in menus[2]" :key="item.text" link @click="navTo(item.link)">
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ item.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-divider />
        <v-subheader>Extension</v-subheader>
        <v-list-item v-for="item in menus[3]" :key="item.text" link @click="navTo(item.link)">
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ item.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-divider />
        <v-subheader>About</v-subheader>
        <v-list-item v-for="item in menus[4]" :key="item.text" link @click="navTo(item.link)">
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ item.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar :dark="dark" app clipped-left color="primary">
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title class="mr-12 align-center">
        <span class="title">AID Studio</span>
      </v-toolbar-title>
      <v-spacer />
      <v-row align="center" style="max-width: 650px">
        <v-text-field
          :append-icon-cb="() => {}"
          placeholder="Search..."
          single-line
          append-icon="mdi-magnify"
          color="white"
          hide-details
        />
      </v-row>
    </v-app-bar>
    <loading-dialog :show="isLoading"></loading-dialog>
    <alert-dialog :info="alertInfo" :title="alertTitle"></alert-dialog>
    <v-content>
      <v-container class="fill-height" style="min-width: 100%">
        <router-view />
      </v-container>
    </v-content>
    <aid-footer/>
  </v-app>
</template>

<script lang="ts">
import { mapState } from "vuex";
import {
  overview_menu,
  system_menu,
  about_menu,
  experiment_menu,
  extension_menu
} from "./router/menu";
import router from "@/router";
import LoadingDialog from "@/components/dialogs/LoadingDialog.vue";
import AlertDialog from "@/components/dialogs/Alert.vue";
import AIDFooter from "@/components/layouts/Footer.vue"
export default {
  props: {
    source: String
  },
  data: () => ({
    dark: true,
    drawer: null,
    menus: [
      overview_menu,
      system_menu,
      experiment_menu,
      extension_menu,
      about_menu
    ]
  }),
  methods: {
    navTo(link: string) {
      if (link.startsWith("http")) {
        window.open(link);
      } else {
        router.replace(link).catch(err => {
          // ignore this err
        });
      }
    }
  },
  computed: {
    ...mapState({
      isLoading: "isLoading",
      alertInfo: "alert_info",
      alertTitle: "alert_title"
    })
  },
  components: {
    "loading-dialog": LoadingDialog,
    "alert-dialog": AlertDialog,
    "aid-footer": AIDFooter,
  }
};
</script>

<style scoped>
.fill-height {
  align-items: baseline !important;
}
</style>