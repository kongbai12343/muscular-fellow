import { StyleSheet, Text, View } from 'react-native';

export default function WorkoutsScreen() {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>训练</Text>
      <Text style={styles.description}>这里将展示训练列表、新建训练入口和筛选条件。</Text>
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
