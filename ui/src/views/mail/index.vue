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
                                v-model="role.charac_name"
                                label="角色名称"
                                clearable
                                :loading="loading"
                                readonly
                                required
                            >
                                <template v-slot:append-outer>
                                    <v-btn @click="showRoleDialog" depressed>选择</v-btn>
                                </template>
                            </v-text-field>
                            <v-text-field
                                v-model="item.item_name"
                                :rules="itemIdRules"
                                label="物品名称"
                                clearable
                                :loading="loading"
                                readonly
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
                            <v-text-field
                                v-model="letter_text"
                                label="文字"
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
        <v-dialog v-model="itemDialog" scrollable max-width="500px">
            <item-selector @on-close="closeItemDialog" @on-confirm="setItem"></item-selector>
        </v-dialog>

        <v-dialog v-model="roleDialog" scrollable max-width="500px">
            <role-selector @on-close="closeRoleDialog" @on-confirm="setRole"></role-selector>
        </v-dialog>
    </div>
</template>

<script>
import {sendMail} from "@/api/mail";
import ItemSelector from "@/components/ItemSelector";
import RoleSelector from "@/components/RoleSelector";
export default {
    components:{
        RoleSelector,
        ItemSelector,
    },
    data: () => ({
        valid: true,
        charac_no:null,
        letter_text:null,
        item_id: null,
        itemIdRules: [
            v => (v!=='' && v !== null) || '物品代码不能为空',
        ],
        item:{
            id:null,
            item_name:null,
        },
        num: 1,
        numRules: [
            v => !!v || '数量不能为空',
        ],
        coin: 0,
        strong: 0,
        loading:false,
        itemDialog:false,
        roleDialog:false,
        role:{
            m_id:null,
            charac_no:null,
            charac_name:null,
        },
    }),
    methods: {
        sendMail () {
            if(!this.$refs.form.validate()){
                return false;
            }
            if(!this.item_id){
                this.$store.dispatch('message/error', '请选择物品');
                return false;
            }
            if(!this.num){
                this.$store.dispatch('message/error', '物品数量不能为空');
                return false;
            }
            if(!this.charac_no){
                this.$store.dispatch('message/error', '请选择角色');
                return false;
            }
            let params = {
                item_id: this.item_id,
                num:this.num,
                coin:this.coin,
                strong:this.strong,
                charac_no:this.charac_no,
                letter_text:this.letter_text,
            };
            this.$store.dispatch('loading/show');
            this.loading = true;
            this.$store.dispatch('loading/show');
            sendMail(params).then((res)=>{
                this.loading = false;
                this.$store.dispatch('loading/hide');
                this.$store.dispatch('message/success', '邮件发送成功，小退后可以收到新邮件');
            }).catch((err)=>{
                this.loading = false;
                this.$store.dispatch('loading/hide');
                this.$store.dispatch('message/error', err.message);
            });
        },

        showItemDialog(){
            this.itemDialog = true;
        },
        showRoleDialog(){
            this.roleDialog = true;
        },
        closeItemDialog(){
            this.itemDialog = false;
        },
        closeRoleDialog(){
            this.roleDialog = false;
        },
        setItem(data){
            this.item = Object.assign(this.item, data);
            this.item_id = this.item.id;
            this.closeItemDialog();
        },
        setRole(data){
            this.role = Object.assign(this.role, data);
            console.log(this.role);
            this.charac_no = this.role.charac_no;
            this.closeRoleDialog();
        },
    },
}
</script>