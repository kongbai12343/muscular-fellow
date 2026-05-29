import { StyleSheet, Text, View } from 'react-native';

export default function ProfileScreen() {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>我的</Text>
      <Text style={styles.description}>这里将展示账户信息、基础统计和退出登录操作。</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    padding: 24,
    backgroundColor: '#F6FAF9',
  },
  title: {
    marginBottom: 12,
    color: '#111827',
    fontSize: 32,
    fontWeight: '800',
  },
  description: {
    color: '#6B7280',
    fontSize: 16,
    lineHeight: 24,
  },
});
