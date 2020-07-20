// Libraries
import Vuetify from 'vuetify';

// Components
import Preview from '@/components/Preview'

// Utilities
import { mount, createLocalVue } from '@vue/test-utils';

describe('Preview.vue', () => {
  let localVue;
  let vuetify;
  let wrapper;

  beforeEach(() => {
    localVue = createLocalVue();
    vuetify = new Vuetify();
    wrapper = mount(Preview, {
      localVue,
      vuetify
    });
  });

  it('should load resources as components', () => {
    const itemArray = wrapper.findAll('.v-list-item');
    //see if the resources have been fetched
    expect(itemArray.length).toEqual(4);
    
  });

  it('should allow for on hover that changes the preview', () => {
    const itemArray = wrapper.findAll(".v-list-item");
    itemArray.trigger('mouseover')
    
    //
  });

  it('should have links to resource', () => {

  });
});
