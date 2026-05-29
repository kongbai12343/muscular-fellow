import { StyleSheet, Text, View } from 'react-native';

export default function ExercisesScreen() {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>动作</Text>
      <Text style={styles.description}>这里将展示动作清单、搜索筛选和新建动作入口。</Text>
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
