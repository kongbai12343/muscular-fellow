import { StyleSheet, Text, View } from 'react-native';

export default function HomeScreen() {
  return (
    <View style={styles.container}>
      <Text style={styles.kicker}>Fitness Tracker</Text>
      <Text style={styles.title}>首页</Text>
      <Text style={styles.description}>这里将展示训练概览、最近训练和快捷入口。</Text>
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
  kicker: {
    marginBottom: 8,
    color: '#2563EB',
    fontSize: 14,
    fontWeight: '700',
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
