shows one reply
input
  - control:
    - reply_id
      difinetly show the reply of some reply
    - isShow?
      if set to false, means that this component should not show
<template>
  {{  isShow  }}
  {{ replyObj }}

</template>

<script setup lang='ts'>
import { onMounted, onUpdated,  ref } from 'vue';
import { IReply, replySoup } from '@/functions/api';

const props = defineProps<{
  // the only control param is id
  id: number,
  // this for 
  isShow?: boolean,

}>()

const isShow = ref(true)
const replyObj = ref<IReply>()

function onMountedOrUpdated() {
  replySoup.patch(1, {id:1, title: "", username: "", content: "content"} as IReply )
  // if props.isShow === false, isShow = false, else true
  if (props.isShow ?? true) {
    isShow.value = true
  } else {
    isShow.value = false
  }
  replySoup.get(props.id)
  .then(obj => {
    replyObj.value = obj
  })
}

onMounted(() => onMountedOrUpdated())
onUpdated(() => onMountedOrUpdated())

</script>