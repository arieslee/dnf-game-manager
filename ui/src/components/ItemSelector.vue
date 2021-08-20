<template>
    <v-card>
        <v-card-title>
            选择物品
            <v-spacer></v-spacer>
            <v-text-field
                v-model="itemKw"
                append-icon="mdi-magnify"
                label="按名称搜索"
                @click:append="onSearchItem"
                single-line
                hide-details
            ></v-text-field>
        </v-card-title>
        <v-divider></v-divider>
        <v-card-text style="height: 300px;">

            <v-data-table
                v-model="selectedItem"
                :headers="itemHeaders"
                :items="itemList"
                :loading="itemLoading"
                :single-select="true"
                item-key="id"
                show-select
                class="elevation-1"
                :hide-default-footer="true"
                :disable-pagination="true"
                loading-text="数据载入中..."
                no-data-text="暂无数据"
            >
                <template slot="footer.page-text" slot-scope="item">
                    {{item.itemsLength}}中的{{item.pageStart}} - {{item.pageStop}}项
                </template>
            </v-data-table>
            <div class="text-center mt-3">
                <v-pagination
                    v-model="currentPage"
                    :length="totalPage"
                    :total-visible="5"
                    @input="pageChange"
                ></v-pagination>
            </div>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
            <v-btn
                color="blue darken-1"
                text
                @click="itemClose">
                关闭
            </v-btn>
            <v-btn
                color="blue darken-1"
                text
                @click="itemSelect">
                确定
            </v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
import {getItemList} from "@/api/item";

export default {
    data(){
        return {
            itemList:[],
            pageSize: 50,
            totalCount:0,
            currentPage:1,
            totalPage:1,
            selectedItem:[],
            item:{
                id:null,
                item_name:null,
            },
            itemKw:null,
            itemLoading:false,
            itemHeaders:[
                {
                    text: '编号', value: 'id',sortable:false,
                },
                {
                    text: '名称', value: 'item_name',sortable:false,
                },
            ],
        };
    },
    created() {
        this.fetchItemList();
    },
    methods:{
        pageChange(v){
            this.fetchItemList({page:v});
        },
        onSearchItem(){
            if(!this.itemKw){
                return ;
            }
            this.itemList = [];
            this.fetchItemList({kw:this.itemKw});
        },
        fetchItemList(params){
            this.$store.dispatch('loading/show');
            this.itemLoading = true;
            let pageParams = {
                page:this.currentPage,
                page_size:this.pageSize,
            };
            if (params){
                pageParams = Object.assign(pageParams,params);
            }
            getItemList(pageParams).then((res)=>{
                this.formatList(res.data);
            }).catch((err)=>{
                this.$store.dispatch('message/error', '无法获取物品列表');
            }).finally(()=>{
                this.itemLoading = false
                this.$store.dispatch('loading/hide');
            });
        },
        formatList(data) {
            this.loading = false;
            if(!data){
                return ;
            }
            this.itemList = data['data_list'] ? data['data_list'] : [];
            if (this.itemList.length <= 0) {
                return false;
            }
            this.totalCount = data['total_count'];
            this.currentPage = data['current_page'];
            this.totalPage = data['total_page'];
        },
        itemSelect(){
            if(!this.selectedItem || this.selectedItem.length <= 0){
                this.$store.dispatch('message/error', '请选择物品');
                return false;
            }
            this.item = Object.assign(this.item,this.selectedItem[0]);
            this.item_id = this.item.id;
            this.$emit('on-confirm',this.item);
        },
        itemClose(){
            this.$emit('on-close');
        }
    },
}
</script>