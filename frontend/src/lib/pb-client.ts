import PocketBase from 'pocketbase';
import { PUBLIC_PB_ENDPOINT } from '$env/static/public';

export const pb = new PocketBase(PUBLIC_PB_ENDPOINT);

export const PB_USER_COLLECTION_ID = "users";
export const PB_DEVICES_COLLECTION_ID = "devices";
export const PB_LAYOUTS_COLLECTION_ID = "layouts";
