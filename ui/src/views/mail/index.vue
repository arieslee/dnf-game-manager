<template>
    <div class="mail-form">
        <v-container fill-height fluid class="align-start">
            <v-row class="mx-auto pt-6" width="90%">
                <v-col>
                    <div class="page-title">
                        发送邮件
                    </div>
                    <div class="box">
                        <v-form
                            ref="form"
                            v-model="valid"
                            lazy-validation
                        >
                            <v-text-field
                                v-model="item_id"
                                :rules="itemIdRules"
                                label="物品编号"
                                type="number"
                                clearable
                                :loading="loading"
                                required
                            >
                                <template v-slot:append-outer>
                                    <v-btn @click="showItemDialog" depressed>选择</v-btn>
                                </template>
                            </v-text-field>

                            <v-text-field
                                v-model="num"
                                :rules="numRules"
                                label="数量"
                                type="number"
                                clearable
                                :loading="loading"
                                required
                            ></v-text-field>

                            <v-text-field
                                v-model="coin"
                                label="金币"
                                type="number"
                                clearable
                                :loading="loading"
                                required
                            ></v-text-field>
                            <v-text-field
                                v-model="strong"
                                label="强化"
                                type="number"
                                clearable
                                :loading="loading"
                                required
                            ></v-text-field>
                            <v-btn
                                :disabled="!valid"
                                block
                                color="primary"
                                class="mr-4"
                                :loading="loading"
                                @click="sendMail">立即发送</v-btn>
                        </v-form>
                    </div>
                </v-col>
            </v-row>
        </v-container>
        <v-dialog
            v-model="itemDialog"
            scrollable
            max-width="500px"
        >
            <v-card>
                <v-card-title>
                    选择物品
                    <v-spacer></v-spacer>
                    <v-text-field
                        v-model="kw"
                        append-icon="mdi-magnify"
                        label="按名称搜索"
                        @click:append="onSearch"
                        single-line
                        hide-details
                    ></v-text-field>
                </v-card-title>
                <v-divider></v-divider>
                <v-card-text style="height: 300px;">

                    <v-data-table
                        v-model="selectedItem"
                        :headers="headers"
                        :items="itemList"
                        :loading="loading"
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
                        @click="itemDialog = false">
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
        </v-dialog>
    </div>
</template>

<script>
import {sendMail} from "@/api/mail";
import {getItemList} from "@/api/item";

export default {
    data: () => ({
        valid: true,
        item_id: null,
        itemIdRules: [
            v => (v!=='' && v !== null) || '物品代码不能为空',
        ],
        num: null,
        numRules: [
            v => !!v || '数量不能为空',
        ],
        coin: null,
        strong: null,
        loading:false,
        itemList:[],
        pageSize: 50,
        totalCount:0,
        currentPage:1,
        totalPage:1,
        pagination:{},
        itemDialog:false,
        selectedItem:null,
        headers:[
            {
                text: '编号', value: 'id',sortable:false,
            },
            {
                text: '名称', value: 'item_name',sortable:false,
            },
        ],
        kw:null,
    }),
    methods: {
        sendMail () {
            if(!this.$refs.form.validate()){
                return false;
            }
            this.$store.dispatch('loading/show');
            this.loading = true;
            let params = {
                item_id:this.item_id,
                num:this.num,
                strong:this.strong,
                coin:this.coin,
            };
            sendMail(params).then((res)=>{
                this.$store.dispatch('message/success', '邮件发送成功');
            }).catch(err=>{
                this.$store.dispatch('message/error', err.message);
            }).finally(()=>{
                this.$store.dispatch('loading/hide');
                this.loading = false;
            });
        },
        pageChange(v){
            this.fetchItemList({page:v});
        },
        onSearch(){
            if(!this.kw){
                return ;
            }
            this.itemList = [];
            this.fetchItemList({kw:this.kw});
        },
        fetchItemList(params){
            this.$store.dispatch('loading/show');
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
        showItemDialog(){
            this.itemDialog = true;
            this.fetchItemList();
        },
        itemSelect(){
            if(this.selectedItem === '' || this.selectedItem === null){
                this.$store.dispatch('message/error', '请选择物品');
                return false;
            }
            this.item_id = this.selectedItem;
            this.itemDialog = false;
        },
    },
}
</script>