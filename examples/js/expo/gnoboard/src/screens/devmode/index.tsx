import { ConsoleView } from '@gno/components/consoleview';
import TextInput from '@gno/components/textinput';
import { screenStyleSheet as styles } from '@gno/styles';
import { useState } from 'react';
import { Linking, ScrollView, StyleSheet, View } from 'react-native';
import Button from '@gno/components/buttons';
import Layout from '@gno/components/pages';
import { GRPCError, useGnoNativeContext, ErrCode } from '@gnolang/gnonative';
import { Buffer } from 'buffer';
import ReenterPassword from '../switch-accounts/ReenterPassword';
import { Spacer } from '@gno/components/row';
import Text from '@gno/components/texts';
import { ConnectError } from '@connectrpc/connect';
import { useNavigation } from '@react-navigation/native';
import { RouterWelcomeStackProp } from '@gno/router/custom-router';
import { RoutePath } from '@gno/router/path';

function DevMode() {
  const [postContent, setPostContent] = useState('');
  const [appConsole, setAppConsole] = useState<string>('');
  const [loading, setLoading] = useState<string | undefined>(undefined);
  const [reenterPassword, setReenterPassword] = useState<string | undefined>(undefined);
  const navigate = useNavigation<RouterWelcomeStackProp>();

  const { gnonative } = useGnoNativeContext();

  const onPostPress = async () => {
    setLoading('Replying to a post...');
    setAppConsole('replying to a post...');
    const gasFee = '1000000ugnot';
    const gasWanted = BigInt(2000000);
    const args: Array<string> = ['1', '1', '1', postContent];
    try {
      for await (const response of await gnonative.call('gno.land/r/demo/boards', 'CreateReply', args, gasFee, gasWanted)) {
        console.log('response: ', response);
        setAppConsole(Buffer.from(response.result).toString());
      }
    } catch (error) {
      if (error instanceof ConnectError) {
        const err = new GRPCError(error);
        if (err.errCode() === ErrCode.ErrDecryptionFailed) {
          const account = await gnonative.getActiveAccount();
          setReenterPassword(account.key?.name);
          return;
        }
      }
      console.log(error);
      setAppConsole('error' + JSON.stringify(error));
    } finally {
      setLoading(undefined);
    }
  };

  const loadInBrowser = () => {
    Linking.openURL('https://gno.land/r/demo/boards:testboard/1').catch((err) => console.error("Couldn't load page", err));
  };

  const onCloseReenterPassword = async () => {
    setReenterPassword(undefined);
  };

  const onRenderBoard = async () => {
    navigate.navigate(RoutePath.Board, {
      board: 'gno.land/r/demo/boards',
      thread: 'testboard/1',
    });
  };

  return (
    <>
      <Layout.Container>
        <Layout.Header />
        <Layout.Body>
          <ScrollView contentContainerStyle={styles.scrollViewContent}>
            <Text.Body>Content:</Text.Body>
            <View style={customStyles.sendGroupLikeWhatsapp}>
              <TextInput style={customStyles.inputMsg} value={postContent} onChangeText={setPostContent} autoCapitalize='none' />
              <Button title='Send' onPress={onPostPress} variant='primary' style={{ width: '30%' }} loading={Boolean(loading)} />
            </View>
            <ConsoleView text={appConsole} />
            <Spacer />
            <Button title='Board Render on Browser' onPress={loadInBrowser} variant='primary' />
            <Spacer />
            <Button title='Board Render on Mobile' onPress={onRenderBoard} variant='primary' />
          </ScrollView>
        </Layout.Body>
      </Layout.Container>
      {reenterPassword ? (
        <ReenterPassword visible={Boolean(reenterPassword)} accountName={reenterPassword} onClose={onCloseReenterPassword} />
      ) : null}
    </>
  );
}

const customStyles = StyleSheet.create({
  sendGroupLikeWhatsapp: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  inputMsg: { width: '70%' },
});

export default DevMode;
